package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// different parameter with single data type
func subtract(x, y int) int {
	return x - y
}

// function without return keyword
func greet(name string) {
	fmt.Println("Hello", name)
}

// function with multiple return values
func calculate(x, y int) (int, int) {
	return add(x, y), subtract(x, y)
}

func main() {
	greet("Crown AI")
	// save function in a variable
	cal := func(x int) string {
		if x%2 == 0 {
			return "Even"
		} else {
			return "Odd"
		}
	}
	fmt.Println(cal(10))
	sum, sub := calculate(10, 5)
	fmt.Println(sum, sub)
	result := add(3, 4)
	fmt.Println(result)

	result1 := subtract(10, 5)
	fmt.Println(result1)
}
