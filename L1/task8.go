package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

// BitDisjunction Поразрядная дизъюнкция (OR +)
func BitDisjunction(i int64, x int64) int64 {
	z := x | 1<<i
	fmt.Printf("%b - %d", z, z)
	return z
}

func main() {
	var x int64 = 20
	BitDisjunction(5, x)
}
