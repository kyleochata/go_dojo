package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func lookup() error {
	return fmt.Errorf("lookup failed: %w", ErrNotFound)
}

func main() {
	err := lookup()
	fmt.Println("err == ErrNotFound:", err == ErrNotFound)
	fmt.Println("errors.Is(err, ErrNotFound):", errors.Is(err, ErrNotFound))
	if err != ErrNotFound && errors.Is(err, ErrNotFound) {
		fmt.Println("BUG REPRODUCED: == fails on a wrapped error where errors.Is succeeds")
	} else {
		fmt.Println("not reproduced")
	}
}
