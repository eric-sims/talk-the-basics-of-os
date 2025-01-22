package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	os.Setenv("FOO", "bar")
	s := os.ExpandEnv("Hello ${FOO}, how are you?") // HL
	fmt.Println(s)
	// END OMIT
}
