package main

import (
	"fmt"
	"slices"
	"sort"
)

func topKcounts(input []int, k int) (output []int) {
	if k <= 0 || len(input) == 0 {
		return
	}
	if len(input) == 1 {
		output = append(output, 1)
		return
	}

	type kcountStruct = struct {
		kcount int
		key    int
	}
	kcounts := []kcountStruct{}

	slices.Sort(input)
	maxLen := len(input) - 1
	count := 1
	index := 0
	for index = 1; index <= maxLen; index++ {
		if input[index] != input[index-1] {
			kcounts = append(kcounts, kcountStruct{count, input[index-1]})
			count = 0
		}
		count++
	}
	// Remaining count of the last element
	if count != 0 {
		kcounts = append(kcounts, kcountStruct{count, input[index-1]})
	}

	// Sort descending
	sort.Slice(kcounts, func(i, j int) bool {
		return kcounts[i].kcount > kcounts[j].kcount
	})
	fmt.Printf("kcounts  %v\n", kcounts)

	// Return the top Ks
	for i, kc := range kcounts {
		if i >= k {
			break
		}
		output = append(output, kc.key)
	}
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
