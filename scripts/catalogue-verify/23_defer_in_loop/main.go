package main

import "fmt"

type resource struct{ id int }

var openCount int

func open(id int) *resource {
	openCount++
	return &resource{id: id}
}

func (r *resource) close() {
	openCount--
}

func processAll(n int) int {
	for i := 0; i < n; i++ {
		r := open(i)
		defer r.close() // all n resources stay open until processAll returns, not each iteration
		_ = r
	}
	return openCount // observed just before any defer runs
}

func main() {
	const n = 5
	peak := processAll(n)
	fmt.Println("open resources DURING loop's last iteration, before any defer runs:", peak)
	fmt.Println("open resources AFTER processAll returns:", openCount)
	if peak == n {
		fmt.Println("BUG REPRODUCED: defer inside the loop held all resources open until function return")
	} else {
		fmt.Println("INCONCLUSIVE: resources were released incrementally, not held to function return")
	}
}
