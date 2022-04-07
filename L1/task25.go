package main

import (
	"fmt"
	"time"
)

/**
Реализовать собственную функцию sleep
*/

func Sleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C
}

func main() {
	Sleep(2 * time.Second)
	fmt.Println("end")
}
