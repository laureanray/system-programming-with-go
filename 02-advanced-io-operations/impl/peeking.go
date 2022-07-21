package impl

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
Peeking is the ability to read content without advancing the reader cursor.
Here, under the hood, the peeked data is stored in the buffer.
Each reading operation checks whether there's data in this buffer and if there is any, that data is returned while removing it from the buffer.
This works like a queue (first in, first out).
*/
func Peek(path string) {
	if path == "" {
		fmt.Println("Please specify a path.")
		return
	}

	f, err := os.Open(path)

	if err != nil {
		fmt.Println("Error:", err)
	}

	defer f.Close()

	r := bufio.NewReader(f) // wrapping the reader with a buffered one

	var rowCount int

	for err == nil {
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Print(string(b))
			}
		}

		// each time moar is false, a line is completely read

		if err == nil {
			fmt.Println()
			rowCount++
		}
	}

	if err != nil && err != io.EOF {
		fmt.Println("\nError:", err)
		return
	}

	fmt.Println("\nRow count:", rowCount)
}
