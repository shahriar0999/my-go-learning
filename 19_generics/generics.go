package main

import "fmt"

func Add[T int | float64](a T, b T) T {
	return a + b
}

func main() {
	var intResult int = Add(4, 9)
	var floatResult float64 = Add(4.5, 10.9)

	fmt.Println("Integer Addition ", intResult)
	fmt.Println("Float Addition ", floatResult)
}
