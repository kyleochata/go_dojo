package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func lookup() error {
	return fmt.Errorf("lookup failed: %s", ErrNotFound) // %s instead of %w
}

func main() {
	err := lookup()
	fmt.Println("err:", err)
	fmt.Println("errors.Is(err, ErrNotFound):", errors.Is(err, ErrNotFound))
	if !errors.Is(err, ErrNotFound) {
		fmt.Println("BUG REPRODUCED: %s instead of %w breaks the wrapped-error chain")
	} else {
		fmt.Println("not reproduced")
	}
}
