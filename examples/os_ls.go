package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			name += "/"
		}
		fmt.Println(name)
	}
	// END OMIT
}
