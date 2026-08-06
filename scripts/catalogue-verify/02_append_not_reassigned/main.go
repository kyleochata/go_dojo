package main

import "fmt"

func addOne(s []int) {
	_ = append(s, 1) // result discarded — pointless without assignment
}

func main() {
	s := make([]int, 0, 1)
	addOne(s)
	fmt.Println("len after addOne:", len(s))
	if len(s) == 0 {
		fmt.Println("BUG REPRODUCED: caller's slice unchanged, append result was dropped")
	} else {
		fmt.Println("not reproduced")
	}
}
