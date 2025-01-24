package main

import (
	"fmt"
	"os/user"
)

func main() {
	// START OMIT
	ericUser, _ := user.Lookup("eric.sims")
	fmt.Printf("Group: %+v \n", ericUser)

	currentUser, _ := user.Current()
	fmt.Printf("Group: %+v \n", currentUser)
	// END OMIT
}
