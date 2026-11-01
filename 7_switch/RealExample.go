package main

import "fmt"

func main() {

	orderStatus := "shipped"

	switch orderStatus {
	case "pending":
		fmt.Println("Your order is pending. Please wait for confirmation.")
	case "shipped":
		fmt.Println("Your order has been shipped. It will arrive soon.")
	case "delivered":
		fmt.Println("Your order has been delivered. Enjoy your purchase!")
	default:
		fmt.Println("Unknown order cancellation.")
	}
}
