package main

import (
	"io/fs"
	"os"
)

func main() {
	// START OMIT
	os.CreateTemp("/tmp", "os_permissions")
	os.Chmod("/tmp", fs.FileMode(0755))
	// END OMIT
}
