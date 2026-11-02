package main

import "fmt"

func main() {
	var name [5]string = [5]string{"Rahim", "Karim", "Jabbar", "Jalal", "Jamal"}

	profession := [5]string{
		"Engineer",
		"Doctor",
		"Teacher",
		"Driver",
		"Farmer",
	}
	// fmt.Println(name)
	// fmt.Println(profession)

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
	fmt.Println("===================================")
	for i := 0; i < len(profession); i++ {
		fmt.Println(profession[i])
	}

	// change the value of an array
	name[0] = "Rahman"
	profession[0] = "MLops Engineer"
	fmt.Println("===================================")

	fmt.Println(name[0], "-----", profession[0])

	// initialize an array without specifying the value of the array
	var age [5]int

	age[0] = 25
	age[1] = 29
	age[2] = 30
	age[3] = 35
	age[4] = 40

	fmt.Println("===================================")
	for i := 0; i < len(age); i++ {
		fmt.Println(name[i], "-----", profession[i], "-----", age[i])
	}
}
