package main

import (
	"container/heap"
	"fmt"

	pqi "github.com/seblkma/go-ultimate/ds/heap/priorityqueueint"
)

// GOWORK=off go run main.go

// See
// https://leetcode.com/problems/top-k-frequent-elements/description/

func topKcounts(input []int, k int) (output []int) {
	if k <= 0 || len(input) == 0 {
		return
	}
	max := len(input)
	if max == 1 {
		output = append(output, 1)
		return
	}

	// A map containing inputs and their counts
	inputCounts := make(map[int]int, max)
	for _, v := range input {
		inputCounts[v] += 1
	}
	fmt.Printf("input counts %v\n", inputCounts)

	// PriorityQueue to store in highest counts order
	pq := make(pqi.PriorityQueue, len(inputCounts))
	i := 0
	for n, count := range inputCounts {
		pq[i] = &pqi.Item{
			Value:    n,
			Priority: count,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Take the items out; they arrive in decreasing priority order.
	i = 1
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqi.Item)
		fmt.Printf("%d:%.2d ", item.Value, item.Priority)
		output = append(output, item.Value)
		if i >= k {
			break
		}
		i++
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
