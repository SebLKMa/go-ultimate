package main

import "fmt"

func matchingCountsNotOptimised(input []string, query []string) (output []int) {
	for _, q := range query {
		count := 0
		for _, i := range input {
			if i == q {
				count++
			}
		}
		output = append(output, count)
	}
	return
}

func main() {
	input := []string{"aba", "baba", "aba", "xzxb"}
	query := []string{"aba", "xzxb", "ab"}
	counts := matchingCountsNotOptimised(input, query)
	fmt.Printf("Before %v\n", input)
	fmt.Printf("Query  %v\n", query)
	fmt.Printf("Counts %v\n", counts) // 2 1 0
}
