package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// TODO: Maybe fun to implement this with goroutines?
// TODO: Or implement cool algo to improve performance in really large directories
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
		return
	}

	root, err := filepath.Abs(os.Args[1]) // Get absolute path

	if err != nil {
		fmt.Println("Cannot get absolute path: ", err)
	}

	fmt.Println("Listing files in", root)

	var c struct {
		files int
		dirs  int
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Walk the tree to count file and folders
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}

		fmt.Println("-", path)
		return nil
	})

	fmt.Printf("Total: %d files in %d directories", c.files, c.dirs)
}
