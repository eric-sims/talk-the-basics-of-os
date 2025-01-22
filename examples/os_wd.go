package main

import (
	"fmt"
	"os"
)

func main() {
	printWorkingDir()
}

// START OMIT
func printWorkingDir() {
	wd, err := os.Getwd() // HL
	if err != nil {
		panic(err)
	}
	fmt.Println("working dir is", wd)
}

// END OMIT
