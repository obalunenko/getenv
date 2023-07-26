package internal

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		_, err := getNumberSliceGen[[]float32](testEnvKey, ",")
		require.NoError(b, err)
	}
}

func Test_getNumberGenInt(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     int
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "128",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     128,
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[int](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

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
		val     string
		wantErr assert.ErrorAssertionFunc
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
				val: "",
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				val:     "newval",
				wantErr: assert.NoError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getString(tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenInt64(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     int64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     1024,
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[int64](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenInt8(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     int8
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     10,
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[int8](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenInt16(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     int16
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     1024,
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[int16](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenInt32(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     int32
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     1024,
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[int32](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenFloat32(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     float32
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: float32(0),
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     float32(1024.123),
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: float32(0),
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[float32](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenFloat64(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     float64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "newval",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     1024.123,
				wantErr: assert.NoError,
			},
		},
		{
			name: "invalid env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "128s7",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[float64](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

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
		val     bool
		wantErr assert.ErrorAssertionFunc
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
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				val:     true,
				wantErr: assert.NoError,
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
				val: false,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getBool(tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getStringSlice(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []string
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "true,newval",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []string{"true", "newval"},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,newval",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getStringSlice(tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenInt(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []int
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []int{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]int](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenFloat32(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []float32
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []float32{1.05, 2.07},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed data value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sss,sss",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]float32](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenFloat64(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []float64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []float64{1.05, 2.07},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed data value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sss,sss",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]float64](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenInt16(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []int16
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []int16{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]int16](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenInt32(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []int32
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []int32{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]int32](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenUint(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []uint
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []uint{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]uint](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenUint8(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []uint8
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []uint8{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]uint8](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenUint16(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []uint16
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []uint16{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) && assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]uint16](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenUint32(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []uint32
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []uint32{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]uint32](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenInt8(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []int8
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []int8{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]int8](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenInt64(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []int64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1.05,2.07",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []int64{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "sssss,999",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]int64](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getTime(t *testing.T) {
	const layout = "2006/02/01 15:04"

	type args struct {
		key    string
		layout string
	}

	type expected struct {
		val     time.Time
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "2018/21/04 22:30",
				},
			},
			args: args{
				key:    testEnvKey,
				layout: layout,
			},
			expected: expected{
				val: time.Time{},
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:    testEnvKey,
				layout: layout,
			},
			expected: expected{
				val:     time.Date(2018, 4, 21, 22, 30, 0, 0, time.UTC),
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:    testEnvKey,
				layout: layout,
			},
			expected: expected{
				val: time.Time{},
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "20222-sdslll",
				},
			},
			args: args{
				key:    testEnvKey,
				layout: layout,
			},
			expected: expected{
				val: time.Time{},
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getTime(tt.args.key, tt.args.layout)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getURL(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     url.URL
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "postgres://user:pass@host.com:5432/path?k=v#f",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: url.URL{},
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     getTestURL(t, "postgres://user:pass@host.com:5432/path?k=v#f"),
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "postgres://user:pass@host.com:5432/path?k=v#f%%2",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: url.URL{},
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: url.URL{},
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getURL(tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getTimeSlice(t *testing.T) {
	const layout = "2006/02/01 15:04"

	type args struct {
		key       string
		layout    string
		separator string
	}

	type expected struct {
		val     []time.Time
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "2018/21/04 22:30,2023/21/04 22:30",
				},
			},
			args: args{
				key:       testEnvKey,
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
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
				key:       testEnvKey,
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: []time.Time{
					time.Date(2018, 4, 21, 22, 30, 0, 0, time.UTC),
					time.Date(2023, 4, 21, 22, 30, 0, 0, time.UTC),
				},
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key:       testEnvKey,
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "20222-sdslll",
				},
			},
			args: args{
				key:       testEnvKey,
				layout:    layout,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getTimeSlice(tt.args.key, tt.args.layout, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getDurationSlice(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []time.Duration
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "2m,3h",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: []time.Duration{
					2 * time.Minute, 3 * time.Hour,
				},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2m,3hddd",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getDurationSlice(tt.args.key, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_duration(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     time.Duration
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12s",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: time.Duration(0),
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     time.Second * 12,
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: time.Duration(0),
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "yyydd88",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: time.Duration(0),
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getDuration(tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenUint64(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     uint64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     12,
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[uint64](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenUint64(t *testing.T) {
	type args struct {
		key string
		sep string
	}

	type expected struct {
		val     []uint64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "1,27",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val:     []uint64{1, 2},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, no separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "env set, wrong separator - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "1,2",
				},
			},
			args: args{
				key: testEnvKey,
				sep: "|",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Error(t, err) &&
						assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
				sep: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]uint64](tt.args.key, tt.args.sep)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenUint8(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     uint8
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     12,
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[uint8](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenUint(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     uint
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     12,
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[uint](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenUint16(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     uint16
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     12,
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[uint16](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenUint32(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     uint32
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "12",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     12,
				wantErr: assert.NoError,
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "malformed env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "iii99",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[uint32](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getIP(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     net.IP
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "192.168.8.0",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     getTestIP(t, "192.168.8.0"),
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "192.168.8.0ssss",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
		{
			name: "env ipv6 set - value returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "2001:db8::68",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val:     getTestIP(t, "2001:db8::68"),
				wantErr: assert.NoError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getIP(tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getURLSlice(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []url.URL
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "https://google.com,https://github.com",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: []url.URL{
					getTestURL(t, "https://google.com"),
					getTestURL(t, "https://github.com"),
				},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "https://google.com,htps://%%2github.com",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getURLSlice(tt.args.key, ",")
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getIPSlice(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []net.IP
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "192.168.8.0,2001:cb8::17",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: []net.IP{
					getTestIP(t, "192.168.8.0"),
					getTestIP(t, "2001:cb8::17"),
				},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "192.168.8.0,sdsdsd",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getIPSlice(tt.args.key, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

// Test_getBoolSlice tests the getBoolSlice function.
func Test_getBoolSlice(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []bool
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "true,false",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:     []bool{true, false},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "true,asd",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Equal(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getBoolSlice(tt.args.key, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberGenUintptr(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     uintptr
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
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
				val: 0,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     123,
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err value returned",
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
				val: 0,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberGen[uintptr](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getNumberSliceGenuintptr(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []uintptr
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "123,456",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:     []uintptr{123, 456},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "123,asd",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Equal(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getNumberSliceGen[[]uintptr](tt.args.key, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getComplexGenComplex64(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     complex64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i)",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     1 + 2i,
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
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
				val: 0,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getComplexGen[complex64](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getComplexSliceGenComplex64(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []complex64
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i),(3+4i)",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Equal(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:     []complex64{1 + 2i, 3 + 4i},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i),asd",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Equal(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getComplexSliceGen[[]complex64](tt.args.key, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getComplexGenComplex128(t *testing.T) {
	type args struct {
		key string
	}

	type expected struct {
		val     complex128
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i)",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.Equal(t, err, ErrNotSet)
				},
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
				key: testEnvKey,
			},
			expected: expected{
				val:     1 + 2i,
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
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
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "",
				},
			},
			args: args{
				key: testEnvKey,
			},
			expected: expected{
				val: 0,
				wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getComplexGen[complex128](tt.args.key)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}

func Test_getComplexSliceGenComplex128(t *testing.T) {
	type args struct {
		key       string
		separator string
	}

	type expected struct {
		val     []complex128
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name     string
		precond  precondition
		args     args
		expected expected
	}{
		{
			name: "env not set - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: false,
					val:   "(1+2i),(3+4i)",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
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
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val:     []complex128{1 + 2i, 3 + 4i},
				wantErr: assert.NoError,
			},
		},
		{
			name: "env set, corrupted - err returned",
			precond: precondition{
				setenv: setenv{
					isSet: true,
					val:   "(1+2i),asd",
				},
			},
			args: args{
				key:       testEnvKey,
				separator: ",",
			},
			expected: expected{
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrInvalidValue)
				},
			},
		},
		{
			name: "empty env value set - err returned",
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
				val: nil,
				wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
					return assert.ErrorIs(t, err, ErrNotSet)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got, err := getComplexSliceGen[[]complex128](tt.args.key, tt.args.separator)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.val, got)
		})
	}
}
