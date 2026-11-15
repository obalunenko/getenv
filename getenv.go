// Package getenv provides a simple way to get environment variables.
// It's generic and supports built-in types, custom types and slices of them.
// Note: empty environment variable values are treated as "not set".
//
// Slices are split by "," and time.Time values are parsed with the time.RFC3339 layout,
// unless option.WithSeparator or option.WithTimeLayout sets another one.
//
// Types supported:
// - string
// - []string
// - int
// - []int
// - int8
// - []int8
// - int16
// - []int16
// - int32
// - []int32
// - int64
// - []int64
// - uint8
// - []uint8
// - uint16
// - []uint16
// - uint64
// - []uint64
// - uint
// - []uint
// - uintptr
// - []uintptr
// - uint32
// - []uint32
// - float32
// - []float32
// - float64
// - []float64
// - time.Time
// - []time.Time
// - time.Duration
// - []time.Duration
// - bool
// - []bool
// - url.URL
// - []url.URL
// - net.IP
// - []net.IP
// - netip.Addr
// - []netip.Addr
// - netip.Prefix
// - []netip.Prefix
// - net.HardwareAddr
// - []net.HardwareAddr
// - complex64
// - []complex64
// - complex128
// - []complex128
//
// Custom types are supported as well, and so are slices of the supported types,
// named slice types such as type Tags []string included.
// A value of a custom type T is parsed with the first one that fits:
// - the ParseEnv method of T, see EnvParser;
// - the UnmarshalText method of *T, or of T if it is a pointer type, see encoding.TextUnmarshaler;
// - the underlying type of T, if it is a string, bool or numeric type, as for type Port uint16.
// Note that type Timeout time.Duration is parsed as int64 nanoseconds, not as a duration:
// its underlying type is int64, and it does not get the methods of time.Duration.
// To parse it as a duration, implement EnvParser (see its example),
// or declare an alias instead, type Timeout = time.Duration, if it needs no methods of its own.
//
// Env and EnvOrDefault panic for any other type, even when the variable is not set.
package getenv

import (
	"fmt"
	"os"
	"time"

	"github.com/obalunenko/getenv/internal"
	"github.com/obalunenko/getenv/option"
)

var (
	// ErrNotSet is an error that is returned when the environment variable is not set.
	ErrNotSet = internal.ErrNotSet
	// ErrInvalidValue is an error that is returned when the environment variable is not valid.
	ErrInvalidValue = internal.ErrInvalidValue
	// ErrUnsupportedType is wrapped by the value that Env and EnvOrDefault panic with for a type they cannot parse.
	ErrUnsupportedType = internal.ErrUnsupportedType
)

// Defaults for the options that were not passed.
const (
	defaultSeparator  = ","
	defaultTimeLayout = time.RFC3339
)

// Parsable is a constraint for the built-in types that Env and EnvOrDefault can parse.
// The list of supported types is in the package documentation.
type Parsable = internal.EnvParsable

// EnvParser is implemented by custom types that parse themselves from an environment variable.
// ParseEnv is called on the zero value of T, so it needs a value receiver, unless T is a pointer type.
//
// Implementing EnvParser is checked at run time.
// To check it at compile time, assign a value of the type to an EnvParser variable:
//
//	var _ getenv.EnvParser[Level] = Level(0)
type EnvParser[T EnvParser[T]] interface {
	ParseEnv(raw string) (T, error)
}

// Env retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value will be parsed and returned.
// Otherwise, an error will be returned.
// Env panics if T is not supported (see the package documentation), even when the variable is not set.
func Env[T any](key string, options ...option.Option) (T, error) {
	var zero T

	raw, ok := os.LookupEnv(key)
	if !ok || raw == "" {
		// An unsupported T is reported even when there is nothing to parse.
		internal.MustSupport[T]()

		return zero, fmt.Errorf("failed to get environment variable[%s]: %w", key, ErrNotSet)
	}

	val, err := internal.Parse[T](raw, newParseParams(options))
	if err != nil {
		return zero, fmt.Errorf("failed to parse environment variable[%s]: %w", key, err)
	}

	return val, nil
}

// EnvOrDefault retrieves the value of the environment variable named by the key.
// If the variable is set and its value can be parsed, the parsed value is returned.
// Otherwise, defaultVal is returned: both when the variable is not set (or empty)
// and when its value is invalid. Use Env to tell these cases apart.
// EnvOrDefault panics if T is not supported, like Env.
func EnvOrDefault[T any](key string, defaultVal T, options ...option.Option) T {
	val, err := Env[T](key, options...)
	if err != nil {
		return defaultVal
	}

	return val
}

// newParseParams creates new parameters from the defaults and options.
func newParseParams(opts []option.Option) internal.Parameters {
	p := internal.Parameters{
		Separator: defaultSeparator,
		Layout:    defaultTimeLayout,
	}

	for _, opt := range opts {
		opt.Apply(&p)
	}

	return p
}
