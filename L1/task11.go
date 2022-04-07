package main

import "fmt"

func IntersectionSet(a, b []string) []string {
	m := make(map[string]int)
	intersection := make([]string, 0)
	for _, v := range a {
		m[v] += 1
	}

	for _, v := range b {
		m[v] += 1
	}

	for k := range m {
		if m[k] > 1 {
			intersection = append(intersection, k)
		}
	}
	return intersection
}

func main()  {
	a := []string{"a", "b", "c", "d", "e", "g"}
	b := []string{"a", "c", "f", "g", "t"}
	fmt.Println(IntersectionSet(a, b))
}
