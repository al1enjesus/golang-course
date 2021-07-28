package rectangle

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	width float64
	height float64
}

func (x Rectangle) Area() (float64, error) {
	if x.width < 0 {
		return 0, errors.New("Negative width\n")
	}
	if x.height < 0 {
		return 0, errors.New("Negative height\n")
	}
	return x.width * x.height, nil
}

func (x Rectangle) Perimeter() (float64, error) {
	if x.width < 0 {
		return 0, errors.New("Negative width\n")
	}
	if x.height < 0 {
		return 0, errors.New("Negative height\n")
	}
	return (x.width + x.height) * 2, nil
}

func (x Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", x.height, x.width)
}

func New(width float64, height float64) Rectangle {
	return Rectangle{width: width, height: height}
}