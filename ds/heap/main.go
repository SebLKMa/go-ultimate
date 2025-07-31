package main

import (
	"container/heap"
	"fmt"

	"github.com/seblkma/go-ultimate/ds/heap/heapint"
	"github.com/seblkma/go-ultimate/ds/heap/priorityqueueitem"
)

// GOWORK=off go run main.go

// This example inserts several ints into an HeapInt, checks the minimum,
// and removes them in order of priority.
func main() {
	fmt.Println("Heap")
	h := &heapint.HeapInt{2, 1, 5}
	heap.Init(h)
	// the heap will is sorted by the heap
	heap.Push(h, 3)
	heap.Push(h, 7)
	heap.Push(h, 4)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("PriorityQueue")

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	// Some items and their priorities.
	items := map[string]int{
		"C++": 3, "Python": 1, "Rust": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(priorityqueueitem.PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &priorityqueueitem.Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &priorityqueueitem.Item{
		Value:    "Go",
		Priority: 99,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 42)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*priorityqueueitem.Item)
		fmt.Printf("%s:%.2d ", item.Value, item.Priority)
	}
	fmt.Println()
}
