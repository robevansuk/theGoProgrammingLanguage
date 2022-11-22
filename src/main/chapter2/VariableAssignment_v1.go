package main

func main() {
	i, j := 3, 5
	isValid := true
	isNotValid := false
	// tuple assignment - swaps j for i and i for j, and swaps isValid (true) for isNotValid (false) and vice versa.
	i, j, isValid, isNotValid = j, i, isNotValid, isValid
}
