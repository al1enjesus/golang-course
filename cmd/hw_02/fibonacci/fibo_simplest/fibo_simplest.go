package fibo_simplest

import "fmt"

func PrintSequence(n int) {
	// Simplest solution implementation O(N)

	defer fmt.Println("Program completed")

	if n == 1 {
		fmt.Println("0")
		return
	}
	if n == 2 {
		fmt.Println("0 1")
		return
	}

	a, b := 0, 1
	fmt.Printf("%d %d ", a, b)
	for i := 2; i < n; i++ {
		temp := b
		b += a
		a = temp
		fmt.Printf("%d ", b)
	}

	fmt.Println()
}