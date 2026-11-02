package main

import "fmt"

func main() {
	role := []string{"Engineer", "Doctor", "Teacher", "Driver", "Farmer"}

	fmt.Println(role)

	var name []string

	name = append(name, "Rahim")
	name = append(name, "Karim")
	name = append(name, "Jabbar")
	name = append(name, "Jalal")
	name = append(name, "Jamal")

	fmt.Println(name)

	// change the value of a slice
	name[0] = "Rahman"
	role[0] = "MLops Engineer"

	fmt.Println(name)
	fmt.Println(role)

	// part of a slice
	fmt.Println(name[2:4])

	// make with capacity
	age := make([]int, 5, 10)

	age[0] = 25
	age[1] = 29
	age[2] = 30
	age[3] = 35
	age[4] = 40

	fmt.Println(age)
	fmt.Println("Length of age slice:", len(age))
	fmt.Println("Capacity of age slice:", cap(age))
}
