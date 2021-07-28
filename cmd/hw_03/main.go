package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

func avgValue(arr [6]float64) float64 {
	// Returns average value
	var sum float64
	for _, element := range arr {
		sum += element
	}
	return sum / float64(len(arr))
}

func max(arr []string) string {
	// Returns the longest word
	var maxLength int
	var answer string
	for _, element := range arr {
		// utf8.RuneCountInString finding real length, when len() finding amount of bytes
		length := utf8.RuneCountInString(element)
		if length > maxLength {
			maxLength = length
			answer = element
		}
	}
	return answer

}

func reverse(sl []float64) []float64 {
	// Returns reversed copy of slice
	tempSlice := make([]float64, len(sl))
	copy(tempSlice, sl)

	for i := len(tempSlice)/2 - 1; i >= 0; i-- {
		opp := len(tempSlice) - 1 - i
		tempSlice[i], tempSlice[opp] = tempSlice[opp], tempSlice[i]
	}

	return tempSlice
}

func printSorted(m map[int]string) {
	// Returns values in sorted keys order
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	values := make([]string, 0, len(m))
	for _, key := range keys {
		values = append(values, m[key])
	}

	fmt.Println(values)
}

func main() {
	a := [6]float64{1, 2, 3, 4, 5, 6}
	fmt.Println(avgValue(a))

	b := []string{"one", "two", "three"}
	fmt.Println(max(b))

	c := []float64{1, 2, 5, 15}
	fmt.Println(reverse(c))

	d := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(d)
}
