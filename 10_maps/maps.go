package main

import "fmt"

func main() {
	// initialize a map
	var person map[string]string = map[string]string{
		"name":    "Rahim",
		"role":    "MLops Engineer",
		"country": "Bangladesh",
	}

	// print the map
	println(person["name"])

	// read values of a map using a loop
	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// new value assignment to a map
	person["name"] = "Karim"
	person["role"] = "Data Scientist"
	person["country"] = "India"

	fmt.Println("===================================")
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	// delete a key-value pair from a map
	delete(person, "country")

	fmt.Println("===================================")
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	//
}
