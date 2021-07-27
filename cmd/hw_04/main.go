package main

import (
	"fmt"
	"golang-course/cmd/hw_04/shape"
)

func DescribeShape(s shape.Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	c := shape.NewCircle(8)
	r := shape.NewRectangle(9, 3)
	DescribeShape(c)
	DescribeShape(r)
}