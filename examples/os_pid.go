package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	fmt.Printf("PID: %d\n", os.Getpid())    // Current process ID
	fmt.Printf("PPID: %d\n", os.Getppid())  // Parent process ID
	fmt.Printf("Command: %s\n", os.Args[0]) // Command name
	// END OMIT
}
