package main

// Vibe-coded from ollama run gpt-oss:20b

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
		sink(a, 0, end)             // restore heap property
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

func main() {
	// TODO
}
