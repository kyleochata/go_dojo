package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Safe struct {
	mu  sync.Mutex
	val int
}

func demo() {
	s := Safe{}
	s.mu.Lock()
	s.val = 1
	s.mu.Unlock()

	cp := s // copies the Mutex by value
	cp.mu.Lock()
	cp.val = 2
	cp.mu.Unlock()
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-child" {
		demo()
		return
	}

	out, _ := exec.Command("go", "vet", "main.go").CombinedOutput()
	fmt.Print(string(out))
	if strings.Contains(string(out), "copies lock value") {
		fmt.Println("BUG REPRODUCED: go vet still flags a sync.Mutex copied by value (copylocks)")
	} else {
		fmt.Println("STALE: go vet no longer flags this copylocks pattern")
	}
}
