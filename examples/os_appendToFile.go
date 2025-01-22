package main

import (
	"os"
)

func main() {
	// START OMIT
	os.CreateTemp("/tmp", "tmp.txt")
	// If the file doesn't exist, create it, or append to the file
	f, _ := os.OpenFile("tmp/tmp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	_, _ = f.Write([]byte("appended some data\n"))

	// END OMIT
}
