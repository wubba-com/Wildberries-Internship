package main

import "fmt"

/**
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Sex  string
	Age  int
}

type Action struct {
	Human
}

func (h *Human) Walk() {
	fmt.Println(h.Name + " is walking...")
}

func main() {
	a := Action{}
	fmt.Printf("sex: %s\n", a.Sex)
	a.Walk()
}
