package heapint

// See
// https://pkg.go.dev/container/heap#example-package-IntHeap

// HeapInt is default a min-heap of ints.
type HeapInt []int

func (h HeapInt) Len() int           { return len(h) }
func (h HeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h HeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Implements the heap interface
func (h *HeapInt) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Implements the heap interface
func (h *HeapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
