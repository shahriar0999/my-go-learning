package main

import "fmt"

func main() {
	age := 30

	fmt.Println(age)
	fmt.Println(&age)

	// Dereference pointer
	agePointer := &age
	fmt.Println(*agePointer)

	// Change value using pointer
	*agePointer = 31
	fmt.Println(age)
}
