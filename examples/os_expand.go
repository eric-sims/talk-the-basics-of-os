package main

import (
	"fmt"
	"os"
)

// START OMIT

func getName(s string) string {
	switch s {
	case "CNAME":
		return "Charlie"
	case "RNAME":
		return "Rachel"
	default:
		return "Friend"
	}
}
func main() {
	s := os.Expand("Hello ${CNAME}, how are you?", getName) // HL
	fmt.Println(s)
}

// END OMIT
