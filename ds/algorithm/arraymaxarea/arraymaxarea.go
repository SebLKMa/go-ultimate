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
func maxAreaWrong(heights []int) int {
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
		if j <= 0 {
			break
		}
	}
	return maxArea
}

// https://leetcode.com/problems/container-with-most-water/solutions/7031738/best-explanation-visualization-simple-fast-clear/
func maxArea(heights []int) int {
	left := 0                 // start of the array on X-axis
	right := len(heights) - 1 // end of the array on X-axis
	maxArea := 0
	for left < right {
		// 1. Current_area = min(height[left], height[right]) × (right - left).
		height := minInt(heights[left], heights[right]) // heights on Y-axis
		width := right - left
		area := height * width
		fmt.Println(height, width, area)
		// 2. Update {max_area}.
		if area > maxArea {
			maxArea = area
		}
		// 3. Move {left} pointer inwards if height[left] < height[right]; else move {right} pointer inwards.
		if heights[left] < heights[right] {
			left++
		} else {
			right--
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
