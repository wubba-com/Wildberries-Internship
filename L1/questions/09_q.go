package main

import "fmt"

func main()  {
	// slice
	sl := make([]string, 0)
	sl2 := []string{"test"}
	sl3 := sl[:]
	var sl4 []string // nil, not init

		fmt.Println(sl, sl2, sl3, sl4 == nil)
	// map
	var x map[string]int // nil, not init
	x2 := make(map[string]int) // init

	fmt.Println(x == nil, x, x2)
}
