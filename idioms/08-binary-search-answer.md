# Idiom 08 — binary search on an answer space

## Canonical

```go
package idiom

import "sort"

// minFeasible returns the smallest x in [lo, hi] with ok(x) true,
// given ok is monotone over the range: false...false true...true.
// If nothing in the range is feasible it returns hi — the caller
// checks ok(result).
func minFeasible(lo, hi int, ok func(int) bool) int {
	for lo < hi {
		mid := lo + (hi-lo)/2 // no overflow, unlike (lo+hi)/2
		if ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// Worked shape: least capacity that ships all weights in d days.
func minCapacity(weights []int, days int) int {
	lo, hi := 0, 0
	for _, w := range weights {
		if w > lo {
			lo = w // capacity must fit the largest single item
		}
		hi += w // one day, everything
	}
	return minFeasible(lo, hi, func(cap int) bool {
		used, load := 1, 0
		for _, w := range weights {
			if load+w > cap {
				used++
				load = 0
			}
			load += w
		}
		return used <= days
	})
}

// Standard-library form over an index space. sort.Search finds the
// smallest i in [0, n) with f(i) true, and returns n if none is.
func firstAtLeast(a []int, target int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= target })
}
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. Write the loop invariant `minFeasible` maintains. Which of
   `lo < hi` / `lo <= hi` and `hi = mid` / `hi = mid-1` pair up, and
   which combination spins forever?
2. What exactly does "monotone predicate" buy you, and how do you
   check that a candidate problem has one before committing?
3. In `minCapacity`, defend both bounds. What goes wrong if `lo`
   starts at 0 or 1 instead of the max weight — a wrong answer, or an
   infinite loop?

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
