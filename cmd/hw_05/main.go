package main

import (
	"fmt"
	"golang-course/cmd/hw_05/shape"
	"golang-course/cmd/hw_05/shape/circle"
	"golang-course/cmd/hw_05/shape/rectangle"
)

func DescribeShape(s shape.Shape) {
	fmt.Println(s)

	resArea, errArea := s.Area()
	if errArea != nil {
		panic(fmt.Sprintf("Error: %v", errArea.Error()))
	}
	fmt.Printf("Area: %.2f\n", resArea)

	resPerimeter, errPerimeter := s.Perimeter()
	if errPerimeter != nil {
		panic(fmt.Sprintf("Error: %v", errPerimeter.Error()))
	}
	fmt.Printf("Perimeter: %.2f\n", resPerimeter)
}

func main() {
	c := circle.New(8)
	r := rectangle.New(9, 3)
	DescribeShape(c)
	DescribeShape(r)
}