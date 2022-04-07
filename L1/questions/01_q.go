package main

import (
	"bytes"
	"fmt"
)

// 1. Какой самый эффективный способ конкатенации строк?
// Строка - это неизменяемая последовательность слайс байт доступные только для чтения

func main()  {
	s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
	fmt.Printf(s)

	strings := []string{"This ", "is ", "even ", "more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}
