package main

import "fmt"

func main() {
	// message := "Hello World"

	// // closure function
	// greet := func() {
	// 	fmt.Println(message)
	// }

	// greet()

	// closure function with change variable value
	count := 0

	increment := func() int {
		count++
		return count
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
