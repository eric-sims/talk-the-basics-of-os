package main

import (
	"fmt"
	"os"
)

func main() {
	createTempFile()
}

// START OMIT
func createTempFile() {
	tempFile, _ := os.CreateTemp("/tmp", "someStuff_*.txt") // HL
	fmt.Println("Created temp file:", tempFile.Name())
}

// END OMIT
