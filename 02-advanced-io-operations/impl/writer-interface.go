package impl

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReverseWrite(fileSrc string, fileDst string) {
	if fileSrc == "" || fileDst == "" {
		fmt.Println("Please provide a fileSrc and fileDst")
		return
	}

	src, err := os.Open(fileSrc)
	dst, err := os.OpenFile(fileDst, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cur, err := src.Seek(0, os.SEEK_END)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	b := make([]byte, 16)

	for step, r, w := int64(16), 0, 0; cur != 0; {
		if cur < step { // ensure cursor is 0 at max
			b, step = b[:cur], cur
		}

		fmt.Printf("[%s]\n, step", string(b))
		cur = cur - step
		_, err = src.Seek(cur, os.SEEK_SET) // go backwards

		if err != nil {
			break
		}

		if r, err = src.Read(b); err != nil || r != len(b) {
			if err == nil { // all buffer should be read
				err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
			}
			break
		}

		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			switch { // swap (\r\n) so they get back in place
			case b[i] == '\r' && b[i+1] == '\n':
				b[i], b[i+1] = b[i+1], b[i]
			case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
				b[j], b[j-1] = b[j-1], b[j]
			}
			b[i], b[j] = b[j], b[i] // swap bytes
		}

		if w, err = dst.Write(b); err != nil || w != len(b) {
			if err != nil {
				log.Println(err)
				err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
			}
		}
	}

	if err != nil && err != io.EOF {
		fmt.Println("\n\nError:", err)
	}
}
