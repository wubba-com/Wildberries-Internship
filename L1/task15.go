package main

import (
	"fmt"
	"unicode/utf8"
)

/**
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
*/

var justString string

func someFunc() {
	n, v := createHugeString(1 << 9)
	justString = v[:n]
	fmt.Println(utf8.RuneCountInString(justString))
}

func createHugeString(i int) (int, string) {
	s := ""
	for i > 0 {
		s += "b"
		i--
	}
	return utf8.RuneCountInString(s), s
}

func main() {
	someFunc()
}
