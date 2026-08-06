package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	results := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			wg.Add(1) // race: Wait() below may run before this executes
			defer wg.Done()
			results <- n
		}(i)
	}

	wg.Wait() // may return immediately since counter can still be 0
	close(results)

	count := 0
	for range results {
		count++
	}
	fmt.Println("collected:", count, "of 10")
	if count < 10 {
		fmt.Println("BUG REPRODUCED: Wait() returned before all Add() calls landed, results dropped")
	} else {
		fmt.Println("did not reproduce this run (race is timing-dependent)")
	}
}
