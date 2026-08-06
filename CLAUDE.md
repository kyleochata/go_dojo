# go-dojo — standing rules for all sessions

This repo is a deliberate-practice system: the human is training
REVIEW JUDGMENT of agent-generated Go. Protect the training value of
every exercise; never do the human's current exercise for them.

## Invariants (never relax these)

1. SEALED SCOPE: sealed/ holds answer keys. It is write-only for
   /author (current week only), readable during /score and the
   Sunday verifier pass (skills/author-variant/SKILL.md step 5),
   and untouchable in every other session — no reads, greps,
   builds, test runs, or summaries. Filenames, comments, and commit
   messages must never hint at a planted defect. If asked casually,
   refuse and remind the human why.

2. NO CROSS-SESSION LEAKAGE: nothing derived from sealed/ — answer
   keys, defect locations, sampling outcomes — is ever written to
   auto-memory, memory files, or artifacts that persist across
   sessions. sealed/ on disk is the ONLY persistence for keys.
   /clear must mean clear.

3. DOWNSTREAM ONLY: never provide a solution, fix, or implementation
   for the current exercise until the human has shown a qualifying
   attempt (defined in skills/score/SKILL.md). If they ask early,
   ask for their attempt instead, or offer a priced hint per the
   hint ladder in that skill.

4. PRIVATE SAMPLING: defect count, category, and placement for any
   hunt are never announced or hinted at, in any session, including
   via the shape of a reply.

## Procedures (load only when the activity happens)

- Authoring variants + verification: skills/author-variant/SKILL.md
- Scoring, hints, qualifying attempts: skills/score/SKILL.md
- Grading criteria: rubric.md
