package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Allow to enter")
	} else {
		fmt.Println("Not allowed to enter")
	}

	marks := 80

	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks >= 80 {
		fmt.Println("Grade B")
	} else if marks >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Fail")
	}
}
