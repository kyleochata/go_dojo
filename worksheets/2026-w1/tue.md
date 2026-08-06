# Tuesday — Idiom Deck + Pattern (30 min) — week 2026-w1

Label: LEARN when the pattern is new, DRILL on review.
Tuesday is NOT a drop day this week (the drop day is Tuesday only in
weeks 3–5, and Wednesday from week 6).

## Part 1 — Idiom deck (10 min)

Cold, from memory, compiling — an idiom you can't type without docs
isn't automatic yet. Every card is new this week; the deck starts
with three and grows. A card leaves the daily deck for spaced
maintenance once it's been automatic twice.

Deck (one file per snippet lives in idioms/):

- [ ] `container/heap` with a concrete type
- [x] set via `map[T]struct{}`
- [x] `sort.Slice` with a custom comparator
- [ ] BFS queue on a slice
- [ ] two pointers
- [ ] stack
- [x] `strings.Builder` accumulation
- [ ] binary search on an answer space

Today's draw: **set via `map[T]struct{}`**, **`sort.Slice` with a
custom comparator**, **`strings.Builder` accumulation**

Type each one into a scratch file cold — no idioms/ file open, no
autocomplete acceptance — then compile. Only after all three are
written do you open idioms/02, idioms/03, idioms/07 to check
yourself.

Bar for "clean": it compiles first try and matches the reference in
substance. Details that count as stumbles, not nits: forgetting
`make` on the map, a comparator that isn't a strict weak ordering
(`<=` instead of `<`), calling `sb.String()` inside the loop.

Result (clean / stumbled / failed): ______, ______, ______

## Part 2 — One pattern slot (20 min)

**Pattern:** two pointers — opposite-ends variant

1. (5 min) **Recognition cues, written down.** What in a problem
   statement tells you this pattern applies? List at least two cues
   and one near-miss (a problem that looks like this pattern but
   isn't).

   - Cue 1: ______________________
   - Cue 2: ______________________
   - Near-miss: ______________________

   Commit your three lines before you start coding. The reference
   cues are sealed with this week's keys; compare after.

2. (13 min) **One problem, in Go.** State complexity before running.

   > `func TripletCount(nums []int, target int) int` — `nums` is
   > sorted ascending and may contain duplicates. Return the number
   > of index triples `i < j < k` with
   > `nums[i]+nums[j]+nums[k] <= target`.
   >
   > Required: O(n²) time, O(1) extra space. Fix `i`, then run two
   > pointers over the remainder.
   >
   > Cases to check yourself against:
   >
   > | nums | target | want |
   > |---|---|---|
   > | `{-2, 0, 1, 3}` | 2 | 3 |
   > | `{0, 0, 0}` | 0 | 1 |
   > | `{1, 2, 3, 4}` | 6 | 1 |
   > | `{}` | 0 | 0 |
   > | `{5}` | 100 | 0 |
   >
   > If you pass all five, you are done — do not go looking for the
   > reference. If you fail one, the failing case is the finding;
   > write down what you assumed before you fix it.

3. (2 min) Log: pattern, problem, solved/stuck, and whether the
   recognition cues actually fired before you started coding.

---

## 10-minute floor

Idiom deck only: the three cards above, cold, compiling. Skip the
pattern slot entirely. Log "floor — deck only."
