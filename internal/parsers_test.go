package internal

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Benchmark_float64SliceOrDefault(b *testing.B) {
	p := precondition{
		setenv: setenv{
			isSet: true,
			val:   "1235.67,87.98",
		},
	}

	p.maybeSetEnv(b, testEnvKey)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = float64SliceOrDefault(testEnvKey, []float64{}, ",")
	}
}

func Test_intOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}

	type expected struct {
		val int
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "128",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 42,
			},
			expected: expected{
				val: 42,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 128,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 44,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := intOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_stringOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}

	type expected struct {
		val string
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: "default",
			},
			expected: expected{
				val: "default",
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: "default",
			},
			expected: expected{
				val: "newval",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := stringOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int64
	}

	type expected struct {
		val int64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 956,
			},
			expected: expected{
				val: 956,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1024",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 42,
			},
			expected: expected{
				val: 1024,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 44,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int64OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int8OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int8
	}

	type expected struct {
		val int8
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 95,
			},
			expected: expected{
				val: 95,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "10",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 42,
			},
			expected: expected{
				val: 10,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 44,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int8OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int16OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int16
	}

	type expected struct {
		val int16
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 956,
			},
			expected: expected{
				val: 956,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1024",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 42,
			},
			expected: expected{
				val: 1024,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 44,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int16OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int32OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int32
	}

	type expected struct {
		val int32
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 956,
			},
			expected: expected{
				val: 956,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1024",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 42,
			},
			expected: expected{
				val: 1024,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 44,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int32OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_float32OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal float32
	}

	type expected struct {
		val float32
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: float32(956.02),
			},
			expected: expected{
				val: float32(956.02),
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1024.123",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: float32(42),
			},
			expected: expected{
				val: float32(1024.123),
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: float32(44),
			},
			expected: expected{
				val: float32(44),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := float32OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_float64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal float64
	}

	type expected struct {
		val float64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 956.02,
			},
			expected: expected{
				val: 956.02,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1024.123",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 42,
			},
			expected: expected{
				val: 1024.123,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 44,
			},
			expected: expected{
				val: 44,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := float64OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_boolOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}

	type expected struct {
		val bool
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "true",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: false,
			},
			expected: expected{
				val: false,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: false,
			},
			expected: expected{
				val: true,
			},
		},
		{
			name: "invalid env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: true,
			},
			expected: expected{
				val: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := boolOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_stringSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []string
		sep        string
	}

	type expected struct {
		val []string
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "true,newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []string{"no values"},
				sep:        ",",
			},
			expected: expected{
				val: []string{"no values"},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []string{"no values"},
				sep:        ",",
			},
			expected: expected{
				val: []string{"true", "newval"},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,newval",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []string{"no values"},
				sep:        "",
			},
			expected: expected{
				val: []string{"no values"},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []string{"no values"},
			},
			expected: expected{
				val: []string{"no values"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := stringSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_intSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int
		sep        string
	}

	type expected struct {
		val []int
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int{-99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int{-99},
				sep:        "",
			},
			expected: expected{
				val: []int{-99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int{-99},
				sep:        "|",
			},
			expected: expected{
				val: []int{-99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int{-99},
			},
			expected: expected{
				val: []int{-99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := intSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_float32SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []float32
		sep        string
	}

	type expected struct {
		val []float32
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float32{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float32{-99.99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float32{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float32{1.05, 2.07},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float32{-99.99},
				sep:        "",
			},
			expected: expected{
				val: []float32{-99.99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float32{-99.99},
				sep:        "|",
			},
			expected: expected{
				val: []float32{-99.99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float32{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float32{-99.99},
			},
		},
		{
			name: "malformed data value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sss,sss",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float32{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float32{-99.99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := float32SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_float64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []float64
		sep        string
	}

	type expected struct {
		val []float64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float64{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float64{-99.99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float64{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float64{1.05, 2.07},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float64{-99.99},
				sep:        "",
			},
			expected: expected{
				val: []float64{-99.99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float64{-99.99},
				sep:        "|",
			},
			expected: expected{
				val: []float64{-99.99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float64{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float64{-99.99},
			},
		},
		{
			name: "malformed data value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sss,sss",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []float64{-99.99},
				sep:        ",",
			},
			expected: expected{
				val: []float64{-99.99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := float64SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int16SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int16
		sep        string
	}

	type expected struct {
		val []int16
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int16{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int16{-99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int16{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int16{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int16{-99},
				sep:        "",
			},
			expected: expected{
				val: []int16{-99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int16{-99},
				sep:        "|",
			},
			expected: expected{
				val: []int16{-99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int16{-99},
			},
			expected: expected{
				val: []int16{-99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int16{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int16{-99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int16SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int32SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int32
		sep        string
	}

	type expected struct {
		val []int32
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int32{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int32{-99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int32{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int32{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int32{-99},
				sep:        "",
			},
			expected: expected{
				val: []int32{-99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int32{-99},
				sep:        "|",
			},
			expected: expected{
				val: []int32{-99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int32{-99},
			},
			expected: expected{
				val: []int32{-99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int32{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int32{-99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int32SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uintSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint
		sep        string
	}

	type expected struct {
		val []uint
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint{99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint{99},
				sep:        "",
			},
			expected: expected{
				val: []uint{99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint{99},
				sep:        "|",
			},
			expected: expected{
				val: []uint{99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint{99},
			},
			expected: expected{
				val: []uint{99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uintSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint8SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint8
		sep        string
	}

	type expected struct {
		val []uint8
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint8{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint8{99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint8{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint8{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint8{99},
				sep:        "",
			},
			expected: expected{
				val: []uint8{99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint8{99},
				sep:        "|",
			},
			expected: expected{
				val: []uint8{99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint8{99},
			},
			expected: expected{
				val: []uint8{99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint8{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint8{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint8SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint16SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint16
		sep        string
	}

	type expected struct {
		val []uint16
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint16{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint16{99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint16{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint16{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint16{99},
				sep:        "",
			},
			expected: expected{
				val: []uint16{99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint16{99},
				sep:        "|",
			},
			expected: expected{
				val: []uint16{99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint16{99},
			},
			expected: expected{
				val: []uint16{99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint16{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint16{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint16SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint32SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint32
		sep        string
	}

	type expected struct {
		val []uint32
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint32{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint32{99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint32{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint32{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint32{99},
				sep:        "",
			},
			expected: expected{
				val: []uint32{99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint32{99},
				sep:        "|",
			},
			expected: expected{
				val: []uint32{99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint32{99},
			},
			expected: expected{
				val: []uint32{99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint32{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint32{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint32SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int8SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int8
		sep        string
	}

	type expected struct {
		val []int8
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int8{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int8{-99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int8{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int8{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int8{-99},
				sep:        "",
			},
			expected: expected{
				val: []int8{-99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int8{-99},
				sep:        "|",
			},
			expected: expected{
				val: []int8{-99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int8{-99},
			},
			expected: expected{
				val: []int8{-99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int8{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int8{-99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int8SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_int64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int64
		sep        string
	}

	type expected struct {
		val []int64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int64{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int64{-99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int64{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int64{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int64{-99},
				sep:        "",
			},
			expected: expected{
				val: []int64{-99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int64{-99},
				sep:        "|",
			},
			expected: expected{
				val: []int64{-99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int64{-99},
			},
			expected: expected{
				val: []int64{-99},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []int64{-99},
				sep:        ",",
			},
			expected: expected{
				val: []int64{-99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := int64SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_timeOrDefault(t *testing.T) {
	const layout = "2006/02/01 15:04"

	type args struct {
		key        string
		defaultVal time.Time
		layout     string
	}

	type expected struct {
		val time.Time
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "2018/21/04 22:30",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2018/21/04 22:30",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2018, 04, 21, 22, 30, 0, 0, time.UTC),
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "20222-sdslll",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := timeOrDefault(tt.args.key, tt.args.defaultVal, tt.args.layout)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_urlOrDefault(t *testing.T) {
	const rawDefault = "https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"

	type args struct {
		key        string
		defaultVal url.URL
	}

	type expected struct {
		val url.URL
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "postgres://user:pass@host.com:5432/path?k=v#f",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getURL(t, rawDefault),
			},
			expected: expected{
				val: getURL(t, rawDefault),
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "postgres://user:pass@host.com:5432/path?k=v#f",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getURL(t, rawDefault),
			},
			expected: expected{
				val: getURL(t, "postgres://user:pass@host.com:5432/path?k=v#f"),
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "postgres://user:pass@host.com:5432/path?k=v#f%%2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getURL(t, rawDefault),
			},
			expected: expected{
				val: getURL(t, rawDefault),
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getURL(t, rawDefault),
			},
			expected: expected{
				val: getURL(t, rawDefault),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := urlOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_timeSliceOrDefault(t *testing.T) {
	const layout = "2006/02/01 15:04"

	type args struct {
		key        string
		defaultVal []time.Time
		layout     string
		separator  string
	}

	type expected struct {
		val []time.Time
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "2018/21/04 22:30,2023/21/04 22:30",
				},
			},
			args: args{
				key: testEnvKey,
				defaultVal: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2018/21/04 22:30,2023/21/04 22:30",
				},
			},
			args: args{
				key: testEnvKey,
				defaultVal: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2018, 04, 21, 22, 30, 0, 0, time.UTC),
					time.Date(2023, 04, 21, 22, 30, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				defaultVal: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "20222-sdslll",
				},
			},
			args: args{
				key: testEnvKey,
				defaultVal: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2021, 04, 21, 22, 30, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := timeSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.layout, tt.args.separator)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_durationSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []time.Duration
		separator  string
	}

	type expected struct {
		val []time.Duration
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "2m,3h",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []time.Duration{time.Second},
				separator:  ",",
			},
			expected: expected{
				val: []time.Duration{time.Second},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2m,3h",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []time.Duration{time.Second},
				separator:  ",",
			},
			expected: expected{
				val: []time.Duration{
					2 * time.Minute, 3 * time.Hour,
				},
			},
		},
		{
			name: "env set, corrupted - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2m,3hddd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []time.Duration{time.Second},
				separator:  ",",
			},
			expected: expected{
				val: []time.Duration{time.Second},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []time.Duration{time.Second},
				separator:  ",",
			},
			expected: expected{
				val: []time.Duration{time.Second},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := durationSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.separator)

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_durationOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal time.Duration
	}

	type expected struct {
		val time.Duration
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12s",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Second * 42,
			},
			expected: expected{
				val: time.Second * 42,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12s",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Second * 42,
			},
			expected: expected{
				val: time.Second * 12,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Second * 42,
			},
			expected: expected{
				val: time.Second * 42,
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "yyydd88",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: time.Second * 42,
			},
			expected: expected{
				val: time.Second * 42,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := durationOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint64
	}

	type expected struct {
		val uint64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 12,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint64OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint64
		sep        string
	}

	type expected struct {
		val []uint64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1,27",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint64{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint64{99},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint64{99},
				sep:        ",",
			},
			expected: expected{
				val: []uint64{1, 2},
			},
		},
		{
			name: "env set, no separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint64{99},
				sep:        "",
			},
			expected: expected{
				val: []uint64{99},
			},
		},
		{
			name: "env set, wrong separator - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint64{99},
				sep:        "|",
			},
			expected: expected{
				val: []uint64{99},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uint64{99},
			},
			expected: expected{
				val: []uint64{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint64SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint8OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint8
	}

	type expected struct {
		val uint8
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 99,
			},
			expected: expected{
				val: 99,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 99,
			},
			expected: expected{
				val: 12,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 99,
			},
			expected: expected{
				val: 99,
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 99,
			},
			expected: expected{
				val: 99,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint8OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uintOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint
	}

	type expected struct {
		val uint
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 12,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uintOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint16OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint16
	}

	type expected struct {
		val uint16
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 12,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint16OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_uint32OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint32
	}

	type expected struct {
		val uint32
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "12",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 12,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
		{
			name: "malformed env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 999,
			},
			expected: expected{
				val: 999,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uint32OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_ipOrDefault(t *testing.T) {
	const rawDefault = "0.0.0.0"

	type args struct {
		key        string
		defaultVal net.IP
	}

	type expected struct {
		val net.IP
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "192.168.8.0",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getIP(t, rawDefault),
			},
			expected: expected{
				val: getIP(t, rawDefault),
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "192.168.8.0",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getIP(t, rawDefault),
			},
			expected: expected{
				val: getIP(t, "192.168.8.0"),
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "192.168.8.0ssss",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getIP(t, rawDefault),
			},
			expected: expected{
				val: getIP(t, rawDefault),
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: getIP(t, rawDefault),
			},
			expected: expected{
				val: getIP(t, rawDefault),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := ipOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_urlSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []url.URL
		separator  string
	}

	type expected struct {
		val []url.URL
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "https://google.com,https://github.com",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []url.URL{getURL(t, "https://bing.com")},
				separator:  ",",
			},
			expected: expected{
				val: []url.URL{getURL(t, "https://bing.com")},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "https://google.com,https://github.com",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []url.URL{getURL(t, "https://bing.com")},
				separator:  ",",
			},
			expected: expected{
				val: []url.URL{
					getURL(t, "https://google.com"),
					getURL(t, "https://github.com"),
				},
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "https://google.com,htps://%%2github.com",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []url.URL{getURL(t, "https://bing.com")},
				separator:  ",",
			},
			expected: expected{
				val: []url.URL{getURL(t, "https://bing.com")},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []url.URL{getURL(t, "https://bing.com")},
				separator:  ",",
			},
			expected: expected{
				val: []url.URL{getURL(t, "https://bing.com")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := urlSliceOrDefault(tt.args.key, tt.args.defaultVal, ",")
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_ipSliceOrDefault(t *testing.T) {
	const rawDefault = "0.0.0.0"

	type args struct {
		key        string
		defaultVal []net.IP
		separator  string
	}

	type expected struct {
		val []net.IP
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "192.168.8.0,2001:cb8::17",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []net.IP{getIP(t, rawDefault)},
				separator:  ",",
			},
			expected: expected{
				val: []net.IP{getIP(t, rawDefault)},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "192.168.8.0,2001:cb8::17",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []net.IP{getIP(t, rawDefault)},
				separator:  ",",
			},
			expected: expected{
				val: []net.IP{
					getIP(t, "192.168.8.0"),
					getIP(t, "2001:cb8::17"),
				},
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "192.168.8.0,sdsdsd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []net.IP{getIP(t, rawDefault)},
				separator:  ",",
			},
			expected: expected{
				val: []net.IP{getIP(t, rawDefault)},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []net.IP{getIP(t, rawDefault)},
				separator:  ",",
			},
			expected: expected{
				val: []net.IP{getIP(t, rawDefault)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := ipSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.separator)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_boolSliceOrDefault tests the boolSliceOrDefault function.
func Test_boolSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []bool
		separator  string
	}

	type expected struct {
		val []bool
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "true,false",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []bool{true},
				separator:  ",",
			},
			expected: expected{
				val: []bool{true},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,false",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []bool{false},
				separator:  ",",
			},
			expected: expected{
				val: []bool{true, false},
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []bool{false},
				separator:  ",",
			},
			expected: expected{
				val: []bool{false},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []bool{false},
				separator:  ",",
			},
			expected: expected{
				val: []bool{false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := boolSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.separator)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_uintptrOrDefault tests the uintptrOrDefault function.
func Test_uintptrOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uintptr
	}

	type expected struct {
		val uintptr
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "123",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 123,
			},
			expected: expected{
				val: 123,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "123",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 123,
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uintptrOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_uintptrSliceOrDefault tests the uintptrSliceOrDefault function.
func Test_uintptrSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uintptr
		separator  string
	}

	type expected struct {
		val []uintptr
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "123,456",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uintptr{123},
				separator:  ",",
			},
			expected: expected{
				val: []uintptr{123},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "123,456",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uintptr{456},
				separator:  ",",
			},
			expected: expected{
				val: []uintptr{123, 456},
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "123,asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []uintptr{456},
				separator:  ",",
			},
			expected: expected{
				val: []uintptr{456},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},

			args: args{
				key:        testEnvKey,
				defaultVal: []uintptr{456},
				separator:  ",",
			},
			expected: expected{
				val: []uintptr{456},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := uintptrSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.separator)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_complex64OrDefault tests the complex64OrDefault function.
func Test_complex64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal complex64
	}

	type expected struct {
		val complex64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 123,
			},
			expected: expected{
				val: 123,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 1 + 2i,
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := complex64OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_complex64SliceOrDefault tests the complex64SliceOrDefault function.
func Test_complex64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []complex64
		separator  string
	}

	type expected struct {
		val []complex64
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i),(3+4i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []complex64{123},
				separator:  ",",
			},
			expected: expected{
				val: []complex64{123},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i),(3+4i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []complex64{456},
				separator:  ",",
			},
			expected: expected{
				val: []complex64{1 + 2i, 3 + 4i},
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i),asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []complex64{456},
				separator:  ",",
			},
			expected: expected{
				val: []complex64{456},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},

			args: args{
				key:        testEnvKey,
				defaultVal: []complex64{456},
				separator:  ",",
			},
			expected: expected{
				val: []complex64{456},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := complex64SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.separator)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_complex128OrDefault tests the complex128OrDefault function.
func Test_complex128OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal complex128
	}

	type expected struct {
		val complex128
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 123,
			},
			expected: expected{
				val: 123,
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 1 + 2i,
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := complex128OrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_complex128SliceOrDefault tests the complex128SliceOrDefault function.
func Test_complex128SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []complex128
		separator  string
	}

	type expected struct {
		val []complex128
	}

	var tests = []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i),(3+4i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []complex128{123},
				separator:  ",",
			},
			expected: expected{
				val: []complex128{123},
			},
		},
		{
			name: "env set - env value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i),(3+4i)",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []complex128{456},
				separator:  ",",
			},
			expected: expected{
				val: []complex128{1 + 2i, 3 + 4i},
			},
		},
		{
			name: "env set, corrupted - default value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i),asd",
				},
			},
			args: args{
				key:        testEnvKey,
				defaultVal: []complex128{456},
				separator:  ",",
			},
			expected: expected{
				val: []complex128{456},
			},
		},
		{
			name: "empty env value set - default returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},

			args: args{
				key:        testEnvKey,
				defaultVal: []complex128{456},
				separator:  ",",
			},
			expected: expected{
				val: []complex128{456},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := complex128SliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.separator)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}
