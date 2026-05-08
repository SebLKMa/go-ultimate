package main

import "fmt"

func rotateLeft(d int, input []int) (output []int) {
	output = input[d:]  // starts from index d
	moving := input[:d] // elements to move
	fmt.Println(output)
	fmt.Println(moving)
	for _, i := range moving {
		output = append(output, i)
	}
	return
}

func main() {
	before := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before move: %v\n", before)
	moveSteps := 2
	after := rotateLeft(moveSteps, before)
	fmt.Printf("After moved %d steps: %v\n", moveSteps, after) // 3 4 5 1 2
	moveSteps = 3
	after = rotateLeft(moveSteps, before)
	fmt.Printf("After moved %d steps: %v\n", moveSteps, after) // 4 5 1 2 3
}
