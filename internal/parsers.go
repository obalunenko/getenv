package internal

import (
	"net"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// stringOrDefault retrieves the string value of the environment variable named
// by the key.
// If the variable is not set or the value is empty - defaultVal will be returned.
func stringOrDefault(key, defaultVal string) string {
	env, ok := os.LookupEnv(key)
	if !ok || env == "" {
		return defaultVal
	}

	return env
}

// boolOrDefault retrieves the bool value of the environment variable named
// by the key.
// If the variable is not set or the value is empty - defaultVal will be returned.
func boolOrDefault(key string, defaultVal bool) bool {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := strconv.ParseBool(env)
	if err != nil {
		return defaultVal
	}

	return val
}

// boolSliceOrDefault retrieves the bool slice value of the environment variable named
// by the key and separated by sep.
// If the variable is not set or the value is empty - defaultVal will be returned.
func boolSliceOrDefault(key string, defaultVal []bool, sep string) []bool {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]bool, 0, len(valraw))

	for _, s := range valraw {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return defaultVal
		}

		val = append(val, b)
	}

	return val
}

// stringSliceOrDefault retrieves the string slice value of the environment variable named
// by the key and separated by sep.
// If the variable is not set or the value is empty - defaultVal will be returned.
func stringSliceOrDefault(key string, defaultVal []string, sep string) []string {
	if sep == "" {
		return defaultVal
	}

	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val := strings.Split(env, sep)

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

func numberSliceOrDefaultGen[S []T, T Number](key string, defaultVal S, sep string) S {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val, err := parseNumberSliceGen[S, T](valraw)
	if err != nil {
		return defaultVal
	}

	return val
}

func numberOrDefaultGen[T Number](key string, defaultVal T) T {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := parseNumberGen[T](env)
	if err != nil {
		return defaultVal
	}

	return val
}

// durationOrDefault retrieves the time.Duration value of the environment variable named
// by the key.
// If the variable is not set or the value is empty - defaultVal will be returned.
func durationOrDefault(key string, defaultVal time.Duration) time.Duration {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := time.ParseDuration(env)
	if err != nil {
		return defaultVal
	}

	return val
}

// timeOrDefault retrieves the time.Time value of the environment variable named
// by the key represented by layout.
// If the variable is not set or the value is empty - defaultVal will be returned.
func timeOrDefault(key string, defaultVal time.Time, layout string) time.Time {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := time.Parse(layout, env)
	if err != nil {
		return defaultVal
	}

	return val
}

// timeSliceOrDefault retrieves the []time.Time value of the environment variable named
// by the key represented by layout.
// If the variable is not set or the value is empty - defaultVal will be returned.
func timeSliceOrDefault(key string, defaultVal []time.Time, layout, separator string) []time.Time {
	valraw := stringSliceOrDefault(key, nil, separator)
	if valraw == nil {
		return defaultVal
	}

	val := make([]time.Time, 0, len(valraw))

	for _, s := range valraw {
		v, err := time.Parse(layout, s)
		if err != nil {
			return defaultVal
		}

		val = append(val, v)
	}

	return val
}

// durationSliceOrDefault retrieves the []time.Duration value of the environment variable named
// by the key represented by layout.
// If the variable is not set or the value is empty - defaultVal will be returned.
func durationSliceOrDefault(key string, defaultVal []time.Duration, separator string) []time.Duration {
	valraw := stringSliceOrDefault(key, nil, separator)
	if valraw == nil {
		return defaultVal
	}

	val := make([]time.Duration, 0, len(valraw))

	for _, s := range valraw {
		v, err := time.ParseDuration(s)
		if err != nil {
			return defaultVal
		}

		val = append(val, v)
	}

	return val
}

// urlOrDefault retrieves the url.URL value of the environment variable named
// by the key represented by layout.
// If variable not set or value is empty - defaultVal will be returned.
func urlOrDefault(key string, defaultVal url.URL) url.URL {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := url.Parse(env)
	if err != nil {
		return defaultVal
	}

	return *val
}

// urlSliceOrDefault retrieves the url.URL slice value of the environment variable named
// by the key and separated by sep.
// If variable not set or value is empty - defaultVal will be returned.
func urlSliceOrDefault(key string, defaultVal []url.URL, sep string) []url.URL {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]url.URL, 0, len(valraw))

	for _, s := range valraw {
		v, err := url.Parse(s)
		if err != nil {
			return defaultVal
		}

		val = append(val, *v)
	}

	return val
}

// ipOrDefault retrieves the net.IP value of the environment variable named
// by the key represented by layout.
// If variable not set or value is empty - defaultVal will be returned.
func ipOrDefault(key string, defaultVal net.IP) net.IP {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val := net.ParseIP(env)
	if val == nil {
		return defaultVal
	}

	return val
}

// ipSliceOrDefault retrieves the net.IP slice value of the environment variable named
// by the key and separated by sep.
// If variable not set or value is empty - defaultVal will be returned.
func ipSliceOrDefault(key string, defaultVal []net.IP, sep string) []net.IP {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]net.IP, 0, len(valraw))

	for _, s := range valraw {
		v := net.ParseIP(s)
		if v == nil {
			return defaultVal
		}

		val = append(val, v)
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
		return tt, ErrInvalidValue
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

func complexOrDefaultGen[T Complex](key string, defaultVal T) T {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := parseComplexGen[T](env)
	if err != nil {
		return defaultVal
	}

	return val
}

func complexSliceOrDefaultGen[S []T, T Complex](key string, defaultVal S, sep string) S {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val, err := parseComplexSliceGen[S, T](valraw)
	if err != nil {
		return defaultVal
	}

	return val
}
