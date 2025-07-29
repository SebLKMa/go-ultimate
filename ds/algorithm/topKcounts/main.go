package main

import (
	"fmt"
	"slices"
)

func topKcounts(input []int, k int) (output []int) {
	if k <= 0 || len(input) == 0 {
		return
	}
	if len(input) == 1 {
		output = append(output, 1)
		return
	}

	slices.Sort(input)
	max := len(input) - 1
	count := 1
	counts := []int{}
	indices := []int{}
	index := 0
	for index = 1; index <= max; index++ {
		if input[index] != input[index-1] {
			counts = append(counts, count)
			indices = append(indices, index)
			count = 0
		}
		count++
	}
	if count != 0 {
		counts = append(counts, count)
		indices = append(indices, index)
	}

	fmt.Printf("counts  %v\n", counts)
	fmt.Printf("indices  %v\n", indices)

	return
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	topKs := topKcounts(arr, 2)
	fmt.Printf("array  %v\n", arr)
	fmt.Printf("Top Ks %v\n", topKs) // 1 2

	arr = []int{1, 2, 2, 3, 3, 3}
	topKs = topKcounts(arr, 2)
	fmt.Printf("array  %v\n", arr)
	fmt.Printf("Top Ks %v\n", topKs) // 3 2
}
