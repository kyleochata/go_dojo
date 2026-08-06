package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var before, afterAlloc, afterGC runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&before)

	const n = 200000
	for i := 0; i < n; i++ {
		select {
		case <-time.After(time.Hour): // never fires in this test; channel goes unreferenced each iteration
		default:
		}
	}
	runtime.ReadMemStats(&afterAlloc)

	runtime.GC()
	runtime.ReadMemStats(&afterGC)

	fmt.Printf("HeapObjects before: %d\n", before.HeapObjects)
	fmt.Printf("HeapObjects right after loop (pre-GC): %d\n", afterAlloc.HeapObjects)
	fmt.Printf("HeapObjects after explicit GC: %d\n", afterGC.HeapObjects)

	reclaimed := afterAlloc.HeapObjects > afterGC.HeapObjects &&
		float64(afterAlloc.HeapObjects-afterGC.HeapObjects) > 0.5*float64(afterAlloc.HeapObjects-before.HeapObjects)

	if reclaimed {
		fmt.Println("STALE ON THIS TOOLCHAIN: unreferenced time.After timers were reclaimed by GC before firing (Go 1.23+ runtime timer GC). The classic 'leaks until it fires' framing overstates the danger now; it's still wasteful/anti-pattern but not an unbounded leak.")
	} else {
		fmt.Println("BUG REPRODUCED AS CLASSICALLY DESCRIBED: pending timers were not reclaimed, they'd persist until firing")
	}
}
