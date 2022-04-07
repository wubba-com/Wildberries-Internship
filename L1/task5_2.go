package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться
*/

// горутина отправитель
func sender2(c chan<- int) {
	for {
		r := rand.Int()
		c <- r
	}
}

// горутина получатель
func receiver5(c <-chan int) {
	for d := range c {
		fmt.Println(d)
	}
}

func main() {
	c := make(chan int)
	ctxWithCancel, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer func() {
		cancel()
	}()

	func(ctx context.Context, c chan int) {
		go receiver5(c)
		go sender2(c)

		select {
		// канал закроется спустя 2 сек и программа войдет в этот case
		case <-ctxWithCancel.Done():
			fmt.Println("exit")
		}
	}(ctxWithCancel, c)
}
