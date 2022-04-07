package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

/**
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout
*/

// Решение задачи с 2-мя горутинами
var wg = &sync.WaitGroup{}

func worker(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		n := <-ch
		fmt.Fprintf(os.Stdout, "square: %d\n", n)
	}
}

func main() {
	ch := make(chan int)
	numbers := []int{2, 4, 6, 8, 10, 12, 14}

	// запускается на чтение и блокируется
	// если перенести за цикл будет дедлок, т.к горутина main заблокируется после первой записи
	wg.Add(1)
	go worker(ch, wg)

	// основная горутина начинает писать в канал
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for _, n := range numbers {
			square := math.Pow(float64(n), 2)
			ch <- int(square)
		}
	}(wg)

	wg.Wait()
}
