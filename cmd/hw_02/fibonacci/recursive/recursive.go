package recursive

import "fmt"

func fib(n int) int64{
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}


func PrintSequence(n int) {
	// Recursive implementation O(N*(e^N))

	defer fmt.Println("Program completed")
	var numbers []int64

	for i := 0; i+1 <= n; i++ {
		numbers = append(numbers, fib(i))
	}

	fmt.Println(numbers)
}
