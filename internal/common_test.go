package internal

import (
	"net"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type panicAssertionFunc func(t assert.TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool

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

func getURL(tb testing.TB, rawURL string) url.URL {
	tb.Helper()

	val, err := url.Parse(rawURL)
	require.NoError(tb, err)

	return *val
}

func getIP(tb testing.TB, raw string) net.IP {
	tb.Helper()

	return net.ParseIP(raw)
}
