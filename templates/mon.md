# Monday — Laddered Build (30 min)

Label: LEARN at rungs 0–2, DRILL at rungs 3–4.

## Primitive

**Primitive:** ______________________ (fill from the rung-state table
in log.md; candidates: LRU cache, worker pool, rate limiter, trie,
heap, graph traversal)

**Entry rung:** ____ — enter at the lowest rung you can't do cold.
Check log.md rung state before choosing; don't re-run a cleared rung.

**Rung 0 entry test:** run rung 0 if you can't state, without looking
anything up, what problem this primitive solves and why one obvious
data structure isn't enough for it. If you can state that, skip to
rung 1 — rung 0 is not a formality to log for primitives you've
already met.

## The ladder

| Rung | Format | Your role |
|---|---|---|
| 0 | Concept | No code, no language. Name the problem, name why the obvious single structure fails, derive the real structure from that failure. Tutor mode may lead. |
| 1 | Worked example | Read a canonical implementation (tutor mode may provide). Annotate every design decision in your own words. |
| 2 | Implement with reference | Build it with the worked example open. No copy-paste; type everything. |
| 3 | From memory | Reference closed. Build it cold inside the clock. |
| 4 | Race the agent | Same spec, you and generator mode in parallel. Compare after; write one sentence on each divergence. |

**Rung 0 clears** when you can explain the design to someone else
without naming a language, a type, or a variable — in the problem's
own terms. Needing a variable name to explain it means you memorized
the code, not the idea. It is per-primitive knowledge, so the
two-runs rule below does not apply to it: one clean explain-back on
this primitive clears it for this primitive only.

**Advancement (rungs 1–4):** a rung clears on two clean unaided runs,
on different primitives, inside the clock. One clean run advances
nothing — log it and queue the second.

## Session

1. (2 min) Write the spec you're building to: operations, expected
   complexity, one edge case you'll test.

   - Operations: ______________________
   - Complexity target: ______________________
   - Edge case: ______________________

   *At rung 0, skip this step.* You can't spec a thing you can't yet
   describe, and guessing at complexity targets before you understand
   the structure teaches you to bluff. Writing this spec unaided is
   the first thing rung 1 asks of you.

2. (~25 min) Run the rung. Timer visible. If rung 3–4, no references,
   no agent help before your attempt (CLAUDE.md rule 3).

3. (3 min) Verify: `go build`, run your edge-case test. Then log:
   rung, clean/failed, what broke, next entry point.

   *At rung 0 there is nothing to build.* Verify by explaining the
   design back with no code open and no code words — if you reach for
   a type or variable name, you haven't cleared it. Log what you
   couldn't explain without the code; that's the data.

**Clock rule:** if the clock expires mid-build, stop, log the failure
point — that's the data. A LEARN unit may spend an overtime token
(max two/week, non-rolling); a DRILL unit may not.

---

## 10-minute floor

One primitive, current rung, one operation only (e.g. LRU: `Get` +
eviction path only, no `Put` resizing; worker pool: submit + drain
only). Spec is one line. Build, compile, log the rung attempt as
"floor". No token spend on a floor day.

At rung 0 the floor is: one paragraph, in your own words, on why the
obvious single structure fails for this primitive. No code, no
explain-back of the full design. Log it as "floor".
