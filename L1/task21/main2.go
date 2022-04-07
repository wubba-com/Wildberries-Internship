package main

import (
	"fmt"
	"math"
)

// PowerFunc - Адаптер, который позволяет использовать производить вычисления с одним числом
type PowerFunc func(x int)

func (f PowerFunc) Call(x int) {
	f(x)
}

func Count(x int, f func(x int)) {
	PowerFunc(f).Call(x)
}

// Конкретные функции, работающие с числами
func test(x int) {
	fmt.Println(math.Pow(float64(2), float64(x)))
}

func test2(x int) {
	fmt.Println(math.Sqrt(float64(x)))
}

func main() {
	Count(2, test)
	Count(2, test2)
}
