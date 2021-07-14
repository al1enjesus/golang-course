package main

import (
	"fmt"
	"golang-course/cmd/hw_02/fibonacci/fibo_dynamic"
	"golang-course/cmd/hw_02/fibonacci/fibo_recursive"
	"golang-course/cmd/hw_02/fibonacci/fibo_simplest"
)

func main() {
	fmt.Println("Enter n, amount of Fibonacci nums. Range: 1-93")
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		fmt.Println(err)
		return
	}

	if n < 1 || n > 93 {
		fmt.Println("Incorrect value")
		return
	}
	fmt.Println("Choose type of solution: 1-dynamic programming, 2-recursion, 3-the simplest")
	var chose int
	_, err = fmt.Scanf("%d\n", &chose)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch chose {
	case 1:
		fmt.Println("Linear solution: ")
		fibo_dynamic.PrintSequence(n)
	case 2:
		fmt.Println("Recursive solution: ")
		fibo_recursive.PrintSequence(n)
	case 3:
		fmt.Println("The simplest solution: ")
		fibo_simplest.PrintSequence(n)
	default:
		fmt.Println("Incorrect value")
	}
}
