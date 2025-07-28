package main

import (
	"fmt"
	"slices"
)

func printMinMaxSum(arr []int, lenSum int) {
	sorted := arr
	slices.Sort(sorted)
	var minSum, maxSum int
	for i := 0; i < lenSum; i++ {
		minSum += sorted[i]
	}
	max := len(sorted) - 1 // zero index
	for i := max; i > (max - lenSum); i-- {
		maxSum += sorted[i]
	}

	fmt.Printf("Min Sum: %d\n", minSum)
	fmt.Printf("Max Sum: %d\n", maxSum)
}

// Given five positive integers, find the minimum and maximum values
// that can be calculated by summing exactly four of the five integers.
// Then print the respective minimum and maximum values
func main() {
	arr := []int{1, 3, 5, 7, 9}
	printMinMaxSum(arr, len(arr)-1)
}
