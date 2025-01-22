package main

import (
	"log"
	"os/exec"
)

func main() {
	// START OMIT
	cmd := exec.Command("sleep", "3")
	err := cmd.Start() // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait() //TODO: comment out // HL
	log.Printf("Command finished with error: %v", err)
	// END OMIT
}
