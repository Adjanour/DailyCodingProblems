package day09

import (
	"testing"
)

func TestSolution(t *testing.T) {

	// Helper function to run each test case
	testCase := func(name string, input []int, expected int) {

		result := largestSumNonAdjacent(input)

		if result != expected {
			t.Errorf("%s: got %d, want %d", name, result, expected)
		}

		// fmt.Printf("%s:\n", name)
		// fmt.Printf("  Input:     %v\n", input)
		// fmt.Printf("  Expected:  %d\n", expected)
		// fmt.Printf("  Returned:  %d %v\n", result, result == expected)
		// fmt.Println()
	}

	testCase("Empty list", []int{}, 0)
	testCase("Single element", []int{5}, 5)
	testCase("Two elements", []int{4, 10}, 10)
	testCase("Three elements", []int{3, 2, 5}, 8)
	testCase("All positives", []int{2, 7, 9, 3, 1}, 12)
	testCase("All negatives", []int{-1, -2, -3}, 0)
	testCase("Mixed numbers", []int{4, -1, 1, 10, -3, 2}, 16)
}
