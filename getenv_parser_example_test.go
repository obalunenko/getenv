package getenv_test

import (
	"fmt"
	"os"
	"time"

	"github.com/obalunenko/getenv"
)

// Timeout is defined over time.Duration: without ParseEnv it would be parsed as int64 nanoseconds.
type Timeout time.Duration

// Check at compile time that Timeout implements getenv.EnvParser.
var _ getenv.EnvParser[Timeout] = Timeout(0)

// ParseEnv parses Timeout as a duration.
func (Timeout) ParseEnv(raw string) (Timeout, error) {
	d, err := time.ParseDuration(raw)

	return Timeout(d), err
}

// String returns Timeout as a duration.
func (t Timeout) String() string {
	return time.Duration(t).String()
}

// This example shows how to parse a type defined over time.Duration as a duration.
func ExampleEnvParser() {
	const (
		key = "GH_GETENV_TEST"
	)

	var err error

	defer func() {
		if err = os.Unsetenv(key); err != nil {
			panic(err)
		}
	}()

	// Timeout
	if err = os.Setenv(key, "2m30s"); err != nil {
		panic(err)
	}

	timeout, err := getenv.Env[Timeout](key)
	fmt.Printf("[%T]: %v; err: %v\n", timeout, timeout, err)

	// []Timeout
	if err = os.Setenv(key, "1s,500ms"); err != nil {
		panic(err)
	}

	timeouts := getenv.EnvOrDefault(key, []Timeout{Timeout(time.Second)})
	fmt.Printf("[%T]: %v\n", timeouts, timeouts)

	// Output:
	// [getenv_test.Timeout]: 2m30s; err: <nil>
	// [[]getenv_test.Timeout]: [1s 500ms]
}
