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
			want:      numberSliceParser[[]int, int]{},
		},
		{
			v:         int8(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int8]{},
		},
		{
			v:         []int8{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]int8, int8]{},
		},
		{
			v:         int16(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int16]{},
		},
		{
			v:         []int16{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]int16, int16]{},
		},
		{
			v:         int32(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int32]{},
		},
		{
			v:         []int32{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]int32, int32]{},
		},
		{
			v:         int64(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[int64]{},
		},
		{
			v:         []int64{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]int64, int64]{},
		},
		{
			v:         uint(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint]{},
		},
		{
			v:         []uint{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]uint, uint]{},
		},
		{
			v:         uint8(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint8]{},
		},
		{
			v:         []uint8{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]uint8, uint8]{},
		},
		{
			v:         uint16(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint16]{},
		},
		{
			v:         []uint16{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]uint16, uint16]{},
		},
		{
			v:         uint32(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint32]{},
		},
		{
			v:         []uint32{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]uint32, uint32]{},
		},
		{
			v:         uint64(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[uint64]{},
		},
		{
			v:         []uint64{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]uint64, uint64]{},
		},
		{
			v:         float32(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[float32]{},
		},
		{
			v:         []float32{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]float32, float32]{},
		},
		{
			v:         float64(0),
			wantPanic: assert.NotPanics,
			want:      numberParser[float64]{},
		},
		{
			v:         []float64{0},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]float64, float64]{},
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
			v:         getTestURL(t, "http://example.com"),
			wantPanic: assert.NotPanics,
			want:      urlParser(getTestURL(t, "http://example.com")),
		},
		{
			v:         []url.URL{getTestURL(t, "http://example.com")},
			wantPanic: assert.NotPanics,
			want:      urlSliceParser([]url.URL{getTestURL(t, "http://example.com")}),
		},
		{
			v:         getTestIP(t, "0.0.0.0"),
			wantPanic: assert.NotPanics,
			want:      ipParser(getTestIP(t, "0.0.0.0")),
		},
		{
			v:         []net.IP{getTestIP(t, "0.0.0.0")},
			wantPanic: assert.NotPanics,
			want:      ipSliceParser([]net.IP{getTestIP(t, "0.0.0.0")}),
		},
		{
			v:         uintptr(2),
			wantPanic: assert.NotPanics,
			want:      numberParser[uintptr]{},
		},
		{
			v:         []uintptr{2},
			wantPanic: assert.NotPanics,
			want:      numberSliceParser[[]uintptr, uintptr]{},
		},
		{
			v:         complex64(1),
			wantPanic: assert.NotPanics,
			want:      complexParser[complex64]{},
		},
		{
			v:         []complex64{1},
			wantPanic: assert.NotPanics,
			want:      complexSliceParser[[]complex64, complex64]{},
		},
		{
			v:         complex128(1),
			wantPanic: assert.NotPanics,
			want:      complexParser[complex128]{},
		},
		{
			v:         []complex128{1},
			wantPanic: assert.NotPanics,
			want:      complexSliceParser[[]complex128, complex128]{},
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
