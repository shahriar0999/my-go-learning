package main

import "fmt"

func main() {
	const language string = "Go"
	const releaseYear int = 2009

	fmt.Println(language)
	fmt.Println(releaseYear)

	// multiple constants can be declared in a single block
	const (
		port = 8000
		host = "localhost"
	)

	fmt.Println(port, host)
}
