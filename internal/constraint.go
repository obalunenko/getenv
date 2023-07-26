package internal

import (
	"net"
	"net/url"
	"time"
)

type (
	// EnvParsable is a constraint for types that can be parsed from environment variable.
	EnvParsable interface {
		String | Int | IntSlice | Uint | UintSlice | Float | FloatSlice | Time | Bool | URL | IP | Complex | ComplexSlice
	}

	// String is a constraint for string and slice of strings.
	String interface {
		string | []string
	}

	// Int is a constraint for integer and slice of integers.
	Int interface {
		int | int8 | int16 | int32 | int64
	}

	// IntSlice is a constraint for slice of integers.
	IntSlice interface {
		[]int | []int8 | []int16 | []int32 | []int64
	}

	// UintSlice is a constraint for slice of unsigned integers.
	UintSlice interface {
		[]uint | []uint8 | []uint16 | []uint32 | []uint64 | []uintptr
	}

	// Uint is a constraint for unsigned integer and slice of unsigned integers.
	Uint interface {
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	}

	// FloatSlice is a constraint for slice of floats.
	FloatSlice interface {
		[]float32 | []float64
	}

	// Float is a constraint for float and slice of floats.
	Float interface {
		float32 | float64
	}

	// Time is a constraint for time.Time and slice of time.Time.
	Time interface {
		time.Time | []time.Time | time.Duration | []time.Duration
	}

	// Bool is a constraint for bool and slice of bool.
	Bool interface {
		bool | []bool
	}

	// URL is a constraint for url.URL and slice of url.URL.
	URL interface {
		url.URL | []url.URL
	}

	// IP is a constraint for net.IP and slice of net.IP.
	IP interface {
		net.IP | []net.IP
	}

	// ComplexSlice is a constraint for slice of complex.
	ComplexSlice interface {
		[]complex64 | []complex128
	}

	// Complex is a constraint for complex and slice of complex.
	Complex interface {
		complex64 | complex128
	}
)
