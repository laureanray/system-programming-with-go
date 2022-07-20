package impl

import (
	"fmt"
	"io"
	"os"
)

func Load(file string) {
	if file == "" {
		fmt.Println("Please specify a file.")
		return
	}

	f, err := os.Open(file)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer f.Close() // ensure close to avoid leaks

	var b = make([]byte, 512)

	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Print(string(b[:n])) // only print whats been read
		}
	}

	if err != nil && err != io.EOF {
		fmt.Println("\n\nError: ", err)
	}
}
