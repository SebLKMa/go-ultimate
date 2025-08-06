package main

// HeapSort sorts the slice in ascending order using a binary max\u2011heap.
func HeapSort(a []int) {
	n := len(a)

	// Build max\u2011heap
	for i := n/2 - 1; i >= 0; i-- {
		sink(a, i, n)
	}

	// Extract elements one by one
	for end := n - 1; end > 0; end-- {
		a[0], a[end] = a[end], a[0] // move max to the end
		sink(a, 0, end)              // restore heap property
	}
}

// sink moves the element at idx down the heap until the heap property holds.
// heapSize is the number of valid elements in the heap.
func sink(a []int, idx, heapSize int) {
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		largest := idx

		if left < heapSize && a[left] > a[largest] {
			largest = left
		}
		if right < heapSize && a[right] > a[largest] {
			largest = right
		}
		if largest == idx {
			break
		}
		a[idx], a[largest] = a[largest], a[idx]
		idx = largest
	}
}


package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// TestHeapSort verifies that HeapSort produces a correctly sorted slice
// for a variety of inputs.
func TestHeapSort(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
	}{
		{"empty", []int{}},
		{"single", []int{42}},
		{"sorted", []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}},
		{"random", func() []int {
			rand.Seed(time.Now().UnixNano())
			size := rand.Intn(100) + 1
			s := make([]int, size)
			for i := range s {
				s[i] = rand.Intn(1000)
			}
			return s
		}()},
	}

	for _, tc := range testCases {
		// Make a copy to avoid modifying the original test data
		got := append([]int(nil), tc.in...)
		HeapSort(got)

		want := append([]int(nil), tc.in...)
		sort.Ints(want)

		if !equalSlices(got, want) {
			t.Errorf("%s: expected %v, got %v", tc.name, want, got)
		}
	}
}

// equalSlices reports whether a and b contain the same elements in the same order.
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
