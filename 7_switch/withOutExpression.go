package main

import "fmt"

func main() {
	age := 25

	switch {
	case age < 18:
		fmt.Println(("Teenager"))
	case age >= 18 && age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior Citizen")
	}

}
