package main

import (
	"fmt"
	"os"
)

func main() {
	changeDirectory()
}

// START OMIT
func changeDirectory() {
	printDirectory()
	err := os.Chdir("examples") // HL
	if err != nil {
		fmt.Println(err)
	}
	printDirectory()
}

func printDirectory() {
	d, _ := os.Getwd() // HL
	fmt.Println(d)
}

// END OMIT
