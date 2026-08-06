package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-child" {
		ch := make(chan int) // unbuffered, no receiver anywhere
		ch <- 1               // deadlocks
		return
	}

	out, _ := exec.Command(os.Args[0], "-child").CombinedOutput()
	fmt.Print(string(out))
	if strings.Contains(string(out), "all goroutines are asleep - deadlock") {
		fmt.Println("BUG REPRODUCED: goroutine blocked forever on a send with no receiver")
	} else {
		fmt.Println("INCONCLUSIVE: expected runtime deadlock detection did not fire")
	}
}
