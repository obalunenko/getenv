package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
