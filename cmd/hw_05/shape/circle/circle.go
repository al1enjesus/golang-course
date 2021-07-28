package circle

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (x Circle) Area() (float64, error) {
	if x.radius < 0 {
		return 0, errors.New("Incorrect radius")
	}
	return math.Pi * x.radius * x.radius, nil
}

func (x Circle) Perimeter() (float64, error) {
	if x.radius < 0 {
		return 0, errors.New("Incorrect radius")
	}
	return 2 * math.Pi * x.radius, nil
}

func (x Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", x.radius)
}

func New(radius float64) Circle {
	return Circle{radius: radius}
}