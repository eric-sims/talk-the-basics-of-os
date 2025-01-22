package main

import (
	"fmt"
	"os"
)

// START OMIT
func main() {
	file, _ := os.Create("example.txt") // HL
	defer file.Close()

	bytesWritten, _ := file.WriteString("Hello, Go!")

	fmt.Println("bytes written:", bytesWritten)
}

// END OMIT
