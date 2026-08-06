package main

import "fmt"

type Big struct {
	data [4]int
	tag  string
}

func main() {
	items := []Big{{data: [4]int{1, 2, 3, 4}, tag: "a"}}

	for _, item := range items {
		item.tag = "mutated" // mutates the loop copy, not items[0]
	}

	fmt.Println("items[0].tag:", items[0].tag)
	if items[0].tag == "a" {
		fmt.Println("BUG REPRODUCED: range gave a copy, original untouched")
	} else {
		fmt.Println("not reproduced")
	}
}
