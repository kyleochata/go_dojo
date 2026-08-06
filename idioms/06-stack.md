# Idiom 06 — stack

## Canonical

```go
package idiom

// Inline form — what you actually type mid-problem.
func balanced(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// Named form — pointer receivers, because push and pop resize.
type stack []int

func (s *stack) push(v int) { *s = append(*s, v) }

func (s *stack) pop() (int, bool) {
	old := *s
	if len(old) == 0 {
		return 0, false
	}
	v := old[len(old)-1]
	*s = old[:len(old)-1]
	return v, true
}

func (s stack) peek() (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	return s[len(s)-1], true
}
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. Write push, pop, and peek as three one-line slice expressions from
   memory. Which one is off-by-one if you reach for `len(s)` instead
   of `len(s)-1`?
2. `peek` takes a value receiver, `pop` a pointer. Justify each — and
   say what silently happens if `pop` uses a value receiver.
3. If the elements were pointers or large structs, `s[:len(s)-1]`
   leaks the popped element. Add the one line that fixes it, and say
   why the `[]rune` version above doesn't need it.

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
