package main

import "fmt"

func main() {
	var language string = "Go"
	fmt.Println(language)

	var releaseYear int = 2009
	fmt.Println(releaseYear)

	var currentVersion float32 = 1.20
	fmt.Println(currentVersion)

	var isOpenSource bool = true
	fmt.Println(isOpenSource)

	// Auto type inference
	var framework = "Gin"
	fmt.Println(framework)

	var isFast = true
	fmt.Println(isFast)

	// Short declaration operator
	year := 2024
	fmt.Println(year)
}
