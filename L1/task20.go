package main

import (
	"fmt"
	"strings"
)

/**
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»
*/

func reverseWorld(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWorld("snow dog sun"))
}
