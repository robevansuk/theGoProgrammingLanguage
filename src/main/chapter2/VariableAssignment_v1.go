package main

import "fmt"

func main() {
	i, j := 3, 5
	isValid := true
	isNotValid := false
	// tuple assignment - swaps j for i and i for j, and swaps isValid (true) for isNotValid (false) and vice versa.
	i, j, isValid, isNotValid = j, i, isNotValid, isValid

	in, err := someFuncReturningAValAndAnErr()
	// Variable declaration only needs one new variable (out) - err can be reused
	out, err := someFuncReturningAValAndAnErr()

	fmt.Printf("%s, %s, %v", in, out, err)
}

func someFuncReturningAValAndAnErr() (string, error) {
	return "", nil
}
