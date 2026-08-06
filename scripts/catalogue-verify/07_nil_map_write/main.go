package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
			fmt.Println("BUG REPRODUCED: write to nil map panics")
		}
	}()

	var m map[string]int
	m["x"] = 1
	fmt.Println("no panic — not reproduced")
}
