package main

import (
	"advanced-io-operations/impl"
	"fmt"
)

func main() {
	fmt.Println("I/O Operations")

	impl.ReverseWrite("data/test.txt", "data/dest.txt")
}
