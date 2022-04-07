package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

// Решение задачи с множеством горутин
func receiver(c <-chan int) {
	for {
		n := <-c
		fmt.Fprintf(os.Stdout, "square: %d\n", n)
	}
}

func main() {
	ch := make(chan int)
	numbers := []int{2, 4, 6, 8, 10, 12, 14}

	// для ожидания основной горутины пока другие горутины продолжают работу
	wg := sync.WaitGroup{}

	// запускается на прием из потока
	go receiver(ch)

	for _, n := range numbers {
		wg.Add(1)
		// воркеры выполняющие вычисления
		go func(n int, c chan<- int) {
			defer wg.Done()
			square := math.Pow(float64(n), 2)
			ch <- int(square)
		}(n, ch)
	}

	wg.Wait()
}
