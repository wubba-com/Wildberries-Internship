package main

import (
	"fmt"
)

/**
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}
*/

func test2(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("v is string\n")
	case int:
		fmt.Printf("v is int\n")
	case bool:
		fmt.Printf("v is bool\n")
	case chan int, chan string, chan bool:
		fmt.Printf("v is chan\n")

	default:
		fmt.Println("type unknown")
	}
}

func main() {
	test2(make(chan bool))
	test2(100)
	test2(true)
	test2("test")
}
