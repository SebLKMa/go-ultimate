package main

import "fmt"

// This is a solution to
// https://www.hackerrank.com/challenges/grading/problem

func nextHigherMultipleOf(in int, multiple int) (out int) {
	if in <= 0 {
		return in
	}

	multipleOf := 0
	for i := 1; multipleOf < in; i++ {
		multipleOf = i * multiple
	}

	return multipleOf
}

func modHigherMultipleOf(in int, multiple int) (out int) {
	if in <= 0 {
		return in
	}

	mod := in % multiple
	multipleOf := in + (5 - mod)

	return multipleOf
}

func getGrade(input int, ceiling int, threshold int, failing int) (output int) {
	if input < failing {
		return input
	}
	if (ceiling - input) >= threshold {
		return input
	}
	return ceiling
}

func computeRoundedGrades(grades []int, threshold int, failing int) (results []int) {
	for _, g := range grades {
		magic := modHigherMultipleOf(g, 5)
		rounded := getGrade(g, magic, threshold, failing)
		results = append(results, rounded)
	}
	return
}

// This is a solution to
// https://www.hackerrank.com/challenges/grading/problem
func main() {
	threshold := 3
	failing := 38
	inputNumber := 73
	magicNumber := nextHigherMultipleOf(inputNumber, 5)
	fmt.Printf("%d nextHigherMultipleOf is %d, grade is %d\n", inputNumber, magicNumber, getGrade(inputNumber, magicNumber, threshold, failing))
	inputNumber = 32
	magicNumber = nextHigherMultipleOf(inputNumber, 5)
	fmt.Printf("%d nextHigherMultipleOf is %d, grade is %d\n", inputNumber, magicNumber, getGrade(inputNumber, magicNumber, threshold, failing))

	originalGrades := []int{73, 67, 38, 33}
	roundedGrades := computeRoundedGrades(originalGrades, threshold, failing)
	fmt.Println("Original Grades:")
	for _, g := range originalGrades {
		fmt.Println(g)
	}
	fmt.Println("Rounded Grades:")
	for _, g := range roundedGrades {
		fmt.Println(g)
	}
}
