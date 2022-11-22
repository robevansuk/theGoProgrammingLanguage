package main

import "fmt"

// package level declaration
const boilingF = 212.0

// package level declaration - main()
func main() {
	// local declarations f and c
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g °C\n", f, c)
}
