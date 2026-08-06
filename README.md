# go-dojo

Deliberate-practice repo for training review judgment of
agent-generated Go. Companion documents (keep OUTSIDE session
context; @-mention only when needed):

- `go-native-plan-final.md` — the frozen spec (drop a copy in the
  repo root for /author to reference)
- the operating manual — human-facing playbooks only; it quotes no
  repo file and defers to them everywhere

## Day 0 bootstrap

```
git init
go version                      # record it at the top of catalogue.md
# then, in Claude Code:
#   /author Build the standing template week into templates/ per
#           the ruling: mon.md (laddered build), tue.md (idiom deck
#           + pattern), wed.md (foundations w/ pre-test), thu.md
#           (hunt sheet, verdict table for up to 5 variants),
#           fri.md (four-part explain-before-accept frame). Each
#           sheet ends with its written 10-minute floor version.
#   /author Create idioms/ with one file per snippet from the idiom
#           deck list in go-native-plan-final.md.
#   > Read catalogue.md. For each entry write a minimal program
#     demonstrating the defect, run it on this toolchain, flag any
#     that no longer reproduce.
git add -A && git commit -m "day 0: operational"
```

Also write the work-diff threshold note (80 lines / any concurrency
or error-handling touch) and put it where you review diffs.

## The seal workflow (mechanical, not just prose)

- **Sunday:** /author runs `scripts/sample.py <week>` (real
  randomness, outcome written straight into sealed/ — never the
  transcript), authors the week, verifies, then a fresh-context
  verifier pass checks the output. **Commit, then**
  `scripts/seal.sh <week>` — git cannot read sealed files, so the
  order matters.
- **Mon–Sat:** sealed/<week> is mode 000. A casual `cat` fails; so
  does a stray grep. CLAUDE.md rule 1 is the second wall, not the
  only one.
- **Scoring:** `scripts/unseal.sh <week>`, then /score. Unsealing is
  a deliberate act by design.

Flash-deck Thursdays: pass `--variants 3` to sample.py.

## Changes vs. the original files (context audit, Jul 2026)

- A1/A12 — rubric.md: clean-variant approve is now explicitly 10/10
  (R4c); multi-defect variants normalize to base 10 (N1); H3
  forfeits R4; half-located claims are F2.
- A2/A6 — the fresh-context verifier pass moved from the manual into
  skills/author-variant/SKILL.md, where sessions actually load it.
  Repo files are the single source of truth.
- A3 — sampling is done by scripts/sample.py; sealing by
  seal.sh/unseal.sh.
- A4 — CLAUDE.md rule 1 rewritten as a read/write scope so /author
  can do its Sunday job without contradiction.
- A5 — the constitution's framing is mode-neutral; the adversarial
  posture lives in author files, clarity lives in tutor.md.
- A7 — grade-inflation, version-pin, and log-discipline rules moved
  into the skills that load when they matter.
- A9 — /author reads only the last ~20 log rows + rung state, not
  the whole log.

Housekeeping: disable or scope auto-memory for this repo in settings
if your surface allows (A8), and periodically ask a fresh session
what it knows about this week's hunt — the correct answer is
nothing. Run `/doctor` after any gate-review edit to these files.

At gate reviews, audit realized sampling frequencies (post-reveal
log data) against the targets documented in scripts/sample.py.
