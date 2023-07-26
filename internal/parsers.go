package internal

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func getRawVal(key string) (string, error) {
	env, ok := os.LookupEnv(key)
	if !ok || env == "" {
		return "", ErrNotSet
	}

	return env, nil
}

// getStringOrDefault retrieves the string value of the environment variable named
// by the key.
// If the variable is not set or the value is empty - defaultVal will be returned.
func getStringOrDefault(key, defaultVal string) string {
	val, err := getRawVal(key)
	if err != nil {
		return defaultVal
	}

	return val
}

func getBool(key string) (bool, error) {
	env, err := getRawVal(key)
	if err != nil {
		return false, err
	}

	val, err := strconv.ParseBool(env)
	if err != nil {
		return false, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	return val, nil
}

// getBoolOrDefault retrieves the bool value of the environment variable named
// by the key.
// If the variable is not set or the value is empty - defaultVal will be returned.
func getBoolOrDefault(key string, defaultVal bool) bool {
	b, err := getBool(key)
	if err != nil {
		return defaultVal
	}

	return b
}

func getBoolSlice(key string, sep string) ([]bool, error) {
	if sep == "" {
		return nil, ErrInvalidValue
	}

	env, err := getRawVal(key)
	if err != nil {
		return nil, err
	}

	val := strings.Split(env, sep)

	b := make([]bool, 0, len(val))

	for _, s := range val {
		v, err := strconv.ParseBool(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
		}

		b = append(b, v)
	}

	return b, nil
}

// getBoolSliceOrDefault retrieves the bool slice value of the environment variable named
// by the key and separated by sep.
// If the variable is not set or the value is empty - defaultVal will be returned.
func getBoolSliceOrDefault(key string, defaultVal []bool, sep string) []bool {
	val, err := getBoolSlice(key, sep)
	if err != nil {
		return defaultVal
	}

	return val
}

func getStringSlice(key string, sep string) ([]string, error) {
	if sep == "" {
		return nil, ErrInvalidValue
	}

	env, err := getRawVal(key)
	if err != nil {
		return nil, err
	}

	val := strings.Split(env, sep)

	return val, nil
}

// getStringSliceOrDefault retrieves the string slice value of the environment variable named
// by the key and separated by sep.
// If the variable is not set or the value is empty - defaultVal will be returned.
func getStringSliceOrDefault(key string, defaultVal []string, sep string) []string {
	val, err := getStringSlice(key, sep)
	if err != nil {
		return defaultVal
	}

	return val
}

func parseNumberGen[T Number](raw string) (T, error) {
	var tt T

	const (
		base = 10
	)

	rt := reflect.TypeOf(tt)

	switch rt.Kind() { //nolint:exhaustive // All supported types are covered.
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(raw, base, rt.Bits())
		if err != nil {
			return tt, ErrInvalidValue
		}

		return any(T(val)).(T), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		val, err := strconv.ParseUint(raw, base, rt.Bits())
		if err != nil {
			return tt, ErrInvalidValue
		}

		return any(T(val)).(T), nil
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(raw, rt.Bits())
		if err != nil {
			return tt, ErrInvalidValue
		}

		return any(T(val)).(T), nil
	default:
		return tt, ErrInvalidValue
	}
}

func parseNumberSliceGen[S []T, T Number](raw []string) (S, error) {
	var tt S

	val := make(S, 0, len(raw))

	for _, s := range raw {
		v, err := parseNumberGen[T](s)
		if err != nil {
			return tt, err
		}

		val = append(val, v)
	}

	return val, nil
}

func getNumberSliceGen[S []T, T Number](key string, sep string) (S, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	return parseNumberSliceGen[S, T](env)
}

func getSliceOrDefaultGen[S []T, T Number](key string, defaultVal S, sep string) S {
	val, err := getNumberSliceGen[S, T](key, sep)
	if err != nil {
		return defaultVal
	}

	return val
}

func getNumberGen[T Number](key string) (T, error) {
	env, err := getRawVal(key)
	if err != nil {
		return 0, err
	}

	return parseNumberGen[T](env)
}

func numberOrDefaultGen[T Number](key string, defaultVal T) T {
	val, err := getNumberGen[T](key)
	if err != nil {
		return defaultVal
	}

	return val
}

func getDuration(key string) (time.Duration, error) {
	env, err := getRawVal(key)
	if err != nil {
		return 0, err
	}

	val, err := time.ParseDuration(env)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	return val, nil
}

// getDurationOrDefault retrieves the time.Duration value of the environment variable named
// by the key.
// If the variable is not set or the value is empty - defaultVal will be returned.
func getDurationOrDefault(key string, defaultVal time.Duration) time.Duration {
	val, err := getDuration(key)
	if err != nil {
		return defaultVal
	}

	return val
}

func getTime(key string, layout string) (time.Time, error) {
	env, err := getRawVal(key)
	if err != nil {
		return time.Time{}, err
	}

	val, err := time.Parse(layout, env)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	return val, nil
}

// timeOrDefault retrieves the time.Time value of the environment variable named
// by the key represented by layout.
// If the variable is not set or the value is empty - defaultVal will be returned.
func timeOrDefault(key string, defaultVal time.Time, layout string) time.Time {
	val, err := getTime(key, layout)
	if err != nil {
		return defaultVal
	}

	return val
}

func getTimeSlice(key string, layout string, sep string) ([]time.Time, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]time.Time, 0, len(env))

	for _, s := range env {
		v, err := time.Parse(layout, s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

// timeSliceOrDefault retrieves the []time.Time value of the environment variable named
// by the key represented by layout.
// If the variable is not set or the value is empty - defaultVal will be returned.
func timeSliceOrDefault(key string, defaultVal []time.Time, layout, separator string) []time.Time {
	val, err := getTimeSlice(key, layout, separator)
	if err != nil {
		return defaultVal
	}

	return val
}

func getDurationSlice(key string, sep string) ([]time.Duration, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]time.Duration, 0, len(env))

	for _, s := range env {
		v, err := time.ParseDuration(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
		}

		val = append(val, v)
	}

	return val, nil
}

// durationSliceOrDefault retrieves the []time.Duration value of the environment variable named
// by the key represented by layout.
// If the variable is not set or the value is empty - defaultVal will be returned.
func durationSliceOrDefault(key string, defaultVal []time.Duration, separator string) []time.Duration {
	val, err := getDurationSlice(key, separator)
	if err != nil {
		return defaultVal
	}

	return val
}

func getURL(key string) (url.URL, error) {
	env, err := getRawVal(key)
	if err != nil {
		return url.URL{}, err
	}

	val, err := url.Parse(env)
	if err != nil {
		return url.URL{}, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	return *val, nil
}

// urlOrDefault retrieves the url.URL value of the environment variable named
// by the key represented by layout.
// If variable not set or value is empty - defaultVal will be returned.
func urlOrDefault(key string, defaultVal url.URL) url.URL {
	val, err := getURL(key)
	if err != nil {
		return defaultVal
	}

	return val
}

func getURLSlice(key string, sep string) ([]url.URL, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]url.URL, 0, len(env))

	for _, s := range env {
		v, err := url.Parse(s)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
		}

		val = append(val, *v)
	}

	return val, nil
}

// urlSliceOrDefault retrieves the url.URL slice value of the environment variable named
// by the key and separated by sep.
// If variable not set or value is empty - defaultVal will be returned.
func urlSliceOrDefault(key string, defaultVal []url.URL, sep string) []url.URL {
	val, err := getURLSlice(key, sep)
	if err != nil {
		return defaultVal
	}

	return val
}

func getIP(key string) (net.IP, error) {
	env, err := getRawVal(key)
	if err != nil {
		return nil, err
	}

	val := net.ParseIP(env)
	if val == nil {
		return nil, ErrInvalidValue
	}

	return val, nil
}

// ipOrDefault retrieves the net.IP value of the environment variable named
// by the key represented by layout.
// If variable not set or value is empty - defaultVal will be returned.
func ipOrDefault(key string, defaultVal net.IP) net.IP {
	val, err := getIP(key)
	if err != nil {
		return defaultVal
	}

	return val
}

func getIPSlice(key string, sep string) ([]net.IP, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	val := make([]net.IP, 0, len(env))

	for _, s := range env {
		v := net.ParseIP(s)
		if v == nil {
			return nil, ErrInvalidValue
		}

		val = append(val, v)
	}

	return val, nil
}

// ipSliceOrDefault retrieves the net.IP slice value of the environment variable named
// by the key and separated by sep.
// If variable not set or value is empty - defaultVal will be returned.
func ipSliceOrDefault(key string, defaultVal []net.IP, sep string) []net.IP {
	val, err := getIPSlice(key, sep)
	if err != nil {
		return defaultVal
	}

	return val
}

func parseComplexGen[T Complex](raw string) (T, error) {
	var tt T

	var bitsize int

	switch any(tt).(type) {
	case complex64:
		bitsize = 64
	case complex128:
		bitsize = 128
	}

	val, err := strconv.ParseComplex(raw, bitsize)
	if err != nil {
		return tt, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	return any(T(val)).(T), nil
}

func parseComplexSliceGen[S []T, T Complex](raw []string) (S, error) {
	var tt S

	val := make(S, 0, len(raw))

	for _, s := range raw {
		v, err := parseComplexGen[T](s)
		if err != nil {
			return tt, err
		}

		val = append(val, v)
	}

	return val, nil
}

func getComplexSliceGen[S []T, T Complex](key string, sep string) (S, error) {
	env, err := getStringSlice(key, sep)
	if err != nil {
		return nil, err
	}

	return parseComplexSliceGen[S, T](env)
}

func getComplexGen[T Complex](key string) (T, error) {
	env, err := getRawVal(key)
	if err != nil {
		return 0, err
	}

	return parseComplexGen[T](env)
}

func complexOrDefaultGen[T Complex](key string, defaultVal T) T {
	val, err := getComplexGen[T](key)
	if err != nil {
		return defaultVal
	}

	return val
}

func complexSliceOrDefaultGen[S []T, T Complex](key string, defaultVal S, sep string) S {
	val, err := getComplexSliceGen[S, T](key, sep)
	if err != nil {
		return defaultVal
	}

	return val
}
