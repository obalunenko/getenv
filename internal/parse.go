// Package internal provides internal implementation logic for environment variables parsing.
package internal

import (
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const decimalBase = 10

// Parse converts the raw value of an environment variable to T.
// Slices are split by Parameters.Separator, time.Time is parsed with Parameters.Layout.
// Every returned error wraps ErrInvalidValue.
//
//nolint:gocyclo,funlen // Flat dispatch: one case per supported type.
func Parse[T EnvParsable](raw string, p Parameters) (T, error) {
	var (
		res T
		err error
	)

	switch v := any(&res).(type) {
	case *string:
		*v = raw
	case *[]string:
		*v, err = parseSlice(raw, p.Separator, parseString)
	case *int:
		*v, err = parseInt[int](raw)
	case *[]int:
		*v, err = parseSlice(raw, p.Separator, parseInt[int])
	case *int8:
		*v, err = parseInt[int8](raw)
	case *[]int8:
		*v, err = parseSlice(raw, p.Separator, parseInt[int8])
	case *int16:
		*v, err = parseInt[int16](raw)
	case *[]int16:
		*v, err = parseSlice(raw, p.Separator, parseInt[int16])
	case *int32:
		*v, err = parseInt[int32](raw)
	case *[]int32:
		*v, err = parseSlice(raw, p.Separator, parseInt[int32])
	case *int64:
		*v, err = parseInt[int64](raw)
	case *[]int64:
		*v, err = parseSlice(raw, p.Separator, parseInt[int64])
	case *uint:
		*v, err = parseUint[uint](raw)
	case *[]uint:
		*v, err = parseSlice(raw, p.Separator, parseUint[uint])
	case *uint8:
		*v, err = parseUint[uint8](raw)
	case *[]uint8:
		*v, err = parseSlice(raw, p.Separator, parseUint[uint8])
	case *uint16:
		*v, err = parseUint[uint16](raw)
	case *[]uint16:
		*v, err = parseSlice(raw, p.Separator, parseUint[uint16])
	case *uint32:
		*v, err = parseUint[uint32](raw)
	case *[]uint32:
		*v, err = parseSlice(raw, p.Separator, parseUint[uint32])
	case *uint64:
		*v, err = parseUint[uint64](raw)
	case *[]uint64:
		*v, err = parseSlice(raw, p.Separator, parseUint[uint64])
	case *uintptr:
		*v, err = parseUint[uintptr](raw)
	case *[]uintptr:
		*v, err = parseSlice(raw, p.Separator, parseUint[uintptr])
	case *float32:
		*v, err = parseFloat[float32](raw)
	case *[]float32:
		*v, err = parseSlice(raw, p.Separator, parseFloat[float32])
	case *float64:
		*v, err = parseFloat[float64](raw)
	case *[]float64:
		*v, err = parseSlice(raw, p.Separator, parseFloat[float64])
	case *complex64:
		*v, err = parseComplex[complex64](raw)
	case *[]complex64:
		*v, err = parseSlice(raw, p.Separator, parseComplex[complex64])
	case *complex128:
		*v, err = parseComplex[complex128](raw)
	case *[]complex128:
		*v, err = parseSlice(raw, p.Separator, parseComplex[complex128])
	case *bool:
		*v, err = strconv.ParseBool(raw)
	case *[]bool:
		*v, err = parseSlice(raw, p.Separator, strconv.ParseBool)
	case *time.Time:
		*v, err = parseTime(p.Layout)(raw)
	case *[]time.Time:
		*v, err = parseSlice(raw, p.Separator, parseTime(p.Layout))
	case *time.Duration:
		*v, err = time.ParseDuration(raw)
	case *[]time.Duration:
		*v, err = parseSlice(raw, p.Separator, time.ParseDuration)
	case *url.URL:
		*v, err = parseURL(raw)
	case *[]url.URL:
		*v, err = parseSlice(raw, p.Separator, parseURL)
	case *net.IP:
		*v, err = parseIP(raw)
	case *[]net.IP:
		*v, err = parseSlice(raw, p.Separator, parseIP)
	case *netip.Addr:
		*v, err = netip.ParseAddr(raw)
	case *[]netip.Addr:
		*v, err = parseSlice(raw, p.Separator, netip.ParseAddr)
	case *netip.Prefix:
		*v, err = netip.ParsePrefix(raw)
	case *[]netip.Prefix:
		*v, err = parseSlice(raw, p.Separator, netip.ParsePrefix)
	case *net.HardwareAddr:
		*v, err = net.ParseMAC(raw)
	case *[]net.HardwareAddr:
		*v, err = parseSlice(raw, p.Separator, net.ParseMAC)
	default:
		// Unreachable while every EnvParsable type has a case above.
		panic(fmt.Sprintf("getenv: no parser for %T", res))
	}

	if err != nil {
		var zero T

		return zero, fmt.Errorf("%w: %w", err, ErrInvalidValue)
	}

	return res, nil
}

// parseSlice splits raw by sep and parses every element with parse.
func parseSlice[T any](raw, sep string, parse func(string) (T, error)) ([]T, error) {
	if sep == "" {
		return nil, errEmptySeparator
	}

	parts := strings.Split(raw, sep)
	res := make([]T, 0, len(parts))

	for i, s := range parts {
		v, err := parse(s)
		if err != nil {
			return nil, fmt.Errorf("element %d: %w", i, err)
		}

		res = append(res, v)
	}

	return res, nil
}

func parseString(s string) (string, error) {
	return s, nil
}

func parseInt[T Int](s string) (T, error) {
	v, err := strconv.ParseInt(s, decimalBase, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

func parseUint[T Uint](s string) (T, error) {
	v, err := strconv.ParseUint(s, decimalBase, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

func parseFloat[T Float](s string) (T, error) {
	v, err := strconv.ParseFloat(s, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

func parseComplex[T Complex](s string) (T, error) {
	v, err := strconv.ParseComplex(s, bitSize[T]())
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

// bitSize returns the size of T in bits, as the strconv parsers expect it.
func bitSize[T Number | Complex]() int {
	return reflect.TypeFor[T]().Bits()
}

func parseTime(layout string) func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		return time.Parse(layout, s)
	}
}

func parseURL(s string) (url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return url.URL{}, err
	}

	return *u, nil
}

func parseIP(s string) (net.IP, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address %q", s)
	}

	return ip, nil
}
