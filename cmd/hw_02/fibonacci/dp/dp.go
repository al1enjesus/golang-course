package dp

import "fmt"

func PrintSequence(n int) {
	// Dynamic programming implementation O(N)

	defer fmt.Println("Program completed")
	numbers := []int64{0, 1}

	if n < 3 {
		fmt.Println(numbers[:n])
	} else {
		for i := 2; i+1 <= n; i++ {
			numbers = append(numbers, numbers[i-2]+numbers[i-1])
		}

		fmt.Println(numbers)
	}
}
