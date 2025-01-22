package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// START OMIT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := <-c
	fmt.Println("Got signal:", s)
	// END OMIT
}
