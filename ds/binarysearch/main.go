package main

import (
	"fmt"
	"slices"
)

func bsearch(arr []int, toSearch int) bool {
	sortedArray := arr
	slices.Sort(sortedArray)
	found := false
	low := 0
	high := len(sortedArray) - 1
	for low <= high {
		mid := (high + low) / 2
		if toSearch == sortedArray[mid] {
			found = true
			break
		}
		if toSearch > sortedArray[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return found
}

func main() {
	array := []int{3, 1, 5, 4}
	to_search1 := 3
	if !bsearch(array, to_search1) {
		fmt.Println("expected true, got false")
	} else {
		fmt.Println("ok")
	}

	to_search2 := 6
	if bsearch(array, to_search2) {
		fmt.Println("expected false, got true")
	} else {
		fmt.Println("ok")
	}
}
