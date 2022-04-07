package main

import (
	"fmt"
	"sync"
)

// Копируют значение блокировки
// Если не передавать по указателю, то счетчик не изменяет значение
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

