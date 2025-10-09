package main

import "fmt"

func main() {
	var age int        // Declaring an integer variable
	var name string    // Declaring an string variable
	var isStudent bool // Declaring a boolean variable

	fmt.Println("Name:", name)            // Default value is ""
	fmt.Println("Age:", age)              // Default value is 0
	fmt.Println("Is Student:", isStudent) // Default value is false

	// values assignment after declaration
	age = 25
	name = "Shahriar"
	isStudent = false

	fmt.Println("After assignment:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)

	// variable in block
	var (
		city    string
		country string
		zipcode int
	)

	fmt.Println("City:", city) // Default value is ""
	fmt.Println("Country:", country)
	fmt.Println("Zipcode:", zipcode) // Default value is 0
}
