# Friday — Explain-Before-Accept (30 min)

Label: DRILL. From week 6, alternate Fridays cede their final ten
minutes to the retro formats (race-the-agent retros alternating with
spec-first trials). The slot never moves; its length alternates.

Target diff: ______________________ (a real diff from this week's
work, or one authored into worksheets/<week>/; agent-generated
preferred — that's the reviewing you're training for)

Standing threshold (applies at work too, posted where you review):
**any diff over 80 lines, or any diff at any size touching
goroutines, channels, mutexes, `context` propagation, or error
handling, gets all four parts. Below that, section 4 alone.**

Produce all four sections IN WRITING, then and only then run or
accept the diff.

## 1. What the change does — one sentence

______________________________________________________________

If you can't do it in one sentence, that is a finding: the diff does
more than one thing.

## 2. Every error path (mandatory; separately scored in the log)

For EACH error path in the diff:

| # | What produces the error | What the code does with it | What the caller can now distinguish |
|---|---|---|---|
| 1 | | | |
| 2 | | | |
| 3 | | | |

A diff with no error paths is itself a finding worth writing down —
say so explicitly rather than leaving the table blank.

## 3. Every copy or share

Every place a value is copied or shared (slices sharing backing
arrays, maps passed by reference, structs copied at call sites or in
`range`, pointers escaping), and whether that was intended:

| Place | Copy or share? | Intended? |
|---|---|---|
| | | |
| | | |

## 4. The one line most likely to be wrong, and why

Line: ______________________
Why: ______________________________________________________________

## Then

Run/accept. Compare what actually happened against sections 1–4;
log the error-path score (rubric S-lines) and any line where reality
disagreed with your writeup — those disagreements seed the real-bug
list and next Thursday's hunt.

---

## 10-minute floor

One small diff (or one function of a big one), sections 2 and 4
only: the error-path table and the most-likely-wrong line. Sections
1 and 3 are skipped, not rushed. Log "floor — §2+§4."
