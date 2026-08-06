---
name: author-variant
description: Procedure for authoring go-dojo hunt variants and worksheets. Use this skill every time the human requests a worksheet, drill, variant, bug hunt, or exercise — any request to generate code for them to review. Never author from memory of the rules; always follow this file.
---

# Authoring a bug-hunt variant

The human trains signal detection, not treasure hunting. Everything
sampled below is private — never announce or hint at an outcome,
including via the shape of your preamble (CLAUDE.md rule 4).

## 1. Roll the dice — with the script, never in your head

Run: `python3 scripts/sample.py <week>` (add `--variants 3` on
flash-deck Thursdays). It draws with real randomness — variant
count, which variants are defective, defects per variant, and the
placement class of each defect — and writes the outcome straight to
`sealed/<week>/sampling.json` so it never appears in a transcript.
Read it from there and follow it exactly. You are a biased die; the
script is not. Never restate its contents in any visible reply.

## 2. Sample the category (judgment, weighted)

Prefer real-bugs.md over catalogue.md as it grows. Roughly 50%
current unit, 50% all prior units; resample early toward any
category the human has scored poorly on. Log context: read only the
last ~20 rows of the Sessions table plus the rung-state table —
never the whole log.

## 3. Write it cold

Vary naming, package layout, and style from previous weeks. The
governing principle: a defective line must be indistinguishable in
every respect from a clean one — its style, salience, and
surroundings carry zero information. Filenames stay neutral
(worker.go, never nil_map_variant.go). Where each defect goes
(low-salience vs. point of interest) comes from sampling.json, not
from taste.

## 4. Verify (same session, mandatory)

Under the Go version pinned at the top of catalogue.md:

1. `go build ./...` — must compile.
2. `go vet ./...` — the planted defect must NOT be caught by vet;
   if vet catches it, it is too shallow for this drill — resample.
3. Each intended defect demonstrably manifests: a failing test, a
   race detector hit, or observably divergent behavior. Reproduce
   it, and record the reproduction in the key.
4. No unintended defects exist beyond the sampled ones.

Write the full answer key to sealed/<week>/ BEFORE presenting
anything. Never serve an unverified variant.

## 5. Fresh-context verifier pass (mandatory — do not skip)

The author cannot grade its own disguise. Spawn a subagent — or have
the human /clear into a fresh session — carrying ONLY this checklist
and none of the authoring context:

- every variant compiles on the pinned Go version
- every key in sealed/<week>/ is complete: defect, line, reproduction
- every planted defect reproduces exactly as its key claims
- nothing near a defect reads as a tell (comments, phrasing, idiom
  shift, suspicious naming)
- clean variants contain no unintended defects

The verifier may read sealed/<week>/ for this purpose only
(CLAUDE.md rule 1). Fix anything it flags, re-verify, then stop.

## 6. Present bare, then hand the seal to the human

Code only. No difficulty rating, no category hint, no phrasing that
implies a count. "Review these" is the whole prompt. Close by
reminding the human: commit first, then `scripts/seal.sh <week>`.
