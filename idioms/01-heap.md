# Idiom 01 — `container/heap` with a concrete type

## Canonical

```go
package idiom

import "container/heap"

// IntHeap is a min-heap. Value receivers for the sort.Interface
// methods, pointer receivers for Push/Pop — they resize the slice.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func drain(vals []int) []int {
	h := &IntHeap{}
	*h = append(*h, vals...)
	heap.Init(h)
	heap.Push(h, 42)

	out := make([]int, 0, h.Len())
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(int))
	}
	return out
}
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. Which of the five methods need a pointer receiver, and why does
   the split matter? What happens if you give `Push` a value receiver?
2. `heap.Pop` pops from the *end* of the slice while `Pop` on the
   heap returns the *minimum*. Who reconciles those, and what does
   `heap.Pop` do before calling your `Pop`?
3. Turn this into a max-heap without touching `Push`, `Pop`, or the
   element type. Then: what changes if elements are structs with a
   priority field?

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
