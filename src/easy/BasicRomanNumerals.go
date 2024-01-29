package main

import "fmt"

// BasicRomanNumerals returns the decimal equivalent of the Roman numeral given
func BasicRomanNumerals(str string) int {
	// Define a map to store the values of each Roman numeral
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize the result
	result := 0

	// Iterate through the Roman numeral string
	for i := 0; i < len(str); i++ {
		// Get the value of the current Roman numeral
		value := romanValues[str[i]]

		// If the current numeral is smaller than the next numeral, subtract its value
		if i < len(str)-1 && romanValues[str[i]] < romanValues[str[i+1]] {
			result -= value
		} else {
			// Otherwise, add its value
			result += value
		}
	}

	return result
}

func main() {
	// Test cases
	result1 := BasicRomanNumerals("MDCXXI")
	fmt.Println(result1) // Output: 24

	result2 := BasicRomanNumerals("XLVI")
	fmt.Println(result2) // Output: 46
}
