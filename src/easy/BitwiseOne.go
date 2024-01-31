package main

import "fmt"

func BitwiseOne(arr []string) string{
	s1 := arr[0]
	s2 := arr[1]

	result := ""

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			result += string(s1[i])
		}else{
			result += string("1")
		}
	}

	return result
}

func main(){
	// Test cases 
	result1 := BitwiseOne([]string{"001110", "100000"})
	fmt.Println(result1)
	result2 := BitwiseOne([]string{"110001", "111100"})
	fmt.Println(result2)
}