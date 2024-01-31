package main

import (
	"fmt"
	"regexp"
	"strings"
)

// CamelCase converts the input string to camel case format.
func CamelCase(str string) string {
	words := splitWords(str)
	var camel strings.Builder
	for i, word := range words {
		if i == 0 {
			camel.WriteString(word)
		} else {
			camel.WriteString(strings.ToUpper(word[:1]) + word[1:])
		}
	}
	return camel.String()
}

// splitWords is a helper function which splits the input string into words.
func splitWords(str string) []string {
	str = strings.ToLower(str)
	re := regexp.MustCompile("[^a-z]+")
	words := re.Split(str, -1)
	var result []string
	for _, word := range words {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	result1 := CamelCase("_good_looking_blues_")
	fmt.Println(result1)

	result2 := CamelCase("-=last-night-on-earth=-")
	fmt.Println(result2)
}
