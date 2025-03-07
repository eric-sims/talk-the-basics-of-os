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

// END OMIT

func printDirectory() {
	d, _ := os.Getwd()
	fmt.Println(d)
}
