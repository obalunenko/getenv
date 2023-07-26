package internal

import (
	"fmt"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// notsupported is a type that is not supported by the parser.
type notsupported struct {
	_ string
}

// TestNewEnvParser tests the NewEnvParser function.
func TestNewEnvParser(t *testing.T) {
	type input struct {
		v         any
		wantPanic panicAssertionFunc
		want      EnvParser
	}

	tests := []input{
		{
			v:         "",
			wantPanic: assert.NotPanics,
			want:      stringParser(""),
		},
		{
			v:         []string{""},
			wantPanic: assert.NotPanics,
			want:      stringSliceParser([]string{""}),
		},
		{
			v:         int(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int]{},
		},
		{
			v:         []int{0},
			wantPanic: assert.NotPanics,
			want:      intSliceParser([]int{0}),
		},
		{
			v:         int8(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int8]{},
		},
		{
			v:         []int8{0},
			wantPanic: assert.NotPanics,
			want:      int8SliceParser([]int8{0}),
		},
		{
			v:         int16(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int16]{},
		},
		{
			v:         []int16{0},
			wantPanic: assert.NotPanics,
			want:      int16SliceParser([]int16{0}),
		},
		{
			v:         int32(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int32]{},
		},
		{
			v:         []int32{0},
			wantPanic: assert.NotPanics,
			want:      int32SliceParser([]int32{0}),
		},
		{
			v:         int64(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int64]{},
		},
		{
			v:         []int64{0},
			wantPanic: assert.NotPanics,
			want:      int64SliceParser([]int64{0}),
		},
		{
			v:         uint(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint]{},
		},
		{
			v:         []uint{0},
			wantPanic: assert.NotPanics,
			want:      uintSliceParser([]uint{0}),
		},
		{
			v:         uint8(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint8]{},
		},
		{
			v:         []uint8{0},
			wantPanic: assert.NotPanics,
			want:      uint8SliceParser([]uint8{0}),
		},
		{
			v:         uint16(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint16]{},
		},
		{
			v:         []uint16{0},
			wantPanic: assert.NotPanics,
			want:      uint16SliceParser([]uint16{0}),
		},
		{
			v:         uint32(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint32]{},
		},
		{
			v:         []uint32{0},
			wantPanic: assert.NotPanics,
			want:      uint32SliceParser([]uint32{0}),
		},
		{
			v:         uint64(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint64]{},
		},
		{
			v:         []uint64{0},
			wantPanic: assert.NotPanics,
			want:      uint64SliceParser([]uint64{0}),
		},
		{
			v:         float32(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[float32]{},
		},
		{
			v:         []float32{0},
			wantPanic: assert.NotPanics,
			want:      float32SliceParser([]float32{0}),
		},
		{
			v:         float64(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[float64]{},
		},
		{
			v:         []float64{0},
			wantPanic: assert.NotPanics,
			want:      float64SliceParser([]float64{0}),
		},
		{
			v:         false,
			wantPanic: assert.NotPanics,
			want:      boolParser(false),
		},
		{
			v:         notsupported{},
			wantPanic: assert.Panics,
			want:      nil,
		},
		{
			v:         nil,
			wantPanic: assert.Panics,
			want:      nil,
		},
		{
			v:         (*string)(nil),
			wantPanic: assert.Panics,
			want:      nil,
		},
		{
			v:         (*int)(nil),
			wantPanic: assert.Panics,
			want:      nil,
		},
		{
			v:         time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantPanic: assert.NotPanics,
			want:      timeParser(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			v:         []time.Time{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
			wantPanic: assert.NotPanics,
			want:      timeSliceParser([]time.Time{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)}),
		},
		{
			v:         time.Minute,
			wantPanic: assert.NotPanics,
			want:      durationParser(time.Minute),
		},
		{
			v:         []time.Duration{time.Minute},
			wantPanic: assert.NotPanics,
			want:      durationSliceParser([]time.Duration{time.Minute}),
		},
		{
			v:         getURL(t, "http://example.com"),
			wantPanic: assert.NotPanics,
			want:      urlParser(getURL(t, "http://example.com")),
		},
		{
			v:         []url.URL{getURL(t, "http://example.com")},
			wantPanic: assert.NotPanics,
			want:      urlSliceParser([]url.URL{getURL(t, "http://example.com")}),
		},
		{
			v:         getIP(t, "0.0.0.0"),
			wantPanic: assert.NotPanics,
			want:      ipParser(getIP(t, "0.0.0.0")),
		},
		{
			v:         []net.IP{getIP(t, "0.0.0.0")},
			wantPanic: assert.NotPanics,
			want:      ipSliceParser([]net.IP{getIP(t, "0.0.0.0")}),
		},
		{
			v:         uintptr(2),
			wantPanic: assert.NotPanics,
			want:      numberParser[uintptr]{},
		},
		{
			v:         []uintptr{2},
			wantPanic: assert.NotPanics,
			want:      uintptrSliceParser([]uintptr{2}),
		},
		{
			v:         complex64(1),
			wantPanic: assert.NotPanics,
			want:      complex64Parser(1),
		},
		{
			v:         []complex64{1},
			wantPanic: assert.NotPanics,
			want:      complex64SliceParser([]complex64{1}),
		},
		{
			v:         complex128(1),
			wantPanic: assert.NotPanics,
			want:      complex128Parser(1),
		},
		{
			v:         []complex128{1},
			wantPanic: assert.NotPanics,
			want:      complex128SliceParser([]complex128{1}),
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt.v), func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.IsType(t, tt.want, NewEnvParser(tt.v))
			})
		})
	}
}

// Implement tests for newUintParser function.
func Test_newUintParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "uint",
			args: args{
				v: uint(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[uint]{},
		},
		{
			name: "uint slice",
			args: args{
				v: []uint{0},
			},
			wantPanic: assert.NotPanics,
			want:      uintSliceParser([]uint{0}),
		},
		{
			name: "uint8",
			args: args{
				v: uint8(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[uint8]{},
		},
		{
			name: "uint8 slice",
			args: args{
				v: []uint8{0},
			},
			wantPanic: assert.NotPanics,
			want:      uint8SliceParser([]uint8{0}),
		},
		{
			name: "uint16",
			args: args{
				v: uint16(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[uint16]{},
		},
		{
			name: "uint16 slice",
			args: args{
				v: []uint16{0},
			},
			wantPanic: assert.NotPanics,
			want:      uint16SliceParser([]uint16{0}),
		},
		{
			name: "uint32",
			args: args{
				v: uint32(0),
			},

			wantPanic: assert.NotPanics,
			want:      numberParser[uint32]{},
		},
		{
			name: "uint32 slice",
			args: args{
				v: []uint32{0},
			},
			wantPanic: assert.NotPanics,
			want:      uint32SliceParser([]uint32{0}),
		},
		{
			name: "uint64",
			args: args{
				v: uint64(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[uint64]{},
		},
		{
			name: "uint64 slice",
			args: args{
				v: []uint64{0},
			},
			wantPanic: assert.NotPanics,
			want:      uint64SliceParser([]uint64{0}),
		},
		{
			name: "uintptr",
			args: args{
				v: uintptr(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[uintptr]{},
		},
		{
			name: "uintptr slice",
			args: args{
				v: []uintptr{0},
			},
			wantPanic: assert.NotPanics,
			want:      uintptrSliceParser([]uintptr{0}),
		},
		{
			name: "invalid type",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newUintParser(tt.args.v))
			})
		})
	}
}

// Implement tests for newFloatParser function.
func Test_newFloatParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "float32",
			args: args{
				v: float32(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[float32]{},
		},
		{
			name: "float32 slice",
			args: args{
				v: []float32{0},
			},
			wantPanic: assert.NotPanics,
			want:      float32SliceParser([]float32{0}),
		},
		{
			name: "float64",
			args: args{
				v: float64(0),
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[float64]{},
		},
		{
			name: "float64 slice",
			args: args{
				v: []float64{0},
			},
			wantPanic: assert.NotPanics,
			want:      float64SliceParser([]float64{0}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newFloatParser(tt.args.v))
			})
		})
	}
}

// Implement tests for newTimeParser function.
func Test_newTimeParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "time.Time",
			args: args{
				v: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantPanic: assert.NotPanics,
			want:      timeParser(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "time.Time slice",
			args: args{
				v: []time.Time{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			wantPanic: assert.NotPanics,
			want:      timeSliceParser([]time.Time{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
		{
			name: "time.Time pointer",
			args: args{
				v: &time.Time{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
		{
			name: "time.Time slice pointer",
			args: args{
				v: &[]time.Time{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
		{
			name: "time.Duration",
			args: args{
				v: time.Second,
			},
			wantPanic: assert.NotPanics,
			want:      durationParser(time.Second),
		},
		{
			name: "time.Duration slice",
			args: args{
				v: []time.Duration{time.Second},
			},
			wantPanic: assert.NotPanics,
			want:      durationSliceParser([]time.Duration{time.Second}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newTimeParser(tt.args.v))
			})
		})
	}
}

// Implement tests for newStringParser function.
func Test_newStringParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "string",
			args: args{
				v: "string",
			},
			wantPanic: assert.NotPanics,
			want:      stringParser("string"),
		},
		{
			name: "string slice",
			args: args{
				v: []string{"string"},
			},
			wantPanic: assert.NotPanics,
			want:      stringSliceParser([]string{"string"}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newStringParser(tt.args.v))
			})
		})
	}
}

// Implement tests for newBoolParser function.
func Test_newBoolParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "bool",
			args: args{
				v: true,
			},
			wantPanic: assert.NotPanics,
			want:      boolParser(true),
		},
		{
			name: "bool slice",
			args: args{
				v: []bool{true},
			},
			wantPanic: assert.NotPanics,
			want:      boolSliceParser([]bool{true}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newBoolParser(tt.args.v))
			})
		})
	}
}

// Test_newIntParser tests newIntParser function.
func Test_newIntParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "int",
			args: args{
				v: 1,
			},
			wantPanic: assert.NotPanics,
			want:      numberParser[int]{},
		},
		{
			name: "int slice",
			args: args{
				v: []int{1},
			},
			wantPanic: assert.NotPanics,
			want:      intSliceParser([]int{1}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newIntParser(tt.args.v))
			})
		})
	}
}

// Test_newIPParser tests newIPParser function.
func Test_newIPParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "net.IP",
			args: args{
				v: net.IPv4(127, 0, 0, 1),
			},
			wantPanic: assert.NotPanics,
			want:      ipParser(net.IPv4(127, 0, 0, 1)),
		},
		{
			name: "net.IP slice",
			args: args{
				v: []net.IP{net.IPv4(127, 0, 0, 1)},
			},
			wantPanic: assert.NotPanics,
			want:      ipSliceParser([]net.IP{net.IPv4(127, 0, 0, 1)}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newIPParser(tt.args.v))
			})
		})
	}
}

// Test_newURLParser tests newURLParser function.
func Test_newURLParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "url.URL",
			args: args{
				v: url.URL{
					Scheme: "http",
					Host:   "localhost",
				},
			},
			wantPanic: assert.NotPanics,
			want: urlParser(url.URL{
				Scheme: "http",
				Host:   "localhost",
			}),
		},
		{
			name: "url.URL slice",
			args: args{
				v: []url.URL{
					{
						Scheme: "http",
						Host:   "localhost",
					},
				},
			},
			wantPanic: assert.NotPanics,
			want: urlSliceParser([]url.URL{
				{
					Scheme: "http",
					Host:   "localhost",
				},
			}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newURLParser(tt.args.v))
			})
		})
	}
}

// Test_newComplexParser tests newComplexParser function.
func Test_newComplexParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		wantPanic panicAssertionFunc
		want      EnvParser
	}{
		{
			name: "complex64",
			args: args{
				v: complex64(1),
			},
			wantPanic: assert.NotPanics,
			want:      complex64Parser(complex64(1)),
		},
		{
			name: "complex64 slice",
			args: args{
				v: []complex64{1},
			},
			wantPanic: assert.NotPanics,
			want:      complex64SliceParser([]complex64{1}),
		},
		{
			name: "complex128",
			args: args{
				v: complex128(1),
			},
			wantPanic: assert.NotPanics,
			want:      complex128Parser(complex128(1)),
		},
		{
			name: "complex128 slice",
			args: args{
				v: []complex128{1},
			},
			wantPanic: assert.NotPanics,
			want:      complex128SliceParser([]complex128{1}),
		},
		{
			name: "not supported",
			args: args{
				v: notsupported{},
			},
			wantPanic: assert.NotPanics,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantPanic(t, func() {
				assert.Equal(t, tt.want, newComplexParser(tt.args.v))
			})
		})
	}
}

func Test_ParseEnv(t *testing.T) {
	type args struct {
		key       string
		defaltVal any
		in2       Parameters
	}

	tests := []struct {
		name    string
		precond precondition
		s       EnvParser
		args    args
		want    any
	}{
		{
			name: "boolParser",
			s:    boolParser(true),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: false,
				in2:       Parameters{},
			},
			want: true,
		},
		{
			name: "stringParser",
			s:    stringParser(""),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "golly",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: "test",
				in2:       Parameters{},
			},
			want: "golly",
		},
		{
			name: "stringSliceParser",
			s:    stringSliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "golly,sally",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []string{},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []string{"golly", "sally"},
		},
		{
			name: "float64Parser",
			s:    float64Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "-1.2",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: float64(99),
				in2:       Parameters{},
			},
			want: -1.2,
		},
		{
			name: "float32Parser",
			s:    float32Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "-1.2",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: float32(99),
				in2:       Parameters{},
			},
			want: float32(-1.2),
		},
		{
			name: "float32SliceParser",
			s:    float32SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "-1.2,0.2",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []float32{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []float32{-1.2, 0.2},
		},
		{
			name: "float64SliceParser",
			s:    float64SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "-1.2,0.2",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []float64{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []float64{-1.2, 0.2},
		},
		{
			name: "uint8Parser",
			s:    uint8Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: uint8(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: uint8(12),
		},
		{
			name: "uint64Parser",
			s:    uint64Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: uint64(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: uint64(12),
		},
		{
			name: "uint64SliceParser",
			s:    uint64SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []uint64{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []uint64{12, 89},
		},
		{
			name: "uint16SliceParser",
			s:    uint16SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []uint16{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []uint16{12, 89},
		},
		{
			name: "uint32Parser",
			s:    uint32Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: uint32(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: uint32(12),
		},
		{
			name: "uint16Parser",
			s:    uint16Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: uint16(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: uint16(12),
		},
		{
			name: "int8Parser",
			s:    int8Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: int8(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: int8(12),
		},
		{
			name: "int16Parser",
			s:    int16Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: int16(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: int16(12),
		},
		{
			name: "uint32SliceParser",
			s:    uint32SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []uint32{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []uint32{12, 89},
		},
		{
			name: "uintParser",
			s:    uintParser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: uint(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: uint(12),
		},
		{
			name: "uintSliceParser",
			s:    uintSliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []uint{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []uint{12, 89},
		},
		{
			name: "int64Parser",
			s:    int64Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: int64(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: int64(12),
		},
		{
			name: "int32Parser",
			s:    int32Parser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: int32(99),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: int32(12),
		},
		{
			name: "int16SliceParser",
			s:    int16SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []int16{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []int16{12, 89},
		},
		{
			name: "int32SliceParser",
			s:    int32SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []int32{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []int32{12, 89},
		},
		{
			name: "int64SliceParser",
			s:    int64SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []int64{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []int64{12, 89},
		},
		{
			name: "intParser",
			s:    intParser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: 99,
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: 12,
		},
		{
			name: "uint8SliceParser",
			s:    uint8SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []uint8{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []uint8{12, 89},
		},
		{
			name: "int8SliceParser",
			s:    int8SliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []int8{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []int8{12, 89},
		},
		{
			name: "intSliceParser",
			s:    intSliceParser(nil),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12,89",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []int{99},
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: []int{12, 89},
		},
		{
			name: "durationParser",
			s:    durationParser(0),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12s",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: time.Duration(0),
				in2: Parameters{
					Separator: ",",
					Layout:    "",
				},
			},
			want: time.Second * 12,
		},
		{
			name: "timeParser",
			s:    timeParser(time.Time{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2022-03-24",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: time.Time{},
				in2: Parameters{
					Separator: "",
					Layout:    time.DateOnly,
				},
			},
			want: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "timeSliceParser",
			s:    timeSliceParser([]time.Time{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2022-03-24,2023-03-24",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []time.Time{},
				in2: Parameters{
					Separator: ",",
					Layout:    time.DateOnly,
				},
			},
			want: []time.Time{
				time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.March, 24, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "durationSliceParser",
			s:    durationSliceParser([]time.Duration{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2m,3h",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []time.Duration{},
				in2: Parameters{
					Separator: ",",
				},
			},
			want: []time.Duration{
				2 * time.Minute,
				3 * time.Hour,
			},
		},
		{
			name: "urlParser",
			s:    urlParser(url.URL{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "https://google.com",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: url.URL{},
				in2: Parameters{
					Separator: "",
					Layout:    time.DateOnly,
				},
			},
			want: getURL(t, "https://google.com"),
		},
		{
			name: "urlSliceParser",
			s:    urlSliceParser([]url.URL{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "https://google.com,https://bing.com",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []url.URL{},
				in2: Parameters{
					Separator: ",",
					Layout:    time.DateOnly,
				},
			},
			want: []url.URL{
				getURL(t, "https://google.com"),
				getURL(t, "https://bing.com"),
			},
		},
		{
			name: "ipParser",
			s:    ipParser(net.IP{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2001:cb8::17",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: net.IP{},
				in2:       Parameters{},
			},
			want: getIP(t, "2001:cb8::17"),
		},
		{
			name: "ipSliceParser",
			s:    ipSliceParser([]net.IP{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2001:cb8::17,192.168.0.1",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []net.IP{},
				in2: Parameters{
					Separator: ",",
				},
			},
			want: []net.IP{
				getIP(t, "2001:cb8::17"),
				getIP(t, "192.168.0.1"),
			},
		},
		{
			name: "boolSliceParser",
			s:    boolSliceParser([]bool{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,false",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []bool{},
				in2: Parameters{
					Separator: ",",
				},
			},
			want: []bool{
				true,
				false,
			},
		},
		{
			name: "uintptrSliceParser",
			s:    uintptrSliceParser([]uintptr{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []uintptr{},
				in2: Parameters{
					Separator: ",",
				},
			},
			want: []uintptr{
				1,
				2,
			},
		},
		{
			name: "uintptrParser",
			s:    uintptrParser(uintptr(0)),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: uintptr(0),
			},
			want: uintptr(1),
		},
		{
			name: "complex64SliceParser",
			s:    complex64SliceParser([]complex64{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1+2i,3+4i",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []complex64{},
				in2: Parameters{
					Separator: ",",
				},
			},
			want: []complex64{
				complex(1, 2),
				complex(3, 4),
			},
		},
		{
			name: "complex64Parser",
			s:    complex64Parser(complex64(0)),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1+2i",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: complex64(0),
			},
			want: complex64(complex(1, 2)),
		},
		{
			name: "complex128SliceParser",
			s:    complex128SliceParser([]complex128{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1+2i,3+4i",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: []complex128{},
				in2: Parameters{
					Separator: ",",
				},
			},
			want: []complex128{
				complex(1, 2),
				complex(3, 4),
			},
		},
		{
			name: "complex128Parser",
			s:    complex128Parser(complex128(0)),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1+2i",
				},
			},
			args: args{
				key:       testEnvKey,
				defaltVal: complex128(0),
			},
			want: complex128(complex(1, 2)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, testEnvKey)

			assert.Equal(t, tt.want, tt.s.ParseEnv(tt.args.key, tt.args.defaltVal, tt.args.in2))
		})
	}
}
