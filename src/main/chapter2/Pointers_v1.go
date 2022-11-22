package main

import "fmt"

func main() {
	someValue := "someString"
	somePointerToSomeValue := &someValue           // & to get pointer
	dereferencedPointer := *somePointerToSomeValue // * to dereference a pointer

	dereferenceAPointer := dereferenceAPointer(somePointerToSomeValue)
	fmt.Printf("%s", dereferencedPointer)
}

func dereferenceAPointer(somePointerToAValue *string) string {
	return *somePointerToAValue
}
