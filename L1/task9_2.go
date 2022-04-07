package main

import (
	"fmt"
	"os"
	"time"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func worker10(c chan<- int, n int) {
	c <- n
}

func sq2(c, c2 chan int) {
	for n := range c {
		c2 <- n * 2
	}
}

func main() {
	c := make(chan int)
	c2 := make(chan int)
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range nums2 {
		go worker10(c, n)
	}

	go sq2(c, c2)

	go func(c <-chan int) {

		for n := range c {
			fmt.Println(n)
		}
	}(c2)

	select {
	case <-time.After(time.Second):
		os.Exit(1)
	}
}
