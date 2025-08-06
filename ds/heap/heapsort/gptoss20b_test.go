package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// GOWORK=off GOFLAGS="-count=1" go test -run TestHeapSort
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
