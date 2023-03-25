package internal

import (
	"time"
)

type (
	// EnvParsable is a constraint for supported environment variable types parsers.
	EnvParsable interface {
		String | Int | Float | Time | bool
	}

	// String is a constraint for strings and slice of strings.
	String interface {
		string | []string
	}

	// Int is a constraint for integer and slice of integers.
	Int interface {
		int | []int | int8 | []int8 | int32 | []int32 | int64 | []int64 | uint64 | []uint64 | uint | []uint | []uint32 | uint32
	}

	// Float is a constraint for floats and slice of floats.
	Float interface {
		float64 | []float64
	}

	// Time is a constraint for time.Time and time.Duration.
	Time interface {
		time.Time | time.Duration
	}
)
