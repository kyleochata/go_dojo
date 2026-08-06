# Go sabotage catalogue
Pinned to: goX.XX.X   <- run `go version`, record it here, and
re-verify every entry whenever you upgrade (Sunday checklist item).
Drilling a defect the toolchain no longer produces is folklore —
loop-variable capture died in Go 1.22.

## Memory and value semantics
- slice aliasing after `append` on a shared backing array
- the result of `append` not reassigned
- a large struct copied in a `range`
- a value receiver mutating a copy
- a method set mismatch between pointer and value
- a `sync.Mutex` copied by value along with its struct

## Maps and nil
- writing to a nil map
- depending on map iteration order
- dropping the second return of a lookup
- a nil interface holding a typed nil pointer, therefore != nil

## Errors
- an error checked and swallowed
- `fmt.Errorf` without `%w`
- `==` where `errors.Is` was needed
- a deferred close whose error is discarded
- a shadowed `err` from `:=` in an inner scope

## Concurrency
- a goroutine blocked forever on a send with no receiver
- `WaitGroup.Add` called inside the goroutine rather than before it
- an unbuffered channel deadlock
- a context accepted but never propagated
- `time.After` inside a loop leaking timers
- a data race on a shared struct field
- a mutex released on one path but not another

## Resources
- `defer` inside a loop holding resources until function return
- deferred arguments evaluated at defer time when execution time was meant
