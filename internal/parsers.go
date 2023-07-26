package internal

import (
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// stringOrDefault retrieves the string value of the environment variable named
// by the key.
// If variable not set or value is empty - defaultVal will be returned.
func stringOrDefault(key, defaultVal string) string {
	env, ok := os.LookupEnv(key)
	if !ok || env == "" {
		return defaultVal
	}

	return env
}

// boolOrDefault retrieves the bool value of the environment variable named
// by the key.
// If variable not set or value is empty - defaultVal will be returned.
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
// If variable not set or value is empty - defaultVal will be returned.
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
// If variable not set or value is empty - defaultVal will be returned.
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

func floatOrDefaultGen[T float32 | float64](key string, defaultVal T) T {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := parseFloatGen[T](env)
	if err != nil {
		return defaultVal
	}

	return val
}

func parseFloatGen[T float32 | float64](raw string) (T, error) {
	var tt T

	const (
		bitsize = 64
	)

	var (
		castFn func(val float64) (T, error)
	)

	switch any(tt).(type) {
	case float32:
		castFn = func(val float64) (T, error) {
			t, ok := any(float32(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case float64:
		castFn = func(val float64) (T, error) {
			t, ok := any(val).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	}

	val, err := strconv.ParseFloat(raw, bitsize)
	if err != nil {
		return tt, ErrInvalidValue
	}

	return castFn(val)
}

func parseIntGen[T int | int8 | int16 | int32 | int64](raw string) (T, error) {
	var tt T

	const (
		base = 10
	)

	var (
		bitsize int
		castFn  func(val int64) (T, error)
	)

	switch any(tt).(type) {
	case int:
		bitsize = 0
		castFn = func(val int64) (T, error) {
			t, ok := any(int(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case int8:
		bitsize = 8
		castFn = func(val int64) (T, error) {
			t, ok := any(int8(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case int16:
		bitsize = 16
		castFn = func(val int64) (T, error) {
			t, ok := any(int16(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case int32:
		bitsize = 32
		castFn = func(val int64) (T, error) {
			t, ok := any(int32(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case int64:
		bitsize = 64
		castFn = func(val int64) (T, error) {
			t, ok := any(val).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	}

	val, err := strconv.ParseInt(raw, base, bitsize)
	if err != nil {
		return tt, ErrInvalidValue
	}

	return castFn(val)
}

func intOrDefaultGen[T int | int8 | int16 | int32 | int64](key string, defaultVal T) T {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := parseIntGen[T](env)
	if err != nil {
		return defaultVal
	}

	return val
}

func parseIntSliceGen[T int | int8 | int16 | int32 | int64](raw []string) ([]T, error) {
	var tt []T

	val := make([]T, 0, len(raw))

	for _, s := range raw {
		v, err := parseIntGen[T](s)
		if err != nil {
			return tt, err
		}

		val = append(val, v)
	}

	return val, nil
}

func parseFloatSliceGen[T float32 | float64](raw []string) ([]T, error) {
	var tt []T

	val := make([]T, 0, len(raw))

	for _, s := range raw {
		v, err := parseFloatGen[T](s)
		if err != nil {
			return tt, err
		}

		val = append(val, v)
	}

	return val, nil
}

func floatSliceOrDefaultGen[T float32 | float64](key string, defaultVal []T, sep string) []T {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val, err := parseFloatSliceGen[T](valraw)
	if err != nil {
		return defaultVal
	}

	return val
}

func intSliceOrDefaultGen[T int | int8 | int16 | int32 | int64](key string, defaultVal []T, sep string) []T {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val, err := parseIntSliceGen[T](valraw)
	if err != nil {
		return defaultVal
	}

	return val
}

// durationOrDefault retrieves the time.Duration value of the environment variable named
// by the key.
// If variable not set or value is empty - defaultVal will be returned.
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
// If variable not set or value is empty - defaultVal will be returned.
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
// If variable not set or value is empty - defaultVal will be returned.
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
// If variable not set or value is empty - defaultVal will be returned.
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

// uint64SliceOrDefault retrieves the uint64 slice value of the environment variable named
// by the key and separated by sep.
// If variable not set or value is empty - defaultVal will be returned.
func uint64SliceOrDefault(key string, defaultVal []uint64, sep string) []uint64 {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]uint64, 0, len(valraw))

	const (
		base    = 10
		bitsize = 64
	)

	for _, s := range valraw {
		v, err := strconv.ParseUint(s, base, bitsize)
		if err != nil {
			return defaultVal
		}

		val = append(val, v)
	}

	return val
}

func parseUintGen[T uint | uint8 | uint16 | uint32 | uint64 | uintptr](raw string) (T, error) {
	var tt T

	const (
		base = 10
	)

	var (
		bitsize int
		castFn  func(val uint64) (T, error)
	)

	switch any(tt).(type) {
	case uint:
		bitsize = 0
		castFn = func(val uint64) (T, error) {
			t, ok := any(uint(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case uint8:
		bitsize = 8
		castFn = func(val uint64) (T, error) {
			t, ok := any(uint8(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case uint16:
		bitsize = 16
		castFn = func(val uint64) (T, error) {
			t, ok := any(uint16(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case uint32:
		bitsize = 32
		castFn = func(val uint64) (T, error) {
			t, ok := any(uint32(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case uint64:
		bitsize = 64
		castFn = func(val uint64) (T, error) {
			t, ok := any(val).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case uintptr:
		bitsize = 0
		castFn = func(val uint64) (T, error) {
			t, ok := any(uintptr(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	}

	val, err := strconv.ParseUint(raw, base, bitsize)
	if err != nil {
		return tt, ErrInvalidValue
	}

	return castFn(val)
}

func uintOrDefaultGen[T uint | uint8 | uint16 | uint32 | uint64 | uintptr](key string, defaultVal T) T {
	env := stringOrDefault(key, "")
	if env == "" {
		return defaultVal
	}

	val, err := parseUintGen[T](env)
	if err != nil {
		return defaultVal
	}

	return val
}

func uintSliceOrDefaultGen[T uint | uint8 | uint16 | uint32 | uint64 | uintptr](key string, defaultVal []T, sep string) []T {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]T, 0, len(valraw))

	for _, s := range valraw {
		v, err := parseUintGen[T](s)
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

// uintptrSliceOrDefault retrieves the uintptr slice value of the environment variable named
// by the key and separated by sep.
// If variable not set or value is empty - defaultVal will be returned.
func uintptrSliceOrDefault(key string, defaultVal []uintptr, sep string) []uintptr {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]uintptr, 0, len(valraw))

	const (
		base    = 10
		bitsize = 0
	)

	for _, s := range valraw {
		v, err := strconv.ParseUint(s, base, bitsize)
		if err != nil {
			return defaultVal
		}

		val = append(val, uintptr(v))
	}

	return val
}

func parseComplexGen[T complex64 | complex128](raw string) (T, error) {
	var tt T

	var (
		bitsize int
		castFn  func(val complex128) (T, error)
	)

	switch any(tt).(type) {
	case complex64:
		bitsize = 64

		castFn = func(val complex128) (T, error) {
			t, ok := any(complex64(val)).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	case complex128:
		bitsize = 128

		castFn = func(val complex128) (T, error) {
			t, ok := any(val).(T)
			if !ok {
				return tt, ErrInvalidValue
			}

			return t, nil
		}
	}

	val, err := strconv.ParseComplex(raw, bitsize)
	if err != nil {
		return tt, ErrInvalidValue
	}

	return castFn(val)
}

func complexOrDefaultGen[T complex64 | complex128](key string, defaultVal T) T {
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

func complexSliceOrDefaultGen[T complex64 | complex128](key string, defaultVal []T, sep string) []T {
	valraw := stringSliceOrDefault(key, nil, sep)
	if valraw == nil {
		return defaultVal
	}

	val := make([]T, 0, len(valraw))

	for _, s := range valraw {
		v, err := parseComplexGen[T](s)
		if err != nil {
			return defaultVal
		}

		val = append(val, v)
	}

	return val
}
