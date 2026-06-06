package internal

import (
	"net"
	"net/netip"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// panicAssertionFunc is a function that asserts that a function panics.
type panicAssertionFunc func(t assert.TestingT, f assert.PanicTestFunc, msgAndArgs ...any) bool

// testEnvKey is a key for test environment variable.
const testEnvKey = "GH_GETENV_TEST"

// setenv is a helper struct for setting environment variables.
type setenv struct {
	isSet bool
	val   string
}

// precondition is a helper struct for setting environment variables.
type precondition struct {
	setenv setenv
}

// maybeSetEnv sets environment variable if it is set.
func (p precondition) maybeSetEnv(tb testing.TB, key string) {
	if p.setenv.isSet {
		tb.Setenv(key, p.setenv.val)
	}
}

// getTestURL is a helper function for getting url.URL from string.
func getTestURL(tb testing.TB, rawURL string) url.URL {
	tb.Helper()

	val, err := url.Parse(rawURL)
	require.NoError(tb, err)

	return *val
}

// getTestIP is a helper function for getting net.IP from string.
func getTestIP(tb testing.TB, raw string) net.IP {
	tb.Helper()

	return net.ParseIP(raw)
}

// getTestNetIPAddr is a helper function for getting netip.Addr from string.
func getTestNetIPAddr(tb testing.TB, raw string) netip.Addr {
	tb.Helper()

	return netip.MustParseAddr(raw)
}

// getTestNetIPPrefix is a helper function for getting netip.Prefix from string.
func getTestNetIPPrefix(tb testing.TB, raw string) netip.Prefix {
	tb.Helper()

	return netip.MustParsePrefix(raw)
}

// getTestHardwareAddr is a helper function for getting net.HardwareAddr from string.
func getTestHardwareAddr(tb testing.TB, raw string) net.HardwareAddr {
	tb.Helper()

	val, err := net.ParseMAC(raw)
	require.NoError(tb, err)

	return val
}

func errorEqual(tb testing.TB, expected error) assert.ErrorAssertionFunc {
	tb.Helper()

	return func(at assert.TestingT, err error, i ...any) bool {
		return assert.Error(at, err, i...) &&
			assert.ErrorIs(at, err, expected, i...) &&
			assert.ErrorContains(at, err, expected.Error(), i...)
	}
}
