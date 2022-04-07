package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка

func Search(sl []int, x int) int {
	sort.Ints(sl)
	return sort.Search(len(sl), func(i int) bool { return x <= sl[i] })
}

func main() {
	a := []int{6, 1, 2, 7, 3, 9}
	x := 3
	i := Search(a, x)
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено %d по индексу %d в %v.\n", x, i, a)
	} else {
		fmt.Printf("Не найдено %d в %v.\n", x, a)
	}
}
