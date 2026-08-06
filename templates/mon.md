# Monday — Laddered Build (30 min)

Label: LEARN at rungs 1–2, DRILL at rungs 3–4.

## Primitive

**Primitive:** ______________________ (fill from the rung-state table
in log.md; candidates: LRU cache, worker pool, rate limiter, trie,
heap, graph traversal)

**Entry rung:** ____ — enter at the lowest rung you can't do cold.
Check log.md rung state before choosing; don't re-run a cleared rung.

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

2. (~25 min) Run the rung. Timer visible. If rung 3–4, no references,
   no agent help before your attempt (CLAUDE.md rule 3).

3. (3 min) Verify: `go build`, run your edge-case test. Then log:
   rung, clean/failed, what broke, next entry point.

**Clock rule:** if the clock expires mid-build, stop, log the failure
point — that's the data. A LEARN unit may spend an overtime token
(max two/week, non-rolling); a DRILL unit may not.

---

## 10-minute floor

One primitive, current rung, one operation only (e.g. LRU: `Get` +
eviction path only, no `Put` resizing; worker pool: submit + drain
only). Spec is one line. Build, compile, log the rung attempt as
"floor". No token spend on a floor day.
