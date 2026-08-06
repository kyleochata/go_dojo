package main

import (
	"errors"
	"fmt"
)

type failingCloser struct{}

func (failingCloser) Close() error { return errors.New("flush failed") }

func writeFile() (err error) {
	c := failingCloser{}
	defer c.Close() // error return value discarded
	return nil
}

func main() {
	err := writeFile()
	fmt.Println("writeFile err:", err)
	if err == nil {
		fmt.Println("BUG REPRODUCED: Close() error was silently discarded by bare defer")
	}
}
