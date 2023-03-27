package getenv_test

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/obalunenko/getenv"
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

	val = getenv.EnvOrDefault(key, "golly")
	fmt.Printf("[%T]: %v\n", val, val)

	// int
	if err := os.Setenv(key, "123"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, -99)
	fmt.Printf("[%T]: %v\n", val, val)

	// time.Time
	if err := os.Setenv(key, "2022-01-20"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key,
		time.Date(1992, 12, 1, 0, 0, 0, 0, time.UTC),
		option.WithTimeLayout("2006-01-02"),
	)
	fmt.Printf("[%T]: %v\n", val, val)

	// []float64
	if err := os.Setenv(key, "26.89,0.67"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, []float64{-99},
		option.WithSeparator(","),
	)
	fmt.Printf("[%T]: %v\n", val, val)

	// time.Duration
	if err := os.Setenv(key, "2h35m"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, time.Second)
	fmt.Printf("[%T]: %v\n", val, val)

	// url.URL
	if err := os.Setenv(key, "https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, url.URL{})
	fmt.Printf("[%T]: %v\n", val, val)

	// net.IP
	if err := os.Setenv(key, "2001:cb8::17"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, net.IP{})
	fmt.Printf("[%T]: %v\n", val, val)

	// Output:
	// [string]: golly
	// [int]: 123
	// [time.Time]: 2022-01-20 00:00:00 +0000 UTC
	// [[]float64]: [26.89 0.67]
	// [time.Duration]: 2h35m0s
	// [url.URL]: {https  test:abcd123 golangbyexample.com:8000 /tutorials/intro  false false type=advance&compact=false history }
	// [net.IP]: 2001:cb8::17
}
