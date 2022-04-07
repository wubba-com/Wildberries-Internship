package main

import (
	"time"
)

func Sleep2(duration time.Duration)  {
	<-time.After(duration)
}

func main()  {
	Sleep2(5*time.Second)
}
