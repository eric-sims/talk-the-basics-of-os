package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	// START OMIT
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe() // HL
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "this is being written to stdin")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// END OMIT
}
