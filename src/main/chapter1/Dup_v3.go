package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Should read contents of a file and report back if any duplicate lines are found
// Usage: go run Dup_v3.go /Users/${me}/SOME_FILE_ON_YOUR_MACHINE
func main() {
	counts := make(map[string]int)

	for _, file := range os.Args[1:] { // slice (array list)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			// %v - any value in natural format
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s'n", n, line)
		}
	}
}

// Cannot redeclare a function with the same name as another within the same package
//func countLines(f *os.File, counts map[string]int) {
//	input := bufio.NewScanner(f)
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//	// ignore any potential errors from input.Err()
//}
