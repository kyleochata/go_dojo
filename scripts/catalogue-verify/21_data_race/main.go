package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Shared struct{ count int }

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-child" {
		s := &Shared{}
		var wg sync.WaitGroup
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.count++ // unsynchronized read-modify-write from many goroutines
			}()
		}
		wg.Wait()
		fmt.Println("final count:", s.count)
		return
	}

	out, _ := exec.Command("go", "run", "-race", "main.go", "-child").CombinedOutput()
	fmt.Print(string(out))
	if strings.Contains(string(out), "WARNING: DATA RACE") {
		fmt.Println("BUG REPRODUCED: data race on shared struct field")
	} else {
		fmt.Println("INCONCLUSIVE: race detector did not flag a race this run")
	}
}
