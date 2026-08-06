# SEALED — 2026-w1 Tuesday key (idiom deck + two pointers)

Open after the attempt, not before (CLAUDE.md rule 3).

## Part 1 — idiom deck marking

The three drawn cards check against `idioms/02-set.md`,
`idioms/03-sort-slice.md`, `idioms/07-strings-builder.md`. Mark
against those files, not against this one.

Counts as a **stumble**, not a nit:

- `map[T]struct{}` written without `make`, or written as
  `map[T]bool` (works, but the card is the `struct{}` form and the
  zero-width value is the point).
- Membership tested as `if seen[k]` where the two-value comma-ok
  form was needed to distinguish "absent" from "present and zero".
- A `sort.Slice` comparator using `<=` — not a strict weak ordering;
  `sort` may panic or misbehave. This is the one that matters.
- A comparator that reads the slice by a captured copy rather than
  by index.
- `sb.String()` called inside the accumulation loop (O(n²), and the
  Wednesday unit will come back to it).
- Taking the address of a `strings.Builder` after it has been used,
  or copying one into another variable.

## Part 2 — recognition cues (reference for §1 of the pattern slot)

Cues that should have been written:

1. **The input is sorted** (or cheaply sortable and you were going to
   sort anyway).
2. **The question is about a pair or a bounded window whose answer
   moves monotonically** — pushing one pointer in only ever makes the
   sum smaller, only ever makes the window shorter, and so on. That
   monotonicity is what licenses never backtracking.

Near-miss: **the same question on unsorted input.** "Find two numbers
summing to a target" on unsorted input is a hash set in O(n); sorting
first to enable two pointers costs O(n log n) and is a worse answer
that feels cleverer. If the human wrote "unsorted two-sum" as their
near-miss, that is a full mark on that line.

Also acceptable as a near-miss: a problem needing *all* triples
enumerated rather than counted — two pointers counts in O(n²) but
cannot enumerate O(n³) results faster than O(n³).

## Part 3 — TripletCount reference

```go
// TripletCount counts index triples i < j < k in the ascending
// slice nums with nums[i]+nums[j]+nums[k] <= target.
// O(n^2) time, O(1) extra space.
func TripletCount(nums []int, target int) int {
	count := 0
	for i := 0; i+2 < len(nums); i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			if nums[i]+nums[lo]+nums[hi] <= target {
				// nums is ascending, so every index in
				// (lo, hi] also satisfies the bound with
				// this lo. That is hi-lo triples at once.
				count += hi - lo
				lo++
			} else {
				hi--
			}
		}
	}
	return count
}
```

Verified against an O(n³) brute force on all five published cases
(go1.26.5): `{-2,0,1,3}/2 → 3`, `{0,0,0}/0 → 1`, `{1,2,3,4}/6 → 1`,
`{}/0 → 0`, `{5}/100 → 0`.

**The scored insight is `count += hi - lo`.** An attempt that writes
`count++` there is O(n²) and wrong; an attempt that fixes it by
looping over the range is correct and O(n³), which fails the stated
complexity requirement. Both failures are worth logging by name —
they are different mistakes.

Second-order points, only if the human's attempt is otherwise clean:

- `i+2 < len(nums)` rather than `i < len(nums)` — same result either
  way (the inner loop is empty), but the intent is clearer.
- Empty and single-element slices need no special case; the loop
  bounds handle them. An attempt with a defensive `if len(nums) < 3`
  guard is fine, not better.
