package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

var person struct {
	Name string
	Age  int
}

func main() {
	// START OMIT
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe() // HL
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(stdout).Decode(&person); err != nil { // person defined before
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	// END OMIT
}
