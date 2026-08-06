# Idiom 07 — `strings.Builder` accumulation

## Canonical

```go
package idiom

import (
	"strconv"
	"strings"
)

// Separator-aware accumulation. Builder's zero value is ready to use
// — no make, no initialisation.
func joinInts(vals []int, sep string) string {
	var b strings.Builder
	for i, v := range vals {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

// Pre-sized: one allocation when you can estimate the output.
func repeatRune(r rune, n int) string {
	var b strings.Builder
	b.Grow(n * len(string(r)))
	for range n {
		b.WriteRune(r)
	}
	return b.String()
}
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. Name the four write methods and their argument types. Which one do
   you reach for with a `byte`, and which with a `rune`?
2. A `strings.Builder` must not be copied after first use — what
   enforces that, and what is the runtime error message you'd see?
3. When is `strings.Join` or `+=` in a loop actually the right call
   instead of a Builder? Give the crossover reasoning, not a rule.

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
