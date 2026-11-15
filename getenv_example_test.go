package getenv_test

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/obalunenko/getenv"
)

func ExampleEnv() {
	// In a real program the variable comes from the environment.
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Fatal(err)
	}

	port, err := getenv.Env[int]("PORT")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(port)
	// Output: 8080
}

func ExampleEnv_errors() {
	if err := os.Setenv("WORKERS", "four"); err != nil {
		log.Fatal(err)
	}

	_, err := getenv.Env[int]("WORKERS")
	fmt.Println(errors.Is(err, getenv.ErrInvalidValue), err)

	_, err = getenv.Env[int]("GETENV_EXAMPLE_NOT_SET")
	fmt.Println(errors.Is(err, getenv.ErrNotSet), err)

	// Output:
	// true environment variable "WORKERS": invalid value: strconv.ParseInt: parsing "four": invalid syntax
	// true environment variable "GETENV_EXAMPLE_NOT_SET": not set
}

func ExampleEnv_slice() {
	if err := os.Setenv("HOSTS", "a.example.com,b.example.com"); err != nil {
		log.Fatal(err)
	}

	hosts, err := getenv.Env[[]string]("HOSTS")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hosts)
	// Output: [a.example.com b.example.com]
}

func ExampleEnv_url() {
	if err := os.Setenv("DATABASE_URL", "postgres://user:secret@db:5432/app"); err != nil {
		log.Fatal(err)
	}

	dsn, err := getenv.Env[url.URL]("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dsn.Host, dsn.Redacted())
	// Output: db:5432 postgres://user:xxxxx@db:5432/app
}

func ExampleEnvOrDefault() {
	// "30" has no unit, so it is not a valid time.Duration: EnvOrDefault falls back
	// to the default, the same as for a variable that is not set. Use Env to catch it.
	if err := os.Setenv("TIMEOUT", "30"); err != nil {
		log.Fatal(err)
	}

	timeout := getenv.EnvOrDefault("TIMEOUT", 5*time.Second)
	retries := getenv.EnvOrDefault("GETENV_EXAMPLE_RETRIES", 3)

	fmt.Println(timeout, retries)
	// Output: 5s 3
}

func ExampleWithSeparator() {
	if err := os.Setenv("PORTS", "80|443"); err != nil {
		log.Fatal(err)
	}

	ports, err := getenv.Env[[]int]("PORTS", getenv.WithSeparator("|"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ports)
	// Output: [80 443]
}

func ExampleWithTimeLayout() {
	if err := os.Setenv("RELEASE_DATE", "2022-01-20"); err != nil {
		log.Fatal(err)
	}

	date, err := getenv.Env[time.Time]("RELEASE_DATE", getenv.WithTimeLayout(time.DateOnly))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	// Output: 2022-01-20 00:00:00 +0000 UTC
}
