package main

import (
	"fmt"
	"strings"
)

// CamelCase converts the input string to camel case format.
func CamelCase(str string) string {
	words := strings.FieldsFunc(strings.ToLower(str), func(r rune) bool {
		return !('a' <= r && r <= 'z')
	})
	var camel string
	for i, word := range words {
		if i == 0 {
			camel += word
		} else {
			camel += strings.Title(word)
		}
	}
	return camel
}

func main() {
	result1 := CamelCase("_good_looking_blues_")
	fmt.Println(result1)

	result2 := CamelCase("-=last-night-on-earth=-")
	fmt.Println(result2)
}
