package main

import (
	"fmt"
	"sort"
)

/**
Реализовать быструю сортировку массива (quicksort) встроенными методами языка
*/

// Sort, SliceStable

type Nums []int

func (n Nums) Search(x int) int {
	i := sort.Search(len(n), func(i int) bool {
		return x <= n[i]
	})

	return i
}

func main() {
	nums := []int{4, 3, 7, 1, 3, 10, 15, 12, 6}
	// Для работы с int
	sort.Ints(nums)
	fmt.Println(nums)

	fmt.Println(Nums(nums).Search(123))
}
