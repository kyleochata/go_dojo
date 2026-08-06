# Thursday — Bug Hunt with Controls (30 min)

Label: DRILL. On flash-deck days the hunt is honestly 25 minutes and
runs three variants, not five — the deck's five minutes come first.

Hunt target: your own implementations from 7–14 days prior, never
the current week. Variant count is 3–5; the number defective is
randomly zero to two and **unknown to you**. False positives are
scored as errors alongside false negatives (rubric F1). A clean
variant correctly approved is 10/10 (rubric R4c) — approving is a
verdict, not an abstention.

## Flash deck (5 min, deck days only)

Cards drawn: ______________________ Results: ______________________

## The hunt

Week's variants live in worksheets/<week>/. Review each fully before
writing any verdict. No running the code until every verdict is
committed in the table below — this is a reading drill first.

Hints cost points (H1 −1, H2 −2, H3 = scored as a miss). Take them
deliberately or not at all.

## Verdict table (fill every row you were given; leave extras blank)

| Variant | Verdict (approve / reject) | Defect location(s) — file:line or mechanism | Why it's wrong | Consequence — what breaks, when | Confidence (L/M/H) |
|---|---|---|---|---|---|
| 1 | | | | | |
| 2 | | | | | |
| 3 | | | | | |
| 4 | | | | | |
| 5 | | | | | |

A "reject" with a vague location earns F2, not R1 — be specific
enough that someone could fix the bug from your row alone.

## After verdicts are committed

1. Now run the code / tests / race detector to check yourself.
2. `scripts/unseal.sh <week>`, then /score against sealed keys and
   rubric.md.
3. Log: score, false positives, false negatives, hint spend, and the
   category of anything missed → spaced re-attempt queue.

---

## 10-minute floor

One variant, full verdict row, no hints, no flash deck. Score just
that row at the next unsealing. Log "floor — 1 variant."
