package option

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/getenv"
)

// The options must keep working for the callers that have not moved to package getenv yet.
func TestOptions(t *testing.T) {
	const key = "GH_GETENV_TEST"

	t.Setenv(key, "2022-01-20|2022-01-21")

	got, err := getenv.Env[[]time.Time](key, WithSeparator("|"), WithTimeLayout(time.DateOnly))
	require.NoError(t, err)
	assert.Equal(t, []time.Time{
		time.Date(2022, 1, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 1, 21, 0, 0, 0, 0, time.UTC),
	}, got)
}
