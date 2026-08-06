package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-child" {
		a := make(chan int)
		b := make(chan int)
		<-a
		<-b
		return
	}

	out, _ := exec.Command(os.Args[0], "-child").CombinedOutput()
	fmt.Print(string(out))
	if strings.Contains(string(out), "all goroutines are asleep - deadlock") {
		fmt.Println("BUG REPRODUCED: unbuffered channel deadlock")
	} else {
		fmt.Println("INCONCLUSIVE: expected runtime deadlock detection did not fire")
	}
}
