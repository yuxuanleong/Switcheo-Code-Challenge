package main

import "fmt"

// Iterative
func sum_to_n_a(n int) int {

	sum := 0

	for i := n; i > 0; i-- {
		sum = sum + i
	}

	return sum
}

// Recursive
func sum_to_n_b(n int) int {

	// Helper function
	var recursiveFunc func(int, int) int

	// Defining helper function
	recursiveFunc = func (n, sum int) int {

		if n == 0 {
			// Base case
			return sum
		} else {
			return recursiveFunc(n - 1, sum + n)
		}
	}

	// Calling helper function
	return recursiveFunc(n, 0)
}

// Sum of n terms of Arithmetic Progression
func sum_to_n_c(n int) int {
	
	numOfPairs := n / 2
	multiplier := 1 + n // first number + last number
	middleNum := 0

	// If n is odd
	if n % 2 != 0 {
		middleNum = (n + 1) / 2
	}
	sum := numOfPairs * multiplier + middleNum
	return sum
}

func main() {
	n := 10
	result_a := sum_to_n_a(n)
	result_b := sum_to_n_b(n)
	result_c := sum_to_n_c(n)
	fmt.Printf("n = %d\nSolution a: %d\nSolution b: %d\nSolution c: %d\n", n, result_a, result_b, result_c)
}