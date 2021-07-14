package main

import (
	"fmt"
	"golang-course/cmd/hw_02/fibonacci/dp"
	"golang-course/cmd/hw_02/fibonacci/recursive"
)

func main() {
	fmt.Println("Enter n, amount of Fibonacci nums. Range: 1-93")
	var n int
	fmt.Scanf("%d\n", &n)

	if n < 1 || n > 93 {
		fmt.Println("Incorrect value")
		return
	}

	fmt.Println("Choose type of solution: 1-dynamic programming, 2-recursion")
	var chose int

	_, err := fmt.Scanf("%d\n", &chose)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch chose {
	case 1:
		fmt.Println("Linear solution: ")
		dp.PrintSequence(n)
	case 2:
		fmt.Println("Recursive solution: ")
		recursive.PrintSequence(n)
	default:
		fmt.Println("Incorrect value")

	}
}
