package main

import (
	"fmt"
)

// See
// https://leetcode.com/problems/container-with-most-water/

func minInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Returns the max area from the array of heights
func maxArea(heights []int) int {
	count := len(heights)
	if count == 0 {
		return 0
	}
	j := count - 1
	maxArea := 0
	for i, h := range heights {
		leftWall := h
		rightWall := heights[j]
		wallHeight := minInt(leftWall, rightWall)
		area := wallHeight * wallHeight
		if area > maxArea {
			maxArea = area
		}
		j -= i
		if j == i || j < 0 {
			break
		}
	}
	return maxArea
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(input)
	fmt.Printf("Input  %v\n", input)
	fmt.Printf("Result %v\n", result) // 49
}
