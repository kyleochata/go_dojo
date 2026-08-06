package main

import "fmt"

type Counter struct{ n int }

func (c Counter) Inc() { // value receiver: mutates a copy
	c.n++
}

func main() {
	c := Counter{n: 0}
	c.Inc()
	fmt.Println("n:", c.n)
	if c.n == 0 {
		fmt.Println("BUG REPRODUCED: value receiver mutated a copy only")
	} else {
		fmt.Println("not reproduced")
	}
}
