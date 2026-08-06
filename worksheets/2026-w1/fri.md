# Friday — Explain-Before-Accept (30 min) — week 2026-w1

Label: DRILL. (From week 6, alternate Fridays cede their final ten
minutes to the retro formats. Not this week — the full thirty
minutes are the worksheet.)

Target diff: **`worksheets/2026-w1/fri-diff/change.diff`**, applied
to `worksheets/2026-w1/fri-diff/sessions/store.go`. Agent-generated,
from the prompt *"add TTL expiry and disk persistence to the session
store."* See the README in that directory.

Standing threshold (applies at work too, posted where you review):
**any diff over 80 lines, or any diff at any size touching
goroutines, channels, mutexes, `context` propagation, or error
handling, gets all four parts. Below that, section 4 alone.**

This diff touches a goroutine, a mutex, and error handling, so it
takes all four parts on the concurrency clause alone — the line count
never enters into it. Say in one line which clause fired; recognising
*why* a diff is in the four-part bucket is half the habit.

Produce all four sections IN WRITING, then and only then run or
accept the diff. Do not apply the patch first.

## 1. What the change does — one sentence

______________________________________________________________

If you can't do it in one sentence, that is a finding: the diff does
more than one thing. (Count the things. Write the number down.)

## 2. Every error path (mandatory; separately scored in the log)

For EACH error path in the diff — every place an error is created,
returned, wrapped, checked, or discarded:

| # | What produces the error | What the code does with it | What the caller can now distinguish |
|---|---|---|---|
| 1 | | | |
| 2 | | | |
| 3 | | | |
| 4 | | | |
| 5 | | | |
| 6 | | | |

Column four is the scored one. For each row, name a specific decision
a caller might want to make — retry, create the file, log and carry
on, give up — and say whether this error lets them make it.

A diff with no error paths is itself a finding worth writing down —
say so explicitly rather than leaving the table blank.

## 3. Every copy or share

Every place a value is copied or shared (slices sharing backing
arrays, maps passed by reference, structs copied at call sites or in
`range`, pointers escaping), and whether that was intended. Pay
attention to what crosses the mutex boundary: a pointer that leaves a
locked region takes the lock's protection with it, and doesn't give
it back.

| Place | Copy or share? | Intended? |
|---|---|---|
| | | |
| | | |
| | | |
| | | |

## 4. The one line most likely to be wrong, and why

Line: ______________________
Why: ______________________________________________________________

## Then

Apply and run:

```
cd worksheets/2026-w1/fri-diff
git apply change.diff
go build ./... && go vet ./... && go test -race ./...
```

Two things that command will NOT tell you, and you should predict
both before running it:

- `go vet` finds nothing here.
- `go test -race` prints `ok ... [no test files]`. There are no
  tests. A green line from a suite that does not exist is the most
  comfortable false signal in Go, and you just watched yourself
  accept it.

If you suspect a race, the honest move is to write the test that
provokes it — a race the detector never sees is a claim, not a
finding.

Compare what actually happened against sections 1–4; log the
error-path score (rubric S-lines) and any line where reality
disagreed with your writeup — those disagreements seed the real-bug
list and next Thursday's hunt. Then `git checkout --
sessions/store.go` to restore the pre-image.

---

## 10-minute floor

Sections 2 and 4 only, and only over the two new persistence
functions at the bottom of the diff: the error-path table and the
most-likely-wrong line. Sections 1 and 3 are skipped, not rushed.
Log "floor — §2+§4."
