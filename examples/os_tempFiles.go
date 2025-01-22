package main

import (
	"fmt"
	"os"
)

func main() {
	createTempFile()
}

func createTempFile() {
	// START OMIT
	tempFile, _ := os.CreateTemp("/tmp", "someStuff_*.txt") // HL
	fmt.Println("Created temp file:", tempFile.Name())
	// END OMIT
}
