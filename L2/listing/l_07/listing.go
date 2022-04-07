package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
	Loop:
		for {
			select {
			case v, ok := <-a:
				if !ok {
					break Loop
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					break Loop
				}
				c <- v
			}
		}
		close(c)
	}()
	return c
}

func merge2(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for v := range a {
			c <- v
		}

		for v := range b {
			c <- v
		}

		close(c)
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)

	for v := range c {
		fmt.Println(v)
	}
}
