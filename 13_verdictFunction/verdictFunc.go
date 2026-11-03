package main

import "fmt"

func totalbill(bill ...int) int {
	total := 0

	for _, bl := range bill {
		total += bl
	}

	return total
}

func main() {
	fmt.Println("Friday Bill ", totalbill(22, 33, 31, 44, 11))
	fmt.Println("Saturday Bill ", totalbill(11, 12, 13, 14, 44, 22))
}
