package main

import (
	"fmt"
)

// ArithGeo takes the array of numbers stored in arr and returns "Arithmetic" if the sequence follows an arithmetic pattern,
// "Geometric" if it follows a geometric pattern, or -1 if neither pattern is found.
func ArithGeo(arr []int) string {
	arithInterval := arr[1] - arr[0]
	geoInterval := arr[1] / arr[0]
	arithCount := 0
	geoCount := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == arithInterval {
			arithCount++
		}
		if arr[i+1]/arr[i] == geoInterval {
			geoCount++
		}
	}

	if arithCount == len(arr)-1 {
		return "Arithmetic"
	}

	if geoCount == len(arr)-1 {
		return "Geometric"
	}

	return "-1"
}

func main() {
	result1 := ArithGeo([]int{2, 4, 6, 8})
	fmt.Println(result1)
	result2 := ArithGeo([]int{2, 6, 18, 54})
	fmt.Println(result2)
	result3 := ArithGeo([]int{-3, -4, -5, -6, -7})
	fmt.Println(result3)
	result4 := ArithGeo([]int{3, 12, 48, 192, 768})
	fmt.Println(result4)
}
