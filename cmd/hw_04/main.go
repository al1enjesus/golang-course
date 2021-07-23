package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

func (x Circle) Area() float64 {
	return math.Pi * x.radius * x.radius
}

func (x Circle) Perimeter() float64 {
	return 2 * math.Pi * x.radius
}

func (x Rectangle) Area() float64 {
	return x.width * x.height
}

func (x Rectangle) Perimeter() float64 {
	return (x.width + x.height) * 2
}

func (x Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", x.radius)
}

func (x Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", x.height, x.width)
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
func main() {
	c := Circle{8}
	r := Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)
}