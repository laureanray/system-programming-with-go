package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()

	fmt.Println(wd, err)
}
