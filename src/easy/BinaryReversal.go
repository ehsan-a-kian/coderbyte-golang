package main

// TO-DO:
//	1: Convert to binary
//	2: pad to N * 8
// 	3: Reverse the binary
// 	4: Convert to decimal

import (
	"fmt"
	"strconv"
)

func BinaryReversal(val int64) int64 {
	binVal := strconv.FormatInt(val, 2)

	for(len(binVal)<8) {
		binVal = string("0") + binVal
	}

	result := ""
	for _,v := range binVal {
		result = string(v) + result
	}

	strVal, _ := strconv.ParseInt(result, 2, 64)
	fmt.Println(result)
	return strVal
}

func main(){
	// Test cases
	result1 := BinaryReversal(47)
	fmt.Println(result1)
}