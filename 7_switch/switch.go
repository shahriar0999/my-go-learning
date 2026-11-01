package main

import "fmt"

func main() {
	role := "admin"

	switch role {
	case "admin":
		fmt.Println("Welcome Admin")

	case "Editor":
		fmt.Println("Welcome Editor")

	case "Viewer":
		fmt.Println("Welcome Viewer")

	default:
		fmt.Println("Unknown role")
	}
}
