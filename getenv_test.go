package getenv_test

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/getenv"
	"github.com/obalunenko/getenv/option"
)

const testEnvKey = "GH_GETENV_TEST"

type setenv struct {
	isSet bool
	val   string
}

type precondition struct {
	setenv setenv
}

func (p precondition) maybeSetEnv(tb testing.TB, key string) {
	if p.setenv.isSet {
		tb.Setenv(key, p.setenv.val)
	}
}

func BenchmarkEnvOrDefault(b *testing.B) {
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
		_ = getenv.EnvOrDefault(testEnvKey, []float64{}, option.WithSeparator(","))
	}
}

func TestIntOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}

	type expected struct {
		val int
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestStringOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}

	type expected struct {
		val string
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int64
	}

	type expected struct {
		val int64
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt32OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int32
	}

	type expected struct {
		val int32
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestFloat32OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal float32
	}

	type expected struct {
		val float32
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestFloat64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal float64
	}

	type expected struct {
		val float64
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestBoolOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}

	type expected struct {
		val bool
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestStringSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []string
		sep        string
	}

	type expected struct {
		val []string
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt32SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int32
		sep        string
	}

	type expected struct {
		val []int32
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestIntSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int
		sep        string
	}

	type expected struct {
		val []int
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestFloat32SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []float32
		sep        string
	}

	type expected struct {
		val []float32
	}

	tests := []struct {
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
			},
			expected: expected{
				val: []float32{-99.99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestFloat64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []float64
		sep        string
	}

	type expected struct {
		val []float64
	}

	tests := []struct {
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
			},
			expected: expected{
				val: []float64{-99.99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int64
		sep        string
	}

	type expected struct {
		val []int64
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestTimeOrDefault(t *testing.T) {
	const layout = "2006/02/01 15:04"

	type args struct {
		key        string
		defaultVal time.Time
		layout     string
	}

	type expected struct {
		val time.Time
	}

	tests := []struct {
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
				defaultVal: time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC),
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
				defaultVal: time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2018, 0o4, 21, 22, 30, 0, 0, time.UTC),
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
				defaultVal: time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC),
				layout:     layout,
			},
			expected: expected{
				val: time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithTimeLayout(tt.args.layout))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestTimeSliceOrDefault(t *testing.T) {
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

	tests := []struct {
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
				key:        testEnvKey,
				defaultVal: []time.Time{time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC)},
				layout:     layout,
				separator:  ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC),
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
				key:        testEnvKey,
				defaultVal: []time.Time{time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC)},
				layout:     layout,
				separator:  ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2018, 0o4, 21, 22, 30, 0, 0, time.UTC),
					time.Date(2023, 0o4, 21, 22, 30, 0, 0, time.UTC),
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
				key:        testEnvKey,
				defaultVal: []time.Time{time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC)},
				layout:     layout,
				separator:  ",",
			},
			expected: expected{
				val: []time.Time{time.Date(2021, 0o4, 21, 22, 30, 0, 0, time.UTC)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithTimeLayout(tt.args.layout), option.WithSeparator(","))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestDurationSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []time.Duration
		separator  string
	}

	type expected struct {
		val []time.Duration
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(","))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestDurationOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal time.Duration
	}

	type expected struct {
		val time.Duration
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint8OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint8
	}

	type expected struct {
		val uint8
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint16OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint16
	}

	type expected struct {
		val uint16
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint64
	}

	type expected struct {
		val uint64
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt8OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int8
	}

	type expected struct {
		val int8
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt8SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int8
		sep        string
	}

	type expected struct {
		val []int8
	}

	tests := []struct {
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
				defaultVal: []int8{99},
				sep:        ",",
			},
			expected: expected{
				val: []int8{99},
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
				defaultVal: []int8{99},
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
				defaultVal: []int8{99},
				sep:        "",
			},
			expected: expected{
				val: []int8{99},
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
				defaultVal: []int8{99},
				sep:        "|",
			},
			expected: expected{
				val: []int8{99},
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
				defaultVal: []int8{99},
			},
			expected: expected{
				val: []int8{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint64
		sep        string
	}

	type expected struct {
		val []uint64
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUintOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint
	}

	type expected struct {
		val uint
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUintSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint
		sep        string
	}

	type expected struct {
		val []uint
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint32SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint32
		sep        string
	}

	type expected struct {
		val []uint32
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint32OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint32
	}

	type expected struct {
		val uint32
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt16OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal int16
	}

	type expected struct {
		val int16
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint8SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint8
		sep        string
	}

	type expected struct {
		val []uint8
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestUint16SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uint16
		sep        string
	}

	type expected struct {
		val []uint16
	}

	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestInt16SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []int16
		sep        string
	}

	type expected struct {
		val []int16
	}

	tests := []struct {
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
				defaultVal: []int16{99},
				sep:        ",",
			},
			expected: expected{
				val: []int16{99},
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
				defaultVal: []int16{99},
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
				defaultVal: []int16{99},
				sep:        "",
			},
			expected: expected{
				val: []int16{99},
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
				defaultVal: []int16{99},
				sep:        "|",
			},
			expected: expected{
				val: []int16{99},
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
				defaultVal: []int16{99},
			},
			expected: expected{
				val: []int16{99},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.sep))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func getURL(tb testing.TB, rawURL string) url.URL {
	tb.Helper()

	val, err := url.Parse(rawURL)
	require.NoError(tb, err)

	return *val
}

func TestURLOrDefault(t *testing.T) {
	const rawDefault = "https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"

	type args struct {
		key        string
		defaultVal url.URL
	}

	type expected struct {
		val url.URL
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func getIP(tb testing.TB, raw string) net.IP {
	tb.Helper()

	return net.ParseIP(raw)
}

func TestIPOrDefault(t *testing.T) {
	const rawDefault = "0.0.0.0"

	type args struct {
		key        string
		defaultVal net.IP
	}

	type expected struct {
		val net.IP
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestIPSliceOrDefault(t *testing.T) {
	const rawDefault = "0.0.0.0"

	type args struct {
		key        string
		defaultVal []net.IP
		separator  string
	}

	type expected struct {
		val []net.IP
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.separator))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func TestURLSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []url.URL
		separator  string
	}

	type expected struct {
		val []url.URL
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(","))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestBoolSliceOrDefault tests the EnvOrDefault function with a bool slice.
func TestBoolSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []bool
		separator  string
	}

	type expected struct {
		val []bool
	}

	tests := []struct {
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
				defaultVal: []bool{false, true},
				separator:  ",",
			},
			expected: expected{
				val: []bool{false, true},
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
				defaultVal: []bool{false, true},
				separator:  ",",
			},
			expected: expected{
				val: []bool{true, false},
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
				defaultVal: []bool{false, true},
				separator:  ",",
			},
			expected: expected{
				val: []bool{false, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.separator))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestUintptrOrDefault tests the EnvOrDefault function with a uintptr.
func TestUintptrOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal uintptr
	}

	type expected struct {
		val uintptr
	}

	tests := []struct {
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
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestUintptrSliceOrDefault tests the EnvOrDefault function with a uintptr slice.
func TestUintptrSliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []uintptr
		separator  string
	}

	type expected struct {
		val []uintptr
	}

	tests := []struct {
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.separator))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestComplex64OrDefault tests the EnvOrDefault function with a complex64.
func TestComplex64OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal complex64
	}

	type expected struct {
		val complex64
	}

	tests := []struct {
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
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestComplex128OrDefault tests the EnvOrDefault function with a complex128.
func TestComplex128OrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal complex128
	}

	type expected struct {
		val complex128
	}

	tests := []struct {
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
				defaultVal: 456,
			},
			expected: expected{
				val: 456,
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestComplex64SliceOrDefault tests the EnvOrDefault function with a []complex64.
func TestComplex64SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []complex64
		separator  string
	}

	type expected struct {
		val []complex64
	}

	tests := []struct {
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
				defaultVal: []complex64{123, 456},
				separator:  ",",
			},
			expected: expected{
				val: []complex64{123, 456},
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
				defaultVal: []complex64{123, 456},
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.separator))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestComplex128SliceOrDefault tests the EnvOrDefault function with a []complex128.
func TestComplex128SliceOrDefault(t *testing.T) {
	type args struct {
		key        string
		defaultVal []complex128
		separator  string
	}

	type expected struct {
		val []complex128
	}

	tests := []struct {
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
				defaultVal: []complex128{123, 456},
				separator:  ",",
			},
			expected: expected{
				val: []complex128{123, 456},
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
				defaultVal: []complex128{123, 456},
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

			got := getenv.EnvOrDefault(tt.args.key, tt.args.defaultVal, option.WithSeparator(tt.args.separator))
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestEnvInt tests the Env function with an int.
func TestEnvInt(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val       int
		wantError assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "123",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val:       0,
				wantError: errorEqual(getenv.ErrNotSet),
			},
		},
		{
			name: "env set",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "123",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val:       123,
				wantError: assert.NoError,
			},
		},
		{
			name: "env set, corrupted",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "asd",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val:       0,
				wantError: errorEqual(getenv.ErrInvalidValue),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getenv.Env[int](tt.args.key)
			if !tt.expected.wantError(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// TestEnvIntSlice tests the Env function with a []int.
func TestEnvIntSlice(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val       []int
		wantError assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1,2,3",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:       nil,
				wantError: errorEqual(getenv.ErrNotSet),
			},
		},
		{
			name: "env set",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2,3",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:       []int{1, 2, 3},
				wantError: assert.NoError,
			},
		},
		{
			name: "env set, corrupted",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,asd,3",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:       nil,
				wantError: errorEqual(getenv.ErrInvalidValue),
			},
		},
		{
			name: "empty env value set",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:       nil,
				wantError: errorEqual(getenv.ErrNotSet),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getenv.Env[[]int](tt.args.key, option.WithSeparator(tt.args.separator))
			if !tt.expected.wantError(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func errorEqual(expected error) assert.ErrorAssertionFunc {
	return func(t assert.TestingT, err error, i ...interface{}) bool {
		return assert.Error(t, err, i...) &&
			assert.ErrorContains(t, err, expected.Error(), i...)
	}
}
