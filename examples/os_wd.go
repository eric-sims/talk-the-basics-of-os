package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	wd, err := os.Getwd() // HL
	if err != nil {
		panic(err)
	}
	fmt.Println("working dir is", wd)
	// END OMIT
}
