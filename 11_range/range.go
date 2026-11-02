package main

import "fmt"

func main() {
	// initialize a map
	var person map[string]string = map[string]string{
		"name":    "Rahim",
		"role":    "MLops Engineer",
		"country": "Bangladesh",
	}

	// read values of a map using range in a loop
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	// text interation using range in a loop
	text := "Hello, World!!"
	for index, char := range text {
		// fmt.Printf("Index: %d, Character: %c\n", index, char)
		fmt.Println("Index:", index, "Character:", string(char))
	}
}
