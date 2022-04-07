package main

import (
	"fmt"
	"reflect"
)

/**
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}
*/

func test(v interface{}) {
	i := reflect.ValueOf(v)
	switch i.Kind() {
	case reflect.String:
		fmt.Printf("v is string - %s\n", i)
	case reflect.Int:
		fmt.Printf("v is int - %v\n", i)
	case reflect.Bool:
		fmt.Printf("v is bool - %v\n", i.Bool())
	case reflect.Chan:
		fmt.Printf("v is chan - %v\n", i.Interface())
	}
}

func main() {
	test(make(chan int))
	test(100)
	test(true)
	test("test")
}
