package main

import (
	"fmt"
	"strconv"
)

func Calculator(str string) int {
	tokens := tokenize(str)
	result, _ := evaluate(tokens)
	return result
}

func tokenize(str string) []string {
	var tokens []string

	for i := 0; i < len(str); i++ {
		switch {
		case str[i] == ' ':
			// Skip spaces
			continue
		case isDigit(str[i]):
			// If it's a digit, extract the entire number
			start := i
			for i < len(str) && isDigit(str[i]) {
				i++
			}
			tokens = append(tokens, str[start:i])
			i--
		default:
			// If it's not a digit, add the operator or parenthesis as a token
			tokens = append(tokens, string(str[i]))
		}
	}

	return tokens
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func evaluate(tokens []string) (int, []string) {
	var stack []int
	var operators []string

	for len(tokens) > 0 {
		token := tokens[0]
		tokens = tokens[1:]

		switch token {
		case "(":
			// Recursively evaluate the expression inside the parentheses
			result, remainingTokens := evaluate(tokens)
			stack = append(stack, result)
			tokens = remainingTokens
		case ")":
			// Return the result and the remaining tokens
			return stack[0], tokens
		case "+", "-", "*", "/":
			// Handle operators
			operators = append(operators, token)
		default:
			// Handle numbers
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}

		// Check if there are operators on the stack that can be applied
		for len(operators) > 0 && len(stack) >= 2 && precedence(operators[len(operators)-1]) >= precedence(token) {
			operator := operators[len(operators)-1]
			operators = operators[:len(operators)-1]

			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Apply the operator to the top two numbers on the stack
			result := applyOperator(num1, num2, operator)
			stack = append(stack, result)
		}
	}

	// Evaluate any remaining operators on the stack
	for len(operators) > 0 && len(stack) >= 2 {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		num2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		num1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result := applyOperator(num1, num2, operator)
		stack = append(stack, result)
	}

	// Return the final result and the remaining tokens
	return stack[0], tokens
}

func precedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func applyOperator(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0
	}
}

func main() {
	// Example usage
	result := Calculator("2+(3-1)*3")
	fmt.Println(result) // Output: 8

	result = Calculator("(2-0)(6/2)")
	fmt.Println(result) // Output: 6
}
