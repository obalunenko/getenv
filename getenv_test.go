package getenv_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/getenv"
)

// Parsing of every supported type is covered in parse_test.go;
// these tests check the public API: lookup, errors and options.

const testEnvKey = "GH_GETENV_TEST"

func TestEnv(t *testing.T) {
	tests := []struct {
		name    string
		set     bool
		env     string
		want    int
		wantErr error
	}{
		{name: "not set", wantErr: getenv.ErrNotSet},
		{name: "empty value is not set", set: true, env: "", wantErr: getenv.ErrNotSet},
		{name: "valid", set: true, env: "123", want: 123},
		{name: "invalid", set: true, env: "asd", wantErr: getenv.ErrInvalidValue},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(testEnvKey, tt.env)
			}

			got, err := getenv.Env[int](testEnvKey)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.ErrorContains(t, err, testEnvKey)
				assert.Zero(t, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEnvErrorSentinels(t *testing.T) {
	t.Setenv(testEnvKey, "not-an-int")

	_, err := getenv.Env[int](testEnvKey)
	require.ErrorIs(t, err, getenv.ErrInvalidValue)
	require.ErrorIs(t, err, strconv.ErrSyntax)
	assert.NotErrorIs(t, err, getenv.ErrNotSet)

	_, err = getenv.Env[int]("GH_GETENV_TEST_NOT_SET")
	require.ErrorIs(t, err, getenv.ErrNotSet)
	assert.NotErrorIs(t, err, getenv.ErrInvalidValue)
}

func TestEnvOptions(t *testing.T) {
	t.Run("separator", func(t *testing.T) {
		t.Setenv(testEnvKey, "1|2|3")

		got, err := getenv.Env[[]int](testEnvKey, getenv.WithSeparator("|"))
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, got)
	})

	t.Run("time layout", func(t *testing.T) {
		t.Setenv(testEnvKey, "2022-01-20")

		got, err := getenv.Env[time.Time](testEnvKey, getenv.WithTimeLayout(time.DateOnly))
		require.NoError(t, err)
		assert.Equal(t, time.Date(2022, 1, 20, 0, 0, 0, 0, time.UTC), got)
	})

	t.Run("default separator", func(t *testing.T) {
		t.Setenv(testEnvKey, "1,2")

		got, err := getenv.Env[[]int](testEnvKey)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2}, got)
	})

	t.Run("default time layout", func(t *testing.T) {
		t.Setenv(testEnvKey, "2022-01-20T10:30:00Z")

		got, err := getenv.Env[time.Time](testEnvKey)
		require.NoError(t, err)
		assert.Equal(t, time.Date(2022, 1, 20, 10, 30, 0, 0, time.UTC), got)
	})

	t.Run("empty separator", func(t *testing.T) {
		t.Setenv(testEnvKey, "1,2")

		_, err := getenv.Env[[]int](testEnvKey, getenv.WithSeparator(""))
		require.ErrorIs(t, err, getenv.ErrInvalidValue)
		assert.ErrorContains(t, err, "empty separator")
	})

	t.Run("not set is reported before empty separator", func(t *testing.T) {
		_, err := getenv.Env[[]int](testEnvKey, getenv.WithSeparator(""))
		require.ErrorIs(t, err, getenv.ErrNotSet)
	})
}

func TestEnvOrDefault(t *testing.T) {
	const defaultVal = 42

	tests := []struct {
		name string
		set  bool
		env  string
		want int
	}{
		{name: "not set - default returned", want: defaultVal},
		{name: "empty value - default returned", set: true, env: "", want: defaultVal},
		{name: "valid - value returned", set: true, env: "128", want: 128},
		{name: "invalid - default returned", set: true, env: "128s7", want: defaultVal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(testEnvKey, tt.env)
			}

			assert.Equal(t, tt.want, getenv.EnvOrDefault(testEnvKey, defaultVal))
		})
	}
}

func TestEnvOrDefaultOptions(t *testing.T) {
	t.Run("separator", func(t *testing.T) {
		t.Setenv(testEnvKey, "a|b")

		got := getenv.EnvOrDefault(testEnvKey, []string{"default"}, getenv.WithSeparator("|"))
		assert.Equal(t, []string{"a", "b"}, got)
	})

	t.Run("default separator", func(t *testing.T) {
		t.Setenv(testEnvKey, "a,b")

		got := getenv.EnvOrDefault(testEnvKey, []string{"default"})
		assert.Equal(t, []string{"a", "b"}, got)
	})
}

// mustEnv is a caller-side generic helper: it compiles only because Parsable is exported.
func mustEnv[T getenv.Parsable](tb testing.TB, key string) T {
	tb.Helper()

	val, err := getenv.Env[T](key)
	require.NoError(tb, err)

	return val
}

func TestParsable(t *testing.T) {
	t.Setenv(testEnvKey, "2h35m")

	assert.Equal(t, 2*time.Hour+35*time.Minute, mustEnv[time.Duration](t, testEnvKey))
}

func BenchmarkEnvOrDefault(b *testing.B) {
	b.Setenv(testEnvKey, "1235.67,87.98")
	b.ReportAllocs()

	for b.Loop() {
		_ = getenv.EnvOrDefault(testEnvKey, []float64{}, getenv.WithSeparator(","))
	}
}
