package internal

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type notsupported struct {
	name string
}

func TestNewEnvParser(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name      string
		args      args
		want      EnvParser
		wantPanic panicAssertionFunc
	}{
		{
			name: "int32",
			args: args{
				v: int32(1),
			},
			want:      int32Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "int64",
			args: args{
				v: int64(1),
			},
			want:      int64Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]int32",
			args: args{
				v: []int32{1},
			},
			want:      int32SliceParser([]int32{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]int64",
			args: args{
				v: []int64{1},
			},
			want:      int64SliceParser([]int64{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "uint64",
			args: args{
				v: uint64(1),
			},
			want:      uint64Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]uint64",
			args: args{
				v: []uint64{1},
			},
			want:      uint64SliceParser([]uint64{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "uint",
			args: args{
				v: uint(1),
			},
			want:      uintParser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]uint",
			args: args{
				v: []uint{1},
			},
			want:      uintSliceParser{1},
			wantPanic: assert.NotPanics,
		},
		{
			name: "uint8",
			args: args{
				v: uint8(1),
			},
			want:      uint8Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "uint16",
			args: args{
				v: uint16(1),
			},
			want:      uint16Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "uint32",
			args: args{
				v: uint32(1),
			},
			want:      uint32Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]uint32",
			args: args{
				v: []uint32{1},
			},
			want:      uint32SliceParser{1},
			wantPanic: assert.NotPanics,
		},
		{
			name: "int",
			args: args{
				v: 1,
			},
			want:      intParser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "int16",
			args: args{
				v: int16(1),
			},
			want:      int16Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "int8",
			args: args{
				v: int8(1),
			},
			want:      int8Parser(1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]int",
			args: args{
				v: []int{1},
			},
			want:      intSliceParser([]int{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]int16",
			args: args{
				v: []int16{1},
			},
			want:      int16SliceParser([]int16{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]uint8",
			args: args{
				v: []uint8{1},
			},
			want:      uint8SliceParser([]uint8{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]int8",
			args: args{
				v: []int8{1},
			},
			want:      int8SliceParser([]int8{1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "string",
			args: args{
				v: "s",
			},
			want:      stringParser("s"),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]string",
			args: args{
				v: []string{"s"},
			},
			want:      stringSliceParser([]string{"s"}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "bool",
			args: args{
				v: true,
			},
			want:      boolParser(true),
			wantPanic: assert.NotPanics,
		},
		{
			name: "float32",
			args: args{
				v: float32(1.1),
			},
			want:      float32Parser(float32(1.1)),
			wantPanic: assert.NotPanics,
		},
		{
			name: "float64",
			args: args{
				v: 1.1,
			},
			want:      float64Parser(1.1),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]float64",
			args: args{
				v: []float64{1.1},
			},
			want:      float64SliceParser([]float64{1.1}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "time.Time",
			args: args{
				v: time.Time{},
			},
			want:      timeParser(time.Time{}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "[]time.Time",
			args: args{
				v: []time.Time{},
			},
			want:      timeSliceParser([]time.Time{}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "time.Duration",
			args: args{
				v: time.Minute,
			},
			want:      durationParser(time.Minute),
			wantPanic: assert.NotPanics,
		},
		{
			name: "url.URL",
			args: args{
				v: url.URL{},
			},
			want:      urlParser(url.URL{}),
			wantPanic: assert.NotPanics,
		},
		{
			name: "not supported - panics",
			args: args{
				v: notsupported{
					name: "name",
				},
			},
			want:      nil,
			wantPanic: assert.Panics,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got EnvParser

			if !tt.wantPanic(t, func() {
				got = NewEnvParser(tt.args.v)
			}) {
				return
			}

			assert.Equal(t, tt.want, got)
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
			name: "float64Parser",
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
			name: "urlParser",
			s:    urlParser(url.URL{}),
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "https.google.com",
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
			want: getURL(t, "https.google.com"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, testEnvKey)

			assert.Equal(t, tt.want, tt.s.ParseEnv(tt.args.key, tt.args.defaltVal, tt.args.in2))
		})
	}
}
