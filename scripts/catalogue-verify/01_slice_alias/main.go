package main

import "fmt"

func main() {
	base := make([]int, 3, 5)
	base[0], base[1], base[2] = 1, 2, 3

	a := base[:2]
	b := append(a, 99) // has capacity, so writes into base's backing array

	fmt.Println("base:", base)
	fmt.Println("b:   ", b)
	if base[2] == 99 {
		fmt.Println("BUG REPRODUCED: append aliased base's backing array")
	} else {
		fmt.Println("not reproduced")
	}
}
