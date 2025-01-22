package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// START OMIT
	os.Chdir("someDir")

	file, _ := os.Open("someOtherDir/hello.txt") // HL
	// file, _ := os.Open("../examples/hello.go")  // HL
	defer file.Close()

	contents, _ := io.ReadAll(file)
	fmt.Println(string(contents))
	// END OMIT
}
