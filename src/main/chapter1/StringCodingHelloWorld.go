package main

import (
	"fmt"
)

// type literal declaration
type Something string

const (
	MY_STRING Something = "Something"
)

func main() {
	fmt.Println()
	fmt.Println(MY_STRING)
	fmt.Println("MY_STRING ("+MY_STRING+") == 'Something'?", (MY_STRING == "Something"))
}
