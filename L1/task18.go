package main

import (
	"fmt"
	"sync"
)

/**
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое
*/

type Incrementer struct {
	Inc int
	mu  sync.Mutex
}

func (i *Incrementer) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	i.mu.Lock()
	defer i.mu.Unlock()
	i.Inc++
}

func main() {
	var i int
	wg11 := &sync.WaitGroup{}
	inc := &Incrementer{mu: sync.Mutex{}}
	for i < 100000 {
		wg11.Add(1)
		go inc.Increment(wg11)
		i++
	}
	wg11.Wait()
	fmt.Println("inc:", inc.Inc)
}
