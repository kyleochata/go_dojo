# Idiom 02 — set via `map[T]struct{}`

## Canonical

```go
package idiom

// Inline form — what you actually type mid-problem.
func dedup(vals []int) []int {
	seen := make(map[int]struct{}, len(vals))
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

// Named form — when the set outlives one loop.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := make(Set[T], len(vals))
	for _, v := range vals {
		s.Add(v)
	}
	return s
}

func (s Set[T]) Add(v T)       { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool  { _, ok := s[v]; return ok }
func (s Set[T]) Del(v T)       { delete(s, v) }
func (s Set[T]) Len() int      { return len(s) }
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. `struct{}{}` vs `struct{}` — which is the type and which is the
   value? Write the composite literal from memory.
2. Why `map[T]struct{}` over `map[T]bool`? Name the concrete cost
   difference, and the one place `map[T]bool` is genuinely nicer.
3. `Add` and `Del` use value receivers on a map type and still
   mutate. Explain why — and name the one operation on `Set[T]` that
   a value receiver *cannot* make stick.

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
