# The Go-Native Plan — Final, Audited (v4)

**Status of this document.** This is the operating version. It supersedes the v3 Final ruling. An independent audit tested the v3 Final for internal consistency, feasibility, and failure under contact; the panel's design survived, and five patches were accepted. All five are integrated below, marked **[AUDIT]** where they change a rule, so the provenance is visible without reading anything else. Nothing else from the v3 Final ruling has been altered.

**The audit's verdict, restated for the record:** the plan cannot fail on pedagogy — its mechanisms (pre-testing, spaced retrieval, laddered generation, error-focused review, signal-detection scoring) are among the best-supported in the learning literature. It can only fail on operations. These patches close the five operational holes: a drop day that briefly carried a load-bearing dependency, a single point of failure on Sunday, an undefined threshold that would have defaulted to never, a closed self-grading loop with no external signal, and a sabotage bank that would have taught the agent's tells instead of Go's defects.

**The governing finding, unchanged since Round 1:** the subject's formal training is in the language they use least; their production work is in the language they taught themselves, mediated by an agent, filtered by a reviewer whose documented instinct under uncertainty is to accept. Every ranking below follows from that sentence.

**Standing facts:** career switcher, under two years professional experience, no CS degree; bootcamp MERN, a web-dev internship, then self-directed work in Go, Python, TypeScript. The stack is Go; TypeScript appears only on frontend work. 30 minutes/day, five days, with flex for completion of understanding. Strong learning ability, appetite for volume, needs direction.

---

## Repository rules

1. **Three agent modes.** *Tutor* — explanations, canonical worked examples, quizzing; upstream, unrestricted. *Author* — problems, specs, rubrics, acceptance criteria, sabotage variants, flash cards; upstream, permitted, provided every solution is written to a separate artifact opened only at scoring. *Generator* — solutions, implementations, fixes, designs for the current exercise; downstream only, after the subject's own attempt, hypothesis, spec, or written explanation.

2. **LEARN/DRILL is a label, not a budget.** Two **overtime tokens** per week, +30 minutes each, LEARN-only, non-rolling. A unit that burns a token without closing is mis-scoped: split it, log it, no second token that week.

3. **Ladder every generation drill** — worked example → implement with reference → from memory → race the agent. Enter at the lowest rung you can't do cold. A rung clears on **two clean unaided runs, different primitives, inside the clock**.

4. **Pre-test gate.** Every foundations unit is preceded by a cold attempt at its own review drill. Pass cold → the unit is cut. Fail → that's the entry point, and the same drill is the post-test.

5. **Bug hunts carry controls.** Three to five variants; number defective randomly zero to two and unknown; verdict per variant; **false positives scored as errors** alongside false negatives.

6. **Hunt aged code** — the subject's own implementations from 7–14 days prior, never the current week.

7. **One file, one floor, no streaks, one restart rule.** Single log for mistakes, spaced queue, rung state, token spend. Every sheet carries a written ten-minute minimum version. Adherence scored 7-of-10 per fortnight, not consecutively. After a break: resume at the current rotation day, flash deck only for two days, no re-attempts in week one back.

8. **Costed prep, with a floor.** Sunday, 30–45 minutes: author the week, refresh the sabotage bank, triage the log, version-check the gotcha catalogue. Honest weekly cost ≈ **3.5 hours**.
   **[AUDIT — Sunday floor.]** Sunday was the plan's only day with no failure rule, and it is the day the whole week is manufactured on. Two provisions close that:
   - **The ten-minute Sunday:** author Monday's build sheet and Thursday's hunt variants only. Nothing else.
   - **The standing template week:** authored once, before week 1, and kept in the repository — a complete default worksheet for every slot (a Tuesday pattern sheet, a Wednesday unit, a Friday explain-before-accept frame). On any Sunday that produces less than the full prep, the unauthored days run from the template. A missed Sunday now degrades the week; it cannot stall it.

9. **Language.** Go for all worksheets. TypeScript only where the work is frontend, and in Phase 1 only under the standing rule — no worksheet slot. One language for interview work: **Go**, chosen once, not switched.

10. **Gates, not dates.** Phase 2 opens when *all* of: rung three on at least three build primitives; bug-hunt accuracy ≥60% across at least six sessions with false positives counted as errors; Big-O unit closed; at least one completed spaced re-attempt cycle. Backstop: **week 12 regardless** — the backstop exists so the gates cannot hold the plan hostage, and its tension with "gates, not dates" is deliberate.

11. **The drop day.** In a bad week, run four days and drop the designated day. This is not falling behind.
    **[AUDIT — drop day moves during the on-ramp.]** In weeks 3–5, Wednesday carries the concurrency on-ramp while Thursday's hunts begin weighting toward concurrency in week 4 — dropping Wednesday in that window would leave Thursday hunting defects the subject was never taught to recognise. Therefore: **weeks 3–5, the drop day is Tuesday** (patterns are the most deferrable content in that window). **From week 6 onward, the drop day is Wednesday**, as originally ruled, where its content — internals and debugging — is genuinely the most spaceable in the plan.

12. **[AUDIT — the reality check.]** The plan's scores are otherwise self-graded against agent-authored rubrics, a closed loop that cannot detect its own drift. One line in the log opens it: **every real bug that reaches review or production gets a verdict — *would this week's drills have caught it?*** Yes, no, or partially, with one sentence of why. **Three consecutive "no"s trigger a rebuild of the sabotage bank from the real-bug list.** This is the only mechanism in the plan by which reality grades the curriculum, and it costs about two minutes per incident.

---

## The standing rule (all diffs, all languages, at work and in practice)

**Explain-before-accept.** For each diff, produce in writing:

1. What the change does, in one sentence.
2. **Every error path** — for each, what produces the error, what the code does with it, and what the caller can now distinguish. *Mandatory; separately scored in the log. A diff with no error paths is itself a finding worth writing down.*
3. Every place a value is copied or shared, and whether that was intended.
4. The one line most likely to be wrong, and why.

Only then run or accept.

**[AUDIT — the threshold is now a number.]** The v3 ruling said "every diff above some size," and undefined thresholds default to never. The rule is now: **any diff over 80 lines, or any diff at any size touching goroutines, channels, mutexes, `context` propagation, or error handling, gets the full four-part treatment. Everything below that threshold gets section four alone.** The threshold may be tuned at a fortnightly review; it may never be undefined.

The rule applies to every diff in every language, without exception. This is where the supervision habit generalises, and it costs no rotation minutes.

---

## Ranked topics

Nine active. Deferrals named rather than ranked.

1. **Explain-before-accept, error-path weighted** — the standing rule above, plus a fixed weekly Go worksheet (Friday). The highest-leverage item in the plan for the third consecutive panel, and the only one with a fully specified shape. *Scheduling note, for consistency: from week 6, alternate Fridays cede their final ten minutes to the retro formats (rotation, below). The worksheet's slot never alternates; its length does, by design, every other week.*

2. **Go concurrency** — promoted from eighth. Concurrency defects are the class that survives review and fails in production, and the subject is currently the only filter on agent-generated concurrent code. Delivered failure-modes-first through a front-loaded three-week on-ramp (weeks 3–5), then drilled continuously as the weighting of the Thursday bug hunt. Formalism arrives last, as the explanation for a catalogue the subject already recognises. *Calibration, per the audit: by week 12 the subject will recognise the common concurrency defect classes in code shaped like their own. A working memory model is Phase 2-and-beyond material; the promotion buys recognition, not mastery, and claims nothing more.*

3. **Laddered build drills** — LRU cache, worker pool, rate limiter, trie, heap, graph traversal, by the four-rung ladder with the two-clean-runs advancement criterion. The backbone, and the source of everything Thursday sabotages.

4. **Bug hunts with controls, aged own code** — 7–14 day old implementations, three to five variants, zero to two defective, false positives scored. Seeded from the Go defect catalogue, weighted toward recently studied units and toward concurrency from week four.

5. **Go memory and data-structure internals** — promoted from seventh. Stack vs. heap, pointers, value semantics, escape analysis; slices, maps, hash tables as implemented. The vocabulary in which the subject forms hypotheses about their own production code. Pre-test gated.

6. **Hypothesis-driven debugging** — failing test, written hypothesis before any tool, generator mode only afterward for comparison. Runs in the same language as everything else and compounds with topics 4 and 5.

7. **Algorithm pattern library, Go, plus the idiom deck** — one pattern at a time with recognition cues written down, *and* a memorised set of Go snippets drilled cold until automatic: `container/heap` with a concrete type; a set via `map[T]struct{}`; `sort.Slice` with a custom comparator; a BFS queue on a slice; two pointers; a stack; a `strings.Builder` accumulation; binary search on an answer space. Ten minutes of idiom at the head of pattern days, dropping to spaced maintenance once each is automatic. Interview *volume* remains a gated sprint, not a standing cost.

8. **Big-O micro-unit, front-loaded** — two weeks, analysing the complexity of code the subject wrote that week. Makes every subsequent drill self-grading.

9. **Foundations flash deck** — spaced retrieval, author mode, five minutes riding *inside* DRILL sessions. Carries the knowledge layer so worksheets can carry the skill layer. *Budget note: the deck's five minutes come out of the host session — Thursday's hunt is honestly 25 minutes on deck days, and variant count should be sized to that (three, not five).*

**Phase 2, in order:** spec-first table-driven tests (first, trialled fortnightly from week six in the Friday retro position); generic-codebase bug hunts; mutation drills against the subject's own artifacts; monthly TypeScript type-level review; monthly design critique; prompt-as-spec. **Cross-cutting format, not a slot:** race-the-agent retros, fortnightly, in the last ten minutes of Friday, alternating with the spec-first trials.

---

## The sabotage catalogue

Author mode seeds Thursday hunts from this catalogue, weighted toward the units recently studied. The catalogue lives in the repository as a file; the subject appends to it every time a real bug reaches production or review.

*Memory and value semantics:* slice aliasing after `append` on a shared backing array; the result of `append` not reassigned; a large struct copied in a `range`; a value receiver mutating a copy; a method set mismatch between pointer and value; a `sync.Mutex` copied by value along with its struct.

*Maps and nil:* writing to a nil map; depending on map iteration order; dropping the second return of a lookup; a nil interface holding a typed nil pointer and therefore not equal to nil.

*Errors:* an error checked and swallowed; `fmt.Errorf` without `%w`; `==` where `errors.Is` was needed; a deferred close whose error is discarded; a shadowed `err` from `:=` in an inner scope.

*Concurrency:* a goroutine blocked forever on a send with no receiver; `WaitGroup.Add` called inside the goroutine rather than before it; an unbuffered channel deadlock; a context accepted but never propagated; `time.After` inside a loop leaking timers; a data race on a shared struct field; a mutex released on one path but not another.

*Resources:* `defer` inside a loop holding resources until function return; deferred arguments evaluated at defer time when execution time was meant.

**Version pinning.** The catalogue is pinned to the Go version in use and re-checked as a Sunday checklist item. (Loop-variable capture was the canonical gotcha for a decade; Go 1.22 changed the semantics to per-iteration variables. Drilling a defect the toolchain no longer produces is folklore.)

**[AUDIT — anti-stereotyping provisions.]** An agent seeding from a fixed catalogue develops stereotyped expressions of each defect within weeks — same variable names, same placement, same surrounding shape — and detection trained on stereotyped signal is detection of the stereotype, not the defect. Three provisions:
- **Author mode is explicitly instructed, in its standing prompt, to vary defect placement, naming, and surrounding code style** between variants and between weeks, and to disguise defects in visually unremarkable positions rather than at the point of interest.
- **The catalogue category for each week's hunt is drawn without the subject knowing which** — the weighting toward recent units biases the draw; it does not announce it.
- **Real bugs progressively displace synthetic ones.** As the real-bug list (rule 12) grows, Author mode seeds from it in preference to the synthetic catalogue. The synthetic catalogue is scaffolding; production is the syllabus.

---

## Rotation

| Day | Slot | Label | Min |
|---|---|---|---|
| Mon | Laddered build (3) | LEARN at rungs 1–2, DRILL at 3–4 | 30 |
| Tue | Algorithm pattern (7) — 10 min idiom deck, 20 min pattern. **Drop day, weeks 3–5 only** | LEARN new / DRILL review | 30 |
| Wed | Wks 1–2 Big-O (8); wks 3–5 concurrency on-ramp (2); wk 6+ internals (5) / debugging (6) alternating. **Drop day, week 6 onward** | LEARN | 30 |
| Thu | Bug hunt with controls (4), concurrency-weighted from wk 4 | DRILL | 30 |
| Fri | Explain-before-accept, error paths (1); from wk 6, alternate Fridays end with 10 min retro | DRILL | 30 |
| Sun | Prep: author sheets, refresh sabotage bank, triage log, version-check catalogue. **Floor: 10-min version (Mon + Thu only); template week covers the rest** | — | 30–45 |

Flash deck (9) rides inside DRILL sessions; the host session shrinks accordingly. Tokens are usually spent Monday and Wednesday.

**150 baseline + ~35 prep + up to 60 overtime = 185–245 minutes/week.**

**Phase 2, on the gate:** Wednesday takes monthly design critique and monthly TypeScript type-level review in place of alternating foundations; Thursday alternates own-code hunts with generic hunts and mutation drills; Friday rotates in spec-first table-driven tests as a full slot; pattern review drops to one spaced DRILL.

---

## Pre-week-1 checklist

Everything the patches require to exist before Monday of week 1:

- [ ] Repository created; single log file with columns for mistakes, spaced queue, rung state, token spend, error-path scores, adherence, and the **real-bug verdict line** (rule 12).
- [ ] Sabotage catalogue committed as a living file, version-pinned to the current Go toolchain.
- [ ] **Standing template week authored** — one default worksheet per slot (rule 8).
- [ ] **Work-diff threshold written into the standing rule as a number** (80 lines / any concurrency or error-handling touch) and posted where reviews happen.
- [ ] Author mode's standing prompt includes the **anti-stereotyping instructions** (catalogue section).
- [ ] Ten-minute floor versions written for all five weekday sheets and for Sunday.

---

## Remaining contingency

**B — the interview horizon.** Still unanswered; governs topic 7's volume only. The ruling assumes a hunt more than six months out. **Under six months:** the sprint starts now, topic 7 takes Tuesday and Thursday until done, bug hunts fall to fortnightly — *with the audit's caveat on the record: fortnightly hunts halve the drilling of topic 2, the plan's highest-expected-loss item, because Thursday is concurrency's delivery vehicle. The trade may still be right under interview pressure, but it is a real trade, and whoever takes that branch takes it knowingly.* **Over twelve months:** the library builds at one pattern a fortnight and Tuesday is shared with foundations. Nothing else moves either way.

---

## Dissents and standing risks

**Tanaka** maintains that nine active topics is too many for a thirty-minute box and predicts topics 7–9 are the first casualties of a bad month. *The audit largely concurs — and notes the plan degrades in the right order: if a bad month costs patterns, Big-O, and the flash deck, the supervision core (1, 2, 4, 6) survives, and that core answers the diagnosed problem. Losing 7–9 for a month is a bruise, not a failure.* Tanaka separately dissents from the monoglot concentration; **Phase 3, month six or on Phase 2 completion, whichever is later, re-opens the language question**, when a consolidated first model makes a second language cheap.

**Ferreira** dissents from the denial of a TypeScript slot, holding that supervision habits generalise from practice rather than stated rules. Named risk on the record: frontend work receives no dedicated practice for roughly twelve weeks; the mitigation is the standing rule and its now-numeric threshold. If the balance of work shifts materially toward frontend, this ruling is the first thing to revisit.

**Iglesias** notes that drilling algorithms in Go costs real throughput against Python; the idiom deck mitigates but does not eliminate it.

**Feld** notes that spec-first table tests are, in an all-Go codebase, ranked below their true value by schedule rather than merit; they open Phase 2 and trial fortnightly from week six.

---

**Ruling entered, audit incorporated, five patches integrated. The next audit is the week-12 gate, conducted against the log — not another panel round. The document has been argued over four times; what it needs now is contact with a Monday.**
