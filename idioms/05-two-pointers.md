# Idiom 05 — two pointers

## Canonical

```go
package idiom

// Converging: a is sorted ascending. Returns i < j with
// a[i]+a[j] == target, or -1, -1.
func twoSumSorted(a []int, target int) (int, int) {
	i, j := 0, len(a)-1
	for i < j {
		switch sum := a[i] + a[j]; {
		case sum == target:
			return i, j
		case sum < target:
			i++
		default:
			j--
		}
	}
	return -1, -1
}

// Same-direction (read/write): compact in place, return the new
// length. The write pointer never outruns the read pointer.
func removeAll(a []int, drop int) int {
	w := 0
	for _, v := range a {
		if v == drop {
			continue
		}
		a[w] = v
		w++
	}
	return w
}

// Sliding window: longest run with at most k distinct values.
func longestAtMostKDistinct(a []int, k int) int {
	count := make(map[int]int)
	best, lo := 0, 0
	for hi, v := range a {
		count[v]++
		for len(count) > k {
			count[a[lo]]--
			if count[a[lo]] == 0 {
				delete(count, a[lo])
			}
			lo++
		}
		if n := hi - lo + 1; n > best {
			best = n
		}
	}
	return best
}
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. `i < j` vs `i <= j` in the converging loop — which problems need
   each, and what does the wrong one do to a "pair with itself" case?
2. In `removeAll`, why is `a[w] = v` safe when `w <= ` the read
   index? Sketch the one aliasing rule that keeps in-place compaction
   correct.
3. The window loop deletes zero-count keys. What breaks if you skip
   the `delete` and only decrement? Name the exact line that goes
   wrong.

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
