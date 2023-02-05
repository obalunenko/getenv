// Package internal provides internal implementation logic for environment variables parsing.
package internal

import (
	"fmt"
	"time"
)

// NewEnvParser is a constructor for EnvParser.
func NewEnvParser(v any) EnvParser {
	var p EnvParser

	switch i := v.(type) {
	case string:
		p = stringParser(v.(string))
	case []string:
		p = stringSliceParser(v.([]string))
	case int:
		p = intParser(v.(int))
	case []int:
		p = intSliceParser(v.([]int))
	case bool:
		p = boolParser(v.(bool))
	case int64:
		p = int64Parser(v.(int64))
	case []int64:
		p = int64SliceParser(v.([]int64))
	case float64:
		p = float64Parser(v.(float64))
	case []float64:
		p = float64SliceParser(v.([]float64))
	case uint64:
		p = uint64Parser(v.(uint64))
	case []uint64:
		p = uint64SliceParser(v.([]uint64))
	case uint:
		p = uintParser(v.(uint))
	case []uint:
		p = uintSliceParser(v.([]uint))
	case time.Time:
		p = timeParser(v.(time.Time))
	case time.Duration:
		p = durationParser(v.(time.Duration))
	default:
		panic(fmt.Sprintf("unsupported type :%T", i))
	}

	return p
}

// EnvParser interface for parsing environment variables.
type EnvParser interface {
	ParseEnv(key string, defaltVal any, options Parameters) any
}

type stringParser string

func (s stringParser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := stringOrDefault(key, defaltVal.(string))

	return val
}

type stringSliceParser []string

func (s stringSliceParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	sep := options.Separator

	val := stringSliceOrDefault(key, defaltVal.([]string), sep)

	return val
}

type intParser int

func (i intParser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := intOrDefault(key, defaltVal.(int))

	return val
}

type intSliceParser []int

func (i intSliceParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	sep := options.Separator

	val := intSliceOrDefault(key, defaltVal.([]int), sep)

	return val
}

type float64SliceParser []float64

func (i float64SliceParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	sep := options.Separator

	val := float64SliceOrDefault(key, defaltVal.([]float64), sep)

	return val
}

type int64Parser int64

func (i int64Parser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := int64OrDefault(key, defaltVal.(int64))

	return val
}

type int64SliceParser []int64

func (i int64SliceParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	sep := options.Separator

	val := int64SliceOrDefault(key, defaltVal.([]int64), sep)

	return val
}

type float64Parser float64

func (f float64Parser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := float64OrDefault(key, defaltVal.(float64))

	return val
}

type boolParser bool

func (b boolParser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := boolOrDefault(key, defaltVal.(bool))

	return val
}

type timeParser time.Time

func (t timeParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	layout := options.Layout

	val := timeOrDefault(key, defaltVal.(time.Time), layout)

	return val
}

type durationParser time.Duration

func (d durationParser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := durationOrDefault(key, defaltVal.(time.Duration))

	return val
}

type uint64Parser uint64

func (d uint64Parser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := uint64OrDefault(key, defaltVal.(uint64))

	return val
}

type uint64SliceParser []uint64

func (i uint64SliceParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	sep := options.Separator

	val := uint64SliceOrDefault(key, defaltVal.([]uint64), sep)

	return val
}

type uintParser uint

func (d uintParser) ParseEnv(key string, defaltVal any, _ Parameters) any {
	val := uintOrDefault(key, defaltVal.(uint))

	return val
}

type uintSliceParser []uint

func (i uintSliceParser) ParseEnv(key string, defaltVal any, options Parameters) any {
	sep := options.Separator

	val := uintSliceOrDefault(key, defaltVal.([]uint), sep)

	return val
}
