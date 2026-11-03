package main

import "fmt"

func main() {
	type Employee struct {
		Name    string
		Age     int
		Address string
	}

	emp1 := Employee{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	emp2 := Employee{
		Name:    "Jane Smith",
		Age:     25,
		Address: "456 Elm St",
	}

	fmt.Println(emp1.Name)
	fmt.Println(emp2.Name)
}
