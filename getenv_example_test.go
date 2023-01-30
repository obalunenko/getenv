package getenv

import (
	"fmt"
	"github.com/obalunenko/getenv/option"
	"os"
	"time"
)

func ExampleEnvOrDefault() {
	key := "GH_GETENV_TEST"
	defer func() {
		if err := os.Unsetenv("GH_GETENV_TEST"); err != nil {
			panic(err)
		}
	}()

	// string
	if err := os.Setenv(key, "golly"); err != nil {
		panic(err)
	}

	var valS string
	valS = EnvOrDefault(key, string("golly"))
	fmt.Println(valS)

	// int
	if err := os.Setenv(key, "123"); err != nil {
		panic(err)
	}

	var valI int
	valI = EnvOrDefault(key, int(123))
	fmt.Println(valI)

	// time.Time
	if err := os.Setenv(key, "2022-01-20"); err != nil {
		panic(err)
	}

	var valT time.Time
	valT = EnvOrDefault(key,
		time.Date(1992, 12, 1, 0, 0, 0, 0, time.UTC),
		option.WithTimeLayout("2006-01-02"),
	)
	fmt.Println(valT.UTC())

	// []float64
	if err := os.Setenv(key, "26.89,0.67"); err != nil {
		panic(err)
	}

	var valSlF64 []float64
	valSlF64 = EnvOrDefault(key, []float64{},
		option.WithSeparator(","),
	)
	fmt.Println(valSlF64)

	// Output:
	// golly
	// 123
	// 2022-01-20 00:00:00 +0000 UTC
	// [26.89 0.67]
}
