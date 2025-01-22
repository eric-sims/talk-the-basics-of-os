package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	file, _ := os.Create("example.txt") // HL
	defer file.Close()

	bytesWritten, _ := file.WriteString("Hello, Go!")

	fmt.Println("bytes written:", bytesWritten)
	// END OMIT
}
