# SEALED — 2026-w1 Wednesday pre-test key (Big-O, part 1)

Open only to mark a completed cold attempt, and again to mark the
post-test. Not before the cold attempt (CLAUDE.md rule 3).

Subject: `worksheets/2026-w1/wed-pretest/analyze.go`. Every function
in it is correct; there is nothing to find.

## Per-function answers

| # | func | time | extra space | what dominates |
|---|---|---|---|---|
| A | `hasDuplicate` | O(n²) worst case, O(1) best (duplicate at the front) | O(1) | the inner loop; ~n²/2 comparisons when all elements are distinct |
| B | `hasDuplicateSeen` | O(n) expected | O(n) | n hash insert/lookups; "expected" because map cost is average-case, not worst-case |
| C | `report` | O(n²·L) — quadratic in n | O(n²·L) allocated in total, O(n·L) live at the end | `out += n + "\n"` allocates a **new** string each iteration and copies everything accumulated so far: L + 2L + … + nL |
| D | `reportBuilder` | O(n·L) amortized | O(n·L) | `strings.Builder` appends into a growable buffer and only materializes a string once, in `String()` |
| E | `topK` | O(n log n) | O(n) | the sort; the `copy` is O(n) and does not change the class |
| F | `contains` | O(log n) | O(1) | the halving loop |
| G | `anyGroupHasDuplicate` | O(g·n²) worst case | O(1) | calls A once per group; early return makes the best case O(n²) or better |
| H | `evens` | O(n) amortized | O(n) | one pass; the appends are amortized O(1) each — see Q3 |

Marking notes:

- "O(n²)" for C is the answer being looked for. Writing O(n) for C is
  the single most common miss and by itself fails the pre-test.
- For B, an answer of "O(n)" is accepted; "O(n) expected / average"
  is the better answer and worth saying so out loud.
- For E, `k` is clamped both ways by `max(0, min(k, len(cp)))`, so
  there is no bounds trap to find. An answer that misses the O(n)
  space is a half-mark. Extra
  credit, not required: `cp[:k]` keeps the **whole** n-element
  backing array alive, so a `topK(hugeSlice, 3)` result retains n
  ints, not 3. That is a real memory-retention bug pattern, and it
  is why the answer is O(n) space rather than O(k).
- For G, "O(g·n²)" and "O(n²) per group, g groups" are the same
  answer. An answer of "O(n²)" that never mentions g is a miss.

## Q1 — the two pairs

**Pair 1: A and B** — both answer "does this slice contain a
duplicate?"

Where B starts to win is a constant-factor question, not an
asymptotic one: A's inner step is an integer compare (about a
nanosecond); B's step is a hash, a bucket probe, and possibly an
allocation (tens of nanoseconds). Measured on the pinned toolchain
(go1.26.5, all-distinct input, the worst case for A). **Absolute
numbers are machine-dependent — a re-run on other hardware moved them
by up to 40%. The crossover position, not the nanoseconds, is the
claim:**

| n | A (nested) | B (set) |
|---|---|---|
| 16 | 134 ns | 1674 ns |
| 32 | 405 ns | 3689 ns |
| 64 | 1833 ns | 7152 ns |
| 128 | 6441 ns | 12130 ns |
| 256 | 31977 ns | 20804 ns |

Crossover is between n = 128 and n = 256 — call it **n ≈ 200**, and
A is more than 10× faster at n = 16. Any answer in the range "around
a hundred to a few hundred, and you'd measure it" is full marks. An
answer of "B is always faster" is the miss this question exists to
catch: asymptotically better is not the same as faster, and for small
n the O(n²) code is the right code.

**Pair 2: C and D** — both build the same string.

Here there is effectively **no crossover**: D wins from about n = 2
upward and the gap grows without bound. This is the contrast worth
holding onto — pair 1's answer is "it depends, measure it"; pair 2's
answer is "just use the Builder". Knowing which kind of question you
are in is the actual skill.

## Q2 — why C is not linear

The loop body runs n times, so it *looks* linear. But Go strings are
immutable: `out += s` cannot extend `out` in place. It allocates a
new string of length `len(out)+len(s)` and copies both operands into
it. On iteration i the copy costs about i·L bytes, so the total is
L·(1 + 2 + … + n) = L·n(n+1)/2 = **O(n²·L)**.

The rule to take away: *a loop is O(n) only if its body is O(1). An
O(1)-looking operation on a value that grows with the loop is not
O(1).* The same trap appears as `s = append(s, xs...)` inside a loop
over slices, and as `m = mergeMaps(m, next)` in a fold.

## Q3 — why repeated `append` is not O(n²)

`append` does reallocate — but not every time, and not by a constant
amount. When capacity runs out the runtime allocates a **larger**
backing array (roughly doubling for small slices, tapering to ~1.25×
growth for large ones) and copies the existing elements once.

Count the copying across n appends: the reallocations happen at
capacities ~1, 2, 4, 8, …, n, and the copy work is
1 + 2 + 4 + … + n < 2n. That geometric series is the whole argument —
total copy work is **linear**, so the cost per append is O(1)
*amortized*.

What "amortized" licenses you to say: n appends cost O(n) **in
total**. It does not license "every append is fast" — any individual
append can be O(n) when it triggers the copy. If the code has a
latency budget per operation rather than per batch, amortized O(1)
is not good enough, and `make([]T, 0, n)` up front is the fix.

The contrast with C is the point of putting them side by side:
`append` grows geometrically and is amortized linear; `+=` on strings
grows by exactly what you added and is quadratic.

## Pass bar

All eight rows correct **and** all three questions correct. Seven of
eight is a fail and the missed row is the unit's entry point. The
floor version (A–F only, no questions) passes on six correct rows.

## If the unit runs

Anchor it to the specific miss, per the sheet. The Monday tie-in is
live this week: an LRU whose `Put` scans a slice for the
least-recently-used key is correct, passes every test, and is O(n).
Complexity analysis is the only review tool that rejects it.
