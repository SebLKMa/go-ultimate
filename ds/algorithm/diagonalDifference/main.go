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

// https://www.apprajapati.com/2016/06/diagonal-difference.html
func diagonaldifference(matrix [][]int) int {
	maxrows := len(matrix) - 1 // zero index
	var d1, d2 int
	fmt.Printf("Matrix Rows: %d\n", maxrows)
	/*
		[1 2 3]
		[4 5 6]
		[9 8 9]

		d1: 1+5+9 = 15
		d2: 3+5+9 = 17
	*/

	// From left
	j := 0 // column position
	for i := 0; i <= maxrows; i++ {
		d1 += matrix[i][j]
		j++ // next column for next row
	}

	// From right
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
	fmt.Printf("Answer: %d\n", dd)

	matrix = [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	dd = diagonaldifference(matrix)
	fmt.Printf("Answer: %d\n", dd)
}
