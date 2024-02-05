package main

import "fmt"

/**
 * Division function.
 *
 * @param num1 input number 1
 * @param num2 input number 2
 * @return the GCF
 */
func division(num1, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return num1 + num2
	}

	return division(num2, num1%num2)
}

func main() {
	result1 := division(20, 40)
	fmt.Println(result1)
	result2 := division(12, 16)
	fmt.Println(result2)
	result3 := division(21, 54)
	fmt.Println(result3)
}
