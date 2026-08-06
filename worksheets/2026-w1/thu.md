# Thursday — Bug Hunt with Controls (30 min) — week 2026-w1

Label: DRILL. No flash deck this week — the foundations deck has
nothing in it until Wednesday's unit produces cards, so the hunt gets
the full thirty minutes.

Hunt target: normally your own implementations from 7–14 days prior.
Week 1 has no aged code, so this week's variants are authored into
`worksheets/2026-w1/hunt/`. From week 3 the target reverts to your
own aged code.

Variant count is 3–5; the number defective is randomly zero to two
and **unknown to you**. False positives are scored as errors
alongside false negatives (rubric F1). A clean variant correctly
approved is 10/10 (rubric R4c) — approving is a verdict, not an
abstention.

## The hunt

Variants live in `worksheets/2026-w1/hunt/`:

- `variant1/` — package `ledger`
- `variant2/` — package `dispatch`
- `variant3/` — package `manifest`

Read the `.go` files. Review each variant fully before writing any
verdict. **No running the code, no `go test`, no `go vet`, no race
detector until every verdict below is committed** — this is a reading
drill first. The `_test.go` files are part of the delivered code and
are fair game to read; a test can be as wrong as the code it covers.

Budget roughly 8 minutes per variant and hold to it. Reading one
variant for twenty minutes and skimming the others is the failure
mode this clock exists to prevent.

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

Write a verdict on every variant you were given. A blank row is not
neutrality; it scores as nothing found.

## After verdicts are committed

1. Now run the code / tests / race detector to check yourself:
   ```
   cd worksheets/2026-w1/hunt
   go build ./... && go vet ./... && go test -race ./...
   ```
2. `scripts/unseal.sh 2026-w1`, then /score against sealed keys and
   rubric.md.
3. Log: score, false positives, false negatives, hint spend, and the
   category of anything missed → spaced re-attempt queue.

---

## 10-minute floor

One variant, picked without choosing — `ls -d variant* | sort -R |
head -1` — full verdict row, no hints. Score
just that row at the next unsealing. Log "floor — 1 variant."
