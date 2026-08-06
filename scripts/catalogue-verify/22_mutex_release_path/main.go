package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu sync.Mutex
	v  int
}

func (s *Store) update(fail bool) error {
	s.mu.Lock()
	if fail {
		return fmt.Errorf("early return without unlocking") // BUG: no unlock on this path
	}
	s.v++
	s.mu.Unlock()
	return nil
}

func main() {
	s := &Store{}
	_ = s.update(true) // takes the buggy early-return path, mutex stays locked

	done := make(chan struct{})
	go func() {
		s.mu.Lock() // will block forever
		s.mu.Unlock()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("not reproduced: second lock acquired")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("BUG REPRODUCED: mutex never released on the error path, second Lock() blocked")
	}
}
