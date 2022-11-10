package main

import (
	"bufio"
	"fmt"
	"os"
)

// Should read contents of a file and report back if any duplicate lines are found
// Usage: go run Dup_v1.go /Users/${me}/SOME_FILE_ON_YOUR_MACHINE
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file) // f is passed to the scanner to be read - Scanner can read files.
			if err != nil {
				// %v - any value in natural format
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s'n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignore any potential errors from input.Err()
}
