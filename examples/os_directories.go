package main

import (
	"os"
)

func main() {
	// START OMIT
	_ = os.Mkdir("example_dir", 0755) // HL
	printDirectory()
	_ = os.Remove("example_dir") // HL
	// END OMIT
}
