package main

import (
	"fmt"
	"os"
)

func main() {
	uid := os.Getuid()
	guid := os.Getgid()
	fmt.Println("User id:", uid)
	fmt.Println("Group id:", guid)

	groups, err := os.Getgroups()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Group IDs: ", groups)
}
