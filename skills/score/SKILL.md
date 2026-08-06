---
name: score
description: Procedure for /score in go-dojo, and for judging whether the human's attempt qualifies to unlock answers or hints. Use whenever the human runs /score, submits verdicts, asks for a hint, asks whether they're right, or asks for the solution before scoring.
---

# Scoring and hints

## Qualifying attempt (gates CLAUDE.md rule 3)

An attempt qualifies only if it commits to something falsifiable:

- **Bug hunt:** a verdict per variant — CLEAN, or for each claimed
  defect: the specific line(s)/mechanism, why it is wrong, and the
  runtime consequence. "Something's off with the channel" does not
  qualify.
- **Spec/implementation drills:** a written spec covering the
  targeted failure modes, or a genuine code attempt.

If it doesn't qualify, name the gap and ask again. Never negotiate
the bar down.

## Hint ladder — priced, logged, never volunteered

1. Name the affected function/region. Cost −1 (rubric H1).
2. Name the defect category. Cost −2 more (H2).
3. The answer — that defect scores as a miss and forfeits R4 for
   that variant (H3).

Record hints used in the log line.

## Scoring

If sealed/<week>/ is unreadable, ask the human to run
`scripts/unseal.sh <week>` — never work around the seal.

Score strictly against rubric.md, citing a line ID for every point
given or withheld — including N1 normalization and the R4c
clean-variant rule. No grade inflation: between two grades, award
the lower and state exactly what would have earned the higher; a
generous grade is a failed drill. False positives cost per F1 —
flagging clean code is a review failure, not caution. A correct
"approve" on a clean variant scores 10 (R4c), not a consolation.

Only during /score (and the Sunday verifier pass defined in
skills/author-variant/SKILL.md) may sealed/ be opened. After
scoring, reveal the full key: every defect, its reproduction, every
false positive explained.

## Log discipline

Append one line to each relevant table in log.md — the header row IS
the schema — including hints used. Category fields are filled
post-reveal (secrecy only holds pre-reveal). Never edit or delete
existing lines; to fix an error, append a CORRECTION: line
referencing the original entry.
