package main

import (
"time"
)

func main() {
	c := make(chan bool)

	go func() {
		for {
			select {
			case <-c:
				return
			}
		}
	}()
	// …

	c <- true
	time.Sleep(1 * time.Second)
}

