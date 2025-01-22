package main

import (
	"fmt"
	"os/user"
)

func main() {
	// START OMIT
	ericUser, _ := user.Lookup("eric.sims")
	printUserGroup(ericUser)

	currentUser, _ := user.Current()
	printUserGroup(currentUser)
	// END OMIT
}

func printUserGroup(group *user.User) {
	fmt.Printf("Group: %+v \n", group)
}
