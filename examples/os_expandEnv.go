package main

import (
	"fmt"
	"os"
)

// START OMIT

func main() {
	os.Setenv("FOO", "bar")
	s := os.ExpandEnv("Hello ${FOO}, how are you?") // HL
	fmt.Println(s)
}

// END OMIT
