package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	dir := "."
	if len(os.Args) > 1 { // HL
		dir = os.Args[1]
	}

	dirEntries, err := os.ReadDir(dir) // HL
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		os.Exit(1) // HL
	}

	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		if dirEntry.IsDir() {
			name += "/"
		}
		fmt.Println(name)
	}
	// END OMIT
}
