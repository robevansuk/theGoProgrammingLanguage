package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * This application replicates the echo command in the Unix terminal
 * by returning whatever args are passed to it
 * Usage:
 * go run ArgsFromCLI_v3.go 123 456
 * Outputs:
 * 123 456
 * [123 456]
 */
// Implementation of Unix Echo command
func main() {
	// Demo the Join function from the strings package
	fmt.Println(strings.Join(os.Args[1:], " "))
	// If we don't care abt formatting
	fmt.Println(os.Args[1:], " ")
}
