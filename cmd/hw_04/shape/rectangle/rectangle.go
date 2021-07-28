package rectangle

import (
	"fmt"
)

type Rectangle struct {
	width float64
	height float64
}

func (x Rectangle) Area() float64 {
	return x.width * x.height
}

func (x Rectangle) Perimeter() float64 {
	return (x.width + x.height) * 2
}

func (x Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", x.height, x.width)
}

func New(width float64, height float64) Rectangle {
	return Rectangle{width: width, height: height}
}