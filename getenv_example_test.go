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
	const (
		key = "GH_GETENV_TEST"
	)

	defer func() {
		if err := os.Unsetenv(key); err != nil {
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

	// []string
	if err := os.Setenv(key, "a,b,c,d"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, []string{}, option.WithSeparator(","))
	fmt.Printf("[%T]: %v\n", val, val)

	// complex128
	if err := os.Setenv(key, "1+2i"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, complex128(0))
	fmt.Printf("[%T]: %v\n", val, val)

	// []complex64
	if err := os.Setenv(key, "1+2i,3+4i"); err != nil {
		panic(err)
	}

	val = getenv.EnvOrDefault(key, []complex64{}, option.WithSeparator(","))
	fmt.Printf("[%T]: %v\n", val, val)

	// Output:
	// [string]: golly
	// [int]: 123
	// [time.Time]: 2022-01-20 00:00:00 +0000 UTC
	// [[]float64]: [26.89 0.67]
	// [time.Duration]: 2h35m0s
	// [url.URL]: {https  test:abcd123 golangbyexample.com:8000 /tutorials/intro  false false type=advance&compact=false history }
	// [net.IP]: 2001:cb8::17
	// [[]string]: [a b c d]
	// [complex128]: (1+2i)
	// [[]complex64]: [(1+2i) (3+4i)]
}

func ExampleEnv() {
	const (
		key = "GH_GETENV_TEST"
	)

	var (
		val any
		err error
	)

	defer func() {
		if err = os.Unsetenv(key); err != nil {
			panic(err)
		}
	}()

	// string
	if err = os.Setenv(key, "golly"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[string](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// int
	if err = os.Setenv(key, "123"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[int](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// int conversion error
	if err = os.Setenv(key, "123s4"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[int](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// time.Time
	if err = os.Setenv(key, "2022-01-20"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[time.Time](key, option.WithTimeLayout("2006-01-02"))
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// []float64
	if err = os.Setenv(key, "26.89,0.67"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[[]float64](key, option.WithSeparator(","))
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// time.Duration
	if err = os.Setenv(key, "2h35m"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[time.Duration](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// url.URL
	if err = os.Setenv(key, "https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[url.URL](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// net.IP
	if err = os.Setenv(key, "2001:cb8::17"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[net.IP](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// []string
	if err = os.Setenv(key, "a,b,c,d"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[[]string](key, option.WithSeparator(","))
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// complex128
	if err = os.Setenv(key, "1+2i"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[complex128](key)
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// []complex64
	if err = os.Setenv(key, "1+2i,3+4i"); err != nil {
		panic(err)
	}

	val, err = getenv.Env[[]complex64](key, option.WithSeparator(","))
	fmt.Printf("[%T]: %v; err: %v\n", val, val, err)

	// Output:
	// [string]: golly; err: <nil>
	// [int]: 123; err: <nil>
	// [int]: 0; err: failed to parse environment variable[GH_GETENV_TEST]: invalid value
	// [time.Time]: 2022-01-20 00:00:00 +0000 UTC; err: <nil>
	// [[]float64]: [26.89 0.67]; err: <nil>
	// [time.Duration]: 2h35m0s; err: <nil>
	// [url.URL]: {https  test:abcd123 golangbyexample.com:8000 /tutorials/intro  false false type=advance&compact=false history }; err: <nil>
	// [net.IP]: 2001:cb8::17; err: <nil>
	// [[]string]: [a b c d]; err: <nil>
	// [complex128]: (1+2i); err: <nil>
	// [[]complex64]: [(1+2i) (3+4i)]; err: <nil>
}
