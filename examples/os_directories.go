package main

import (
	"os"
)

func main() {
	makeAndRemoveDirectory()
}

// START OMIT
func makeAndRemoveDirectory() {
	_ = os.Mkdir("example_dir", 0755) // HL
	printDirectory()
	_ = os.Remove("example_dir") // HL
}

// END OMIT
