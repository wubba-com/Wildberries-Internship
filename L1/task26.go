package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой
*/

func Unique2(input string) bool {
	if utf8.RuneCountInString(input) < 2 {
		return true
	}

	input = strings.ToLower(input)
	m := make(map[string]int)
	for _, b := range input {
		m[string(b)] += 1
		if v, ok := m[string(b)]; ok {
			if v > 1 {
				return false
			}
		}
	}

	return true
}

func Unique(input string) bool {
	if utf8.RuneCountInString(input) < 2 {
		return true
	}
	input = strings.ToLower(input)
	// независимый
	b := []byte(input)
	buf := []byte(input)
	for i, v := range b {
		for _, v2 := range append(buf[:i], buf[i+1:]...) {
			if string(v) == string(v2) {
				return false
			}
		}
		buf = []byte(input)
	}

	return true
}

func main() {
	// Unique
	fmt.Printf("unique: %t\n", Unique("abcd"))
	fmt.Printf("unique: %t\n", Unique("abCdefAaf"))
	fmt.Printf("unique: %t\n", Unique("aabcd"))
	fmt.Printf("unique: %t\n", Unique(""))
	fmt.Printf("unique: %t\n", Unique("f"))

	fmt.Println()

	// Unique2
	fmt.Printf("unique: %t\n", Unique2("abcd"))
	fmt.Printf("unique: %t\n", Unique2("abCdefAaf"))
	fmt.Printf("unique: %t\n", Unique2("aabcd"))
	fmt.Printf("unique: %t\n", Unique2(""))
	fmt.Printf("unique: %t\n", Unique2("f"))

}
