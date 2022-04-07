package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

/**
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте
*/

func generateGo(n uint, c <-chan int) {
	var i uint
	for i < n {
		go func(i uint) {
			for j := range c {
				_, err := fmt.Fprintf(os.Stdout, "%d\n", j)
				if err != nil {
					return
				}
			}
		}(i)
		i++
	}
}

func main() {
	osSign := make(chan os.Signal)
	signal.Notify(osSign, syscall.SIGINT, syscall.SIGTERM) // SIGINT - сигнал прерывания

	c := make(chan int)
	var n uint
	flag.UintVar(&n, "workers", 10, "amount workers")
	flag.Parse()

	generateGo(n, c)

	go func() {
		for {
			r := rand.Int()
			c <- r
		}
	}()

	select {
	case <-osSign:
		fmt.Println("Exit")
	}
}
