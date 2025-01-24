package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	err := os.Mkdir("mkdir_ex", 0755) // HL
	if err != nil {
		fmt.Println(err)
	}
	//err := os.Remove("mkdir_ex")   // HL
	//if err != nil {
	//	fmt.Println(err)
	//}
	// END OMIT
}

//os.MkdirAll("mkdir_ex/subdir/parent/child", 0755)
