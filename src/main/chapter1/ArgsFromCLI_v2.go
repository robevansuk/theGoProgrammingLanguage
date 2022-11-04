package main

import (
	"fmt"
	"os"
)

/**
 * This application replicates the echo command in the Unix terminal
 * by returning whatever args are passed to it
 * Usage:
 * go run ArgsFromCLI_v2.go 123 456
 * Outputs:
 * 123 456
 */
// Implementation of Unix Echo command
func main() {
	// modify the assignment
	s, sep := "", ""
	// modified the loop to use a range rather than a counter
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
