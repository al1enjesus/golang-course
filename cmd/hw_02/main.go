package main

import (
	"fmt"
	"golang-course/cmd/hw_02/fibonacci/dp"
)

func main() {
	fmt.Println("Enter n, amount of Fibonacci nums. Range: 1-93")
	var n int
	fmt.Scanf("%d", &n)

	if n < 1 || n > 93 {
		fmt.Println("Incorrect value")
	} else {
		dp.PrintSequence(n)
	}
}
