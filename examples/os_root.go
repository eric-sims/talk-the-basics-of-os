package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// START OMIT
	root, err := os.OpenRoot("someDir")
	if err != nil {
		panic(err)
	}
	defer root.Close()

	_, err = root.Open("../examples/hello.go") // HL
	if err != nil {
		fmt.Println("expected error:", err)
	}

	file, err := root.Open("someOtherDir/../someOtherDir/hello.txt")
	if err != nil {
		panic(err)
	}
	contents, err := io.ReadAll(file)
	fmt.Println(string(contents))
	// END OMIT
}
