package main

import "fmt"

func rotateLeft(d int, input []int) (output []int) {
	output = input[d:]
	for _, i := range input[:d] {
		output = append(output, i)
	}
	return
}

func main() {
	before := []int{1, 2, 3, 4, 5}
	after := rotateLeft(2, before)
	fmt.Printf("Before %v\n", before)
	fmt.Printf("After  %v\n", after) // 3 4 5 1 2
}
