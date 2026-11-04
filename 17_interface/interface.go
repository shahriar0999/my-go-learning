package main

import "fmt"

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// challenge : the area of rectangle and circle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	d := Dog{}
	c := Cat{}

	fmt.Println(d.Speak())
	fmt.Println(c.Speak())

	fmt.Println("--------------------")

	r := Rectangle{Height: 5.0, Width: 10.0}
	ct := Circle{Radius: 7.0}

	fmt.Println("Area of Rectangle:", r.Area())
	fmt.Println("Area of Circle:", ct.Area())

}
