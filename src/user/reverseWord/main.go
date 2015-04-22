package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "this is my book"
	secondInput := strings.Split(reverse(input), " ")
	fmt.Println(secondInput)
	for _, c := range secondInput {
		fmt.Print(reverse(c) + " ")
	}
}

func reverse(s string) string {
	if len(s) <= 0 {
		return "error"
	}

	result := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		result[i], result[len(s)-1-i] = result[len(s)-1-i], result[i]
	}
	return string(result)
}