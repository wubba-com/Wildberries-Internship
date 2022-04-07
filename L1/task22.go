package main

import (
	"flag"
	"fmt"
	"math"
)

/**
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20
*/

func CheckNumber(x int) bool {
	if x > int(math.Pow(2, 20)) {
		return true
	}

	return false
}

func main() {
	var a int
	var b int
	flag.IntVar(&a, "a", int(math.Pow(2, 20)+1), "num a")
	flag.IntVar(&b, "b", int(math.Pow(2, 20)+1), "num b")
	flag.Parse()

	if (!CheckNumber(a)) || !CheckNumber(b) {
		fmt.Println("Small num")
		return
	}
	// При делении стоит быть внимательным, так как если в операции участвуют два целых числа, то результат деления будет округляться до целого числа
	fmt.Printf("%d \n %g \n %d \n %d", a*b, float32(a)/float32(b), a+b, a-b)
}
