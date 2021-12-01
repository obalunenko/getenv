package getenv_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/getenv"
)

const testEnvKey = "GH_GETENV_TEST"

type setenv struct {
	isSet bool
	val   string
}

type precond struct {
	setenv setenv
}

func (p precond) maybeSetEnv(tb testing.TB, key string) {
	if p.setenv.isSet {
		tb.Setenv(key, p.setenv.val)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			precond: precond{
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

			got := getenv.IntOrDefault(tt.args.key, tt.args.defaultVal)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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

			got := getenv.StringOrDefault(tt.args.key, tt.args.defaultVal)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			precond: precond{
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

			got := getenv.Int64OrDefault(tt.args.key, tt.args.defaultVal)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			precond: precond{
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

			got := getenv.Float64OrDefault(tt.args.key, tt.args.defaultVal)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			precond: precond{
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

			got := getenv.BoolOrDefault(tt.args.key, tt.args.defaultVal)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			name: "empty env value set - default returned",
			precond: precond{
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

			got := getenv.StringSliceOrDefault(tt.args.key, tt.args.defaultVal, tt.args.sep)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			precond: precond{
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precond.maybeSetEnv(t, tt.args.key)

			got := getenv.TimeOrDefault(tt.args.key, tt.args.defaultVal, tt.args.layout)
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

	var tests = []struct {
		name     string
		precond  precond
		args     args
		expected expected
	}{
		{
			name: "env not set - default returned",
			precond: precond{
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
			precond: precond{
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
			precond: precond{
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

			got := getenv.DurationOrDefault(tt.args.key, tt.args.defaultVal)
			assert.Equal(t, tt.expected.val, got)
		})
	}
}
