package main

import (
	"fmt"
	"strconv"
)

func parse(s string) int {
	n, _ := strconv.Atoi(s) // error checked-and-discarded via blank identifier
	return n
}

func main() {
	n := parse("not-a-number")
	fmt.Println("n:", n)
	if n == 0 {
		fmt.Println("BUG REPRODUCED: invalid input silently became zero value, error swallowed")
	}
}
