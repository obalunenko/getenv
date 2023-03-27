// Package getenv provides functionality for loading environment variables and parse them into go builtin types.
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
// - uint32
// - []uint32
// - float32
// - float64
// - []float64
// - time.Time
// - []time.Time
// - time.Duration
// - bool
// - url.URL
package getenv

import (
	"github.com/obalunenko/getenv/internal"
	"github.com/obalunenko/getenv/option"
)

// EnvOrDefault retrieves the value of the environment variable named
// by the key.
// If variable not set or value is empty - defaultVal will be returned.
func EnvOrDefault[T internal.EnvParsable](key string, defaultVal T, options ...option.Option) T {
	w := internal.NewEnvParser(defaultVal)

	params := newParseParams(options)

	val := w.ParseEnv(key, defaultVal, params)

	return val.(T)
}

func newParseParams(opts []option.Option) internal.Parameters {
	var p internal.Parameters

	for _, opt := range opts {
		opt.Apply(&p)
	}

	return p
}
