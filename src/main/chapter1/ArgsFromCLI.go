package main

import (
	"fmt"
	"os"
)

/**
 * This application replicates the echo command in the Unix terminal
 * by returning whatever args are passed to it
 * Usage:
 * go run ArgsFromCLI.go 123 456
 * Outputs:
 * 123 456
 */
// Implementation of Unix Echo command
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
