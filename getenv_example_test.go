package getenv

import (
	"fmt"
	"os"
	"time"

	"github.com/obalunenko/getenv/option"
)

func ExampleEnvOrDefault() {
	key := "GH_GETENV_TEST"

	defer func() {
		if err := os.Unsetenv("GH_GETENV_TEST"); err != nil {
			panic(err)
		}
	}()

	var val any

	// string
	if err := os.Setenv(key, "golly"); err != nil {
		panic(err)
	}

	// var valS string
	val = EnvOrDefault(key, "golly")
	fmt.Printf("[%T]: %v\n", val, val)

	// int
	if err := os.Setenv(key, "123"); err != nil {
		panic(err)
	}

	val = EnvOrDefault(key, -99)
	fmt.Printf("[%T]: %v\n", val, val)

	// time.Time
	if err := os.Setenv(key, "2022-01-20"); err != nil {
		panic(err)
	}

	val = EnvOrDefault(key,
		time.Date(1992, 12, 1, 0, 0, 0, 0, time.UTC),
		option.WithTimeLayout("2006-01-02"),
	)
	fmt.Printf("[%T]: %v\n", val, val)

	// []float64
	if err := os.Setenv(key, "26.89,0.67"); err != nil {
		panic(err)
	}

	val = EnvOrDefault(key, []float64{-99},
		option.WithSeparator(","),
	)
	fmt.Printf("[%T]: %v\n", val, val)

	// time.Duration
	if err := os.Setenv(key, "2h35m"); err != nil {
		panic(err)
	}

	val = EnvOrDefault(key, time.Second)
	fmt.Printf("[%T]: %v\n", val, val)

	// Output:
	// [string]: golly
	// [int]: 123
	// [time.Time]: 2022-01-20 00:00:00 +0000 UTC
	// [[]float64]: [26.89 0.67]
	// [time.Duration]: 2h35m0s
}
