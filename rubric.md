# rubric.md — grading criteria for go-dojo

Every point awarded or withheld in /score must cite a line ID from
this file (e.g. R3, F1, N1). Edit weights at gate reviews only; keep
IDs stable so old log entries stay interpretable.

## Bug-hunt verdicts (per variant)

- R1 (+3 per defect): defect located — correct line(s) or mechanism.
  Right function but wrong line AND wrong mechanism is not R1;
  score it F2 instead.
- R2 (+2 per defect): defect explained — why it is wrong, not just
  where.
- R3 (+2 per defect): consequence stated — what breaks, under what
  conditions (load, shutdown, error path, race timing).
- R4 (+3): correct overall verdict on a DEFECTIVE variant — all
  defects found with no false positives.
- R4c: on a CLEAN variant, a correct "approve" IS the whole score:
  10/10. Full marks, not a consolation — this is the hardest verdict
  in the drill.
- N1 (normalization): a defective variant's raw score is scaled to
  base 10, rounded DOWN: 10 × raw ÷ max-raw, where max-raw = 10 for
  one planted defect and 17 for two. Rounding down implements
  "award the lower." Floor: a variant never scores below 0.

## Penalties

- F1 (−2 each): false positive — a claimed defect that is not real.
- F2 (−1 each): vague or half-located claim — a real defect flagged
  without qualifying specificity (see skills/score/SKILL.md);
  earns no R1–R3 credit.
- H1 (−1): hint 1 taken. H2 (−2 more): hint 2 taken.
- H3: the reveal — that defect scores as a miss: no R1–R3 for it,
  and R4 is forfeited on that variant.

## Spec / implementation drills (base 10)

- S1 (+3): all targeted failure modes covered.
- S2 (+3): behavior under each failure mode specified, not just
  named.
- S3 (+2): boundary edge cases (zero values, closed channels,
  canceled contexts) addressed.
- S4 (+2): the spec/code would actually catch or prevent the
  targeted defect.

## Grade bands

- 9–10: reviewer you'd trust unsupervised.
- 7–8: solid; cite what was missed.
- 5–6: partial credit; the gap is the next drill.
- <5: rerun a variant from the same category within two sessions.

Between two bands, award the lower and state exactly what would have
earned the higher (no-grade-inflation rule, skills/score/SKILL.md).
