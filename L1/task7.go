package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map
*/

func main() {
	m := make(map[string]int)
	mut := sync.Mutex{}
	i := 0

	for i < 10 {
		go func() {
			mut.Lock()
			defer mut.Unlock()
			m["counter"]++
		}()
		i++
	}

	fmt.Println(m)

}
