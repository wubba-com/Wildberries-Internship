package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

/**
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться
*/

func sender(c chan<- int) {
	go func() {
		for {
			r := rand.Int()
			c <- r
		}
	}()
}

func receiver4(c <-chan int) {
	go func() {
		for d := range c {
			fmt.Println(d)
		}
	}()
}

func main() {
	c := make(chan int)
	var sec uint
	flag.UintVar(&sec, "seconds", 3, "program running time in second")
	flag.Parse()

	sender(c)
	receiver4(c)

	select {
	case <-time.After(3 * time.Second):
		os.Exit(1)
	}
}
