package getenv_test

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/obalunenko/getenv"
	"github.com/obalunenko/getenv/option"
)

// Level is a custom type: it implements getenv.EnvParser.
type Level int

// Level values. The zero value is not a valid level.
const (
	LevelDebug Level = iota + 1
	LevelInfo
	LevelWarn
	LevelError
)

// Check at compile time that Level implements getenv.EnvParser.
var _ getenv.EnvParser[Level] = Level(0)

// String returns the name of the level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	default:
		return fmt.Sprintf("Level(%d)", int(l))
	}
}

// ParseEnv parses Level from the value of an environment variable.
// It has a value receiver: getenv calls it on the zero value.
func (Level) ParseEnv(raw string) (Level, error) {
	switch strings.ToLower(raw) {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	default:
		return 0, fmt.Errorf("unknown level %q", raw)
	}
}

// Port is a custom type over a basic type: it is parsed as uint16.
type Port uint16

// This example shows how to read custom types from the environment with the same Env and EnvOrDefault:
// types that implement getenv.EnvParser or encoding.TextUnmarshaler, types over basic types and slices of them.
func Example_customType() {
	const (
		key = "GH_GETENV_TEST"
	)

	var (
		lvl  Level
		lvls []Level
		err  error
	)

	defer func() {
		if err = os.Unsetenv(key); err != nil {
			panic(err)
		}
	}()

	// Level
	if err = os.Setenv(key, "warn"); err != nil {
		panic(err)
	}

	lvl, err = getenv.Env[Level](key)
	fmt.Printf("[%T]: %v; err: %v\n", lvl, lvl, err)

	// []Level
	if err = os.Setenv(key, "debug|error"); err != nil {
		panic(err)
	}

	lvls, err = getenv.Env[[]Level](key, option.WithSeparator("|"))
	fmt.Printf("[%T]: %v; err: %v\n", lvls, lvls, err)

	// Level parse error
	if err = os.Setenv(key, "verbose"); err != nil {
		panic(err)
	}

	lvl, err = getenv.Env[Level](key)
	fmt.Printf("[%T]: %v; err: %v\n", lvl, lvl, err)
	fmt.Printf("is invalid value: %t\n", errors.Is(err, getenv.ErrInvalidValue))

	// Level default
	lvl = getenv.EnvOrDefault(key, LevelInfo)
	fmt.Printf("[%T]: %v\n", lvl, lvl)

	// slog.Level implements encoding.TextUnmarshaler
	if err = os.Setenv(key, "WARN,ERROR"); err != nil {
		panic(err)
	}

	slvls, err := getenv.Env[[]slog.Level](key)
	fmt.Printf("[%T]: %v; err: %v\n", slvls, slvls, err)

	// Port
	if err = os.Setenv(key, "8080"); err != nil {
		panic(err)
	}

	port := getenv.EnvOrDefault(key, Port(80))
	fmt.Printf("[%T]: %v\n", port, port)

	// Output:
	// [getenv_test.Level]: warn; err: <nil>
	// [[]getenv_test.Level]: [debug error]; err: <nil>
	// [getenv_test.Level]: Level(0); err: failed to parse environment variable[GH_GETENV_TEST]: unknown level "verbose": invalid value
	// is invalid value: true
	// [getenv_test.Level]: info
	// [[]slog.Level]: [WARN ERROR]; err: <nil>
	// [getenv_test.Port]: 8080
}
