package main

import "fmt"

// go have only for loop, no while or do while loop

func main() {
	// while loop implementation
	i := 0
	for i < 11 {
		fmt.Println(i)

		i = i + 1
	}

	// for loop implementation
	for j := 0; j < 11; j++ {

		if j%2 == 0 {
			fmt.Println(j)
		}
	}

	// using range in for loop
	for k := range 11 {
		fmt.Println(k)

	}
}
