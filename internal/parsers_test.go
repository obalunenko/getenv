package internal

import (
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
