# Monday — Laddered Build (30 min) — week 2026-w1

Label: LEARN at rungs 1–2, DRILL at rungs 3–4.

## Primitive

**Primitive:** LRU cache (`Get`, `Put`, fixed capacity, eviction of
the least-recently-used entry)

**Entry rung:** ____ — enter at the lowest rung you can't do cold.
The rung-state table in log.md is empty, so nothing is cleared and
nothing is barred: if you can't type an LRU cold right now, that is
rung 1 or 2, and starting there is the correct move, not a concession.

## The ladder

| Rung | Format | Your role |
|---|---|---|
| 1 | Worked example | Read a canonical implementation (tutor mode may provide). Annotate every design decision in your own words. |
| 2 | Implement with reference | Build it with the worked example open. No copy-paste; type everything. |
| 3 | From memory | Reference closed. Build it cold inside the clock. |
| 4 | Race the agent | Same spec, you and generator mode in parallel. Compare after; write one sentence on each divergence. |

**Advancement:** a rung clears on two clean unaided runs, on
different primitives, inside the clock. One clean run advances
nothing — log it and queue the second.

## Session

1. (2 min) Write the spec you're building to: operations, expected
   complexity, one edge case you'll test.

   - Operations: ______________________
   - Complexity target: ______________________
   - Edge case: ______________________

   Acceptance criteria for this primitive (author-supplied — your
   spec above should imply all of them, and Wednesday's Big-O unit
   grades the second one):

   - `New(capacity int)` rejects or documents capacity <= 0.
   - `Get` and `Put` are both O(1) — if either is O(n), the build
     fails the drill no matter how correct it is.
   - `Get` on a hit promotes the key to most-recently-used.
   - `Put` on an existing key updates the value AND promotes it; it
     does not grow the cache.
   - At capacity, `Put` evicts exactly one entry, the least-recently
     used, and the evicted key is gone from every internal structure
     (the classic failure is evicting from the list and leaving the
     map entry behind — a slow leak, not a wrong answer).

   Edge cases worth choosing from: capacity 1; `Put` of a key that is
   already the LRU; `Get` of a missing key; `Put`ting the same key
   twice in a row.

2. (~25 min) Run the rung. Timer visible. If rung 3–4, no references,
   no agent help before your attempt (CLAUDE.md rule 3).

3. (3 min) Verify: `go build`, run your edge-case test. Then log:
   rung, clean/failed, what broke, next entry point.

**Clock rule:** if the clock expires mid-build, stop, log the failure
point — that's the data. A LEARN unit may spend an overtime token
(max two/week, non-rolling); a DRILL unit may not.

---

## 10-minute floor

`Get` and its promotion path only — no eviction, no `Put` resizing.
Fixed-capacity map plus whatever list you choose; one line of spec.
Build, compile, log the rung attempt as "floor". No token spend on a
floor day.
