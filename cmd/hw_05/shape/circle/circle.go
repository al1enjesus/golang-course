package circle

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (x Circle) Area() float64 {
	return math.Pi * x.radius * x.radius
}

func (x Circle) Perimeter() float64 {
	return 2 * math.Pi * x.radius
}

func (x Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", x.radius)
}

func New(radius float64) Circle {
	return Circle{radius: radius}
}