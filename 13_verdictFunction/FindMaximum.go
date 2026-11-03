package main

import "fmt"

func FindMax(number ...int) int {
	max := 0

	for _, num := range number {
		if num > max {
			max = num
		}
	}

	return max

}

func main() {
	fmt.Println("Maximum Number ", FindMax(22, 3, 3, 1, 44, 2, 3, 555, 44, 99, 333))
}
