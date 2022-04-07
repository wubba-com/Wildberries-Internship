package main

import (
	"fmt"
	"math"
)

/**
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений
*/

func f1(nums []int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func f2(c <-chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- int(math.Pow(float64(n), 2))
		}
		close(out)
	}()

	return out
}

func main() {
	var sum int
	numbers := []int{2, 4, 6, 8, 10}
	out := f1(numbers)
	out2 := f2(out)

	for n := range out2 {
		sum += n
	}
	fmt.Println(sum)
}
