package main

import (
	"errors"
	"fmt"
)

/**
Удалить i-ый элемент из слайса
*/

// Remove можно убрать проверку на max i, тогда за несуществующий индекс будем получать панику
func Remove(i int, sl []int) ([]int, error) {
	if i > len(sl)-1 {
		return nil, errors.New("i > last index slice")
	}
	return append(sl[:i], sl[i+1:]...), nil
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	a, err := Remove(10, a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(a)
}
