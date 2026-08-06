package main

import (
	"context"
	"fmt"
	"time"
)

func slowWork(ctx context.Context) error {
	// BUG: ctx is accepted but never used to bound the work
	time.Sleep(200 * time.Millisecond)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := slowWork(ctx)
	elapsed := time.Since(start)

	fmt.Println("elapsed:", elapsed, "err:", err)
	if elapsed >= 200*time.Millisecond {
		fmt.Println("BUG REPRODUCED: work ran to completion despite a 20ms context timeout — ctx never propagated")
	} else {
		fmt.Println("not reproduced")
	}
}
