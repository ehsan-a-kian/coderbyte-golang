package main

import (
	"fmt"
	"sort"
)

// ArrayAdditionOne checks if any combination of numbers in the array (excluding the largest number) can be added up to equal the largest number
func ArrayAdditionOne(arr []int) string {
	// Find the largest number in the array
	sort.Ints(arr)
	largest := arr[len(arr)-1]

	// Recursive function to find the sum of combinations
	var sumCombination func(int, int) bool
	sumCombination = func(target, index int) bool {
		// Base case: If target is 0, combination found
		if target == 0 {
			return true
		}
		// Base case: If index is out of range or target is negative, no combination found
		if index < 0 || target < 0 {
			return false
		}
		// Try including the current number in the sum
		include := sumCombination(target-arr[index], index-1)
		// Try excluding the current number from the sum
		exclude := sumCombination(target, index-1)
		// Return true if either including or excluding the current number results in a valid combination
		return include || exclude
	}

	// Start exploring combinations, excluding the largest number
	return map[bool]string{true: "true", false: "false"}[sumCombination(largest, len(arr)-2)]
}

func main() {
	result1 := ArrayAdditionOne([]int{4, 6, 23, 10, 1, 3})
	fmt.Println(result1)
	result2 := ArrayAdditionOne([]int{1,2,3, 5 ,7,8, 10})
	fmt.Println(result2)
}
