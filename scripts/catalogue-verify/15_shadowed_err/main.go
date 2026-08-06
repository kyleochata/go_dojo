package main

import (
	"errors"
	"fmt"
)

func step() (int, error) {
	return 0, errors.New("step failed")
}

func run() error {
	var err error
	if true {
		v, err := step() // := shadows outer err in this block
		_ = v
		if err != nil {
			// handled locally, but outer err never set
		}
	}
	return err // still nil!
}

func main() {
	err := run()
	fmt.Println("run() err:", err)
	if err == nil {
		fmt.Println("BUG REPRODUCED: shadowed err in inner scope, outer err never set despite failure")
	}
}
