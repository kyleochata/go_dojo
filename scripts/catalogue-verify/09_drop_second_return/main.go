package main

import "fmt"

func main() {
	m := map[string]int{"a": 0}
	v := m["missing"] // zero value looks legit, no way to tell absent vs present-zero
	fmt.Println("v:", v)
	if v == 0 {
		fmt.Println("BUG REPRODUCED: dropped ok bool, can't distinguish absent key from zero value")
	}
}
