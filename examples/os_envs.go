package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	fmt.Println(os.Getenv("RANDOM_ENV"))
	_ = os.Setenv("RANDOM_ENV", "Hello, world!") // HL
	defer os.Unsetenv("RANDOM_ENV")

	if str, ok := os.LookupEnv("RANDOM_ENV"); ok { // HL
		fmt.Println(str)
	}

	allEnvs := os.Environ() // HL
	fmt.Println(allEnvs[len(allEnvs)-1:])
	// END OMIT
}
