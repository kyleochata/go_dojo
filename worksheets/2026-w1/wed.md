# Wednesday — Foundations Unit (30 min) — week 2026-w1

Label: LEARN. Wednesday is NOT a drop day this week (it becomes one
from week 6; in weeks 3–5 the drop day is Tuesday).

**Unit:** Big-O micro-unit, part 1 of 2 — the cost of code you
actually write, not asymptotics for their own sake.

## Part 1 — Pre-test (10 min, cold, gates the unit)

Plan rule 4: every foundations unit is preceded by a cold attempt at
its own review drill. Attempt it before any teaching, before any
reading about complexity, and before running anything.

- Pre-test drill: **`worksheets/2026-w1/wed-pretest/analyze.go`.**
  Eight functions, labelled A–H. For each one write, on paper or in
  a scratch file:
  1. worst-case **time** complexity in terms of the stated n,
  2. **extra space** complexity (what it allocates beyond its input),
  3. one clause saying *what makes it that* — the line or operation
     that dominates.

  Then answer three questions:
  - Two **pairs** of functions in this file compute the same thing
    at different cost. Name both pairs, and for each say roughly
    where the cheaper one starts to win.
  - Function C looks linear and is not. Say why.
  - Function H: why is repeated `append` **not** O(n²)?

  Every function in that file is correct Go. This is a cost drill,
  not a bug hunt — do not spend the ten minutes hunting for defects.

- Attempt result (cold): pass / fail

  **Pass = all eight complexities right AND all three questions
  right.** Seven of eight is a fail; the missed one is the entry
  point. Marking yourself generously here cuts a unit you needed.

- **Pass cold → the unit is CUT.** Log it, spend the freed 20 min on
  the spaced re-attempt queue or stop early. Do not run the unit
  anyway "for completeness."
- **Fail → that's the entry point.** Note exactly what you got
  wrong; the unit targets that, and the same drill is the post-test.

What I got wrong cold: ______________________

## Part 2 — The unit (15 min, only on a failed pre-test)

Tutor mode is unrestricted here: explanations, canonical examples,
quizzing. Anchor everything to the pre-test miss — the unit exists
to close that specific gap, not to survey the topic.

Ask tutor mode for exactly the missed item. Sensible anchors,
depending on what you missed:

- **String concatenation in a loop** — why `out += s` copies, and
  what `strings.Builder` does instead.
- **Amortized cost** — why Go's `append` growth policy makes n
  appends O(n) total, and what "amortized" licenses you to say.
- **Hash vs. scan** — the constant-factor honesty: at what n does
  the map actually beat the nested loop, and why the answer isn't
  "always".
- **Cost of the sort you didn't write** — `sort.Slice`,
  `sort.Sort(sort.Reverse(...))`, and the O(n log n) you inherit
  from calling them.

Then:

- Core mechanism, in my own words: ______________________
- One canonical example, typed out and run: ______________________
- How this shows up in agent-generated code: ______________________

For that last line, use Monday's LRU: an agent implementation that
scans a slice to find the least-recently-used key is *correct* and
O(n) per `Put`. It passes every test you'd write. Complexity is the
only tool that rejects it — that is the whole reason this unit is
front-loaded.

## Part 3 — Post-test (5 min)

Re-run the same eight functions and three questions. Pass → unit
closed; log it. Fail → the unit is mis-scoped or unfinished: split
it, log the split, queue the remainder. A LEARN unit may burn one
overtime token to close; never a second in the same week.

---

## 10-minute floor

Pre-test only: functions A–F, skip G and H, skip the three questions.
Attempt cold, log pass (unit cut — a full result, not a partial one)
or fail (entry point recorded for the next full Wednesday). No
teaching on a floor day.
