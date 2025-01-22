package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	s, _ := os.Hostname()
	fmt.Println(s)
	// END OMIT
}
