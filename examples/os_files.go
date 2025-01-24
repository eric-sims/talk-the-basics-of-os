package main

import (
	"os"
)

func main() {
	// START OMIT
	file, _ := os.Create("example.txt") // HL
	file.WriteString("Hello, there!")
	file.Close()

	f, _ := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // HL
	f.WriteString("some more text!")
	f.Close()
	// END OMIT
}
