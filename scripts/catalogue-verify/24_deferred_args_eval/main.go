package main

import "fmt"

var observed int

func main() {
	x := 1
	defer func(captured int) {
		observed = captured
		fmt.Println("deferred arg evaluated at defer-time, x was:", captured)
		if captured == 1 {
			fmt.Println("BUG REPRODUCED: deferred call used x==1 (defer-time value), not the x==2 set afterward")
		} else {
			fmt.Println("INCONCLUSIVE: deferred arg picked up the later value")
		}
	}(x) // x evaluated and copied NOW, at defer time
	x = 2
	fmt.Println("x at end of main:", x)
}
