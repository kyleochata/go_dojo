package main

import "fmt"

func main() {
	m := map[int]int{}
	for i := 0; i < 10; i++ {
		m[i] = i
	}

	orders := map[string]bool{}
	for run := 0; run < 5; run++ {
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		orders[fmt.Sprint(keys)] = true
	}

	fmt.Println("distinct orderings observed:", len(orders))
	if len(orders) > 1 {
		fmt.Println("BUG REPRODUCED: map iteration order is not stable across runs")
	} else {
		fmt.Println("not reproduced this run (order is still randomized by spec, but happened to match — Go explicitly does not guarantee order)")
	}
}
