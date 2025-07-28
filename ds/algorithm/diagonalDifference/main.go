package main

import (
	"fmt"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func diagonaldifference(matrix [][]int) int {
	maxrows := len(matrix) - 1 // zero index
	var d1, d2 int

	/*
		[1 2 3]
		[4 5 6]
		[9 8 9]
	*/
	j := 0
	for i := 0; i <= maxrows; i++ {
		d1 += matrix[i][j]
		j++
	}

	j = len(matrix[0]) - 1
	for i := 0; i <= maxrows; i++ {
		d2 += matrix[i][j]
		j--
	}

	fmt.Printf("d1: %d\n", d1)
	fmt.Printf("d2: %d\n", d2)
	return absInt(d1 - d2)
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	dd := diagonaldifference(matrix)
	fmt.Println(dd)
}
