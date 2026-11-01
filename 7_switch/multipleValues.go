package main

import "fmt"

func main() {
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	default:
		fmt.Println("Unknown day")
	}
}
