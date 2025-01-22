package main

import (
	"log"
	"os/exec"
)

func main() {
	// START OMIT
	cmd := exec.Command("sleep", "3")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run() // HL
	log.Printf("Command finished with error: %v", err)
	// END OMIT
}
