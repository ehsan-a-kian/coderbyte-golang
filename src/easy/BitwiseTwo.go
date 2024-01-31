package main

import "fmt"

func BitwiseTwo(arr []string) string {
	s1 := arr[0]
	s2 := arr[1]

	result := ""

	for i := 0; i < len(s1); i++ {
		if string(s1[i]) == "1" && string(s2[i]) == "1" {
			result += string("1")
		} else {
			result += string("0")
		}
	}

	return result
}

func main() {

	// Test cases
	result1 := BitwiseTwo([]string{"010111", "101101"})
	fmt.Println(result1)
	result2 := BitwiseTwo([]string{"110001", "111100"})
	fmt.Println(result2)
}
