package main

import "fmt"

//func index(s string, w []string) int {
//	for i, v := range w {
//		if s == v {
//			return i
//		}
//	}
//
//	return -1
//}
//
//func Set(w []string) []string {
//	set := make([]string, 0)
//	for _, s := range w {
//		if index(s, set) < 0 {
//			set = append(set, s)
//		}
//	}
//
//	return set
//}

func Set(sl []string) []string  {
	m := make(map[string]int)
	set := make([]string, 0)
	for _, v := range sl {
		m[v] += 1
	}

	for k := range m {
		if m[k] < 2 {
			set = append(set, k)
		}
	}

	return set
}

func main() {
	words := []string{"fox", "cat", "dog", "cat", "tree"}
	fmt.Println(Set(words))
}
