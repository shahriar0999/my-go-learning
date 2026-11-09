package main

import "fmt"

type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
	Canceled
)

func getOrderStatus(status OrderStatus) string {

	switch status {

	case Pending:
		return "Your Order is Pending"

	case Processing:
		return "Your Order is Processing"

	case Shipped:
		return "Your Order is Shipped"

	case Delivered:
		return "Your Order is Delivered"

	case Canceled:
		return "Your Order is Canceled"

	default:
		return "Invalid Order Status"
	}

}

func main() {
	var myOrder OrderStatus = Shipped

	message := getOrderStatus(myOrder)

	fmt.Println(message)
}
