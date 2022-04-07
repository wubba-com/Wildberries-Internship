package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		// здесь слайс будет ссылаться уже на другой массив, в котором и будут происходить изменения
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)

	// slice - ссылается на старый массив
	fmt.Print(slice)
}

