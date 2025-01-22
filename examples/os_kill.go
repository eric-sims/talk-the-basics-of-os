package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// START OMIT
	if len(os.Args) < 2 { // HL
		fmt.Println("Usage: kill <PID>")
		os.Exit(1)
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid PID: %v\n", err)
		os.Exit(1)
	}

	process, err := os.FindProcess(pid) // HL
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding process: %v\n", err)
		os.Exit(1)
	}

	if err = process.Kill(); err != nil { // HL
		fmt.Fprintf(os.Stderr, "Error killing process: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Killed process with PID %d\n", pid)
	// END OMIT
}
