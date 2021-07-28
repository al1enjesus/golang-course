package main

import (
	"fmt"
	"golang-course/cmd/hw_04/shape"
	"golang-course/cmd/hw_04/shape/circle"
	"golang-course/cmd/hw_04/shape/rectangle"
)

func DescribeShape(s shape.Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	c := circle.New(8)
	r := rectangle.New(9, 3)
	DescribeShape(c)
	DescribeShape(r)
}