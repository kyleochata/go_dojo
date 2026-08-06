# SEALED — 2026-w1 bug-hunt answer key

Do not open outside /score or the Sunday verifier pass
(CLAUDE.md rule 1).

Toolchain: go1.26.5 darwin/arm64 (matches the pin at the top of
catalogue.md).

## Sampling outcome

See `sampling.json` in this directory. This key is written to match
it exactly; if the two ever disagree, `sampling.json` is authoritative
and the week was authored wrong.

- variants: 3
- defective variants: 0
- planted defects: none

**Every variant in `worksheets/2026-w1/hunt/` is CLEAN.** The correct
verdict on all three is *approve*, worth 10/10 each under rubric R4c.
Any claimed defect is a false positive and costs −2 under F1. Maximum
score for the week is 30/30; "rejected one variant on a
plausible-looking construct" is 28/30 at best, and that is the exact
failure this hunt is built to measure.

## Verification performed (this session, go1.26.5)

Run from `worksheets/2026-w1/hunt/`:

| check | result |
|---|---|
| `gofmt -l .` | no output |
| `go build ./...` | success |
| `go vet ./...` | no findings |
| `go test -count=1 ./...` | 33 tests pass across 3 packages (26 top-level: 10 / 8 / 8, plus 7 subtests) |
| `go test -race -count=2 ./...` | 66 passes, no race reports |

Two fresh-context verifier subagents (skills/author-variant/SKILL.md
§5) independently audited all three variants against the full
catalogue. **Neither found a correctness defect.** Both rounds of tell
findings and gray-area findings were acted on before sealing — see
"Verifier rounds" at the bottom.

No defect reproduces because none was planted. The reproduction
requirement in §4.3 of the authoring skill is vacuous this week; §4.4
(no unintended defects) is the load-bearing one, and the per-variant
analysis below plus the two verifier rounds is the evidence for it.

## Variant 1 — package `ledger` (`variant1/ledger.go`)

Clean. Constructs a reviewer is expected to stop on, and why each is
correct:

| function | the tempting reading | why it is actually correct |
|---|---|---|
| `Book.Post` — `b.order = append(b.order, e.Account)` | "append result not reassigned" | it *is* reassigned, to the struct field. |
| `Book.Post` — `b.accounts[e.Account] = acct` | "writing to a nil map on a zero-value `Book`" | the lazy `if b.accounts == nil` init two lines above makes the zero value usable, and the type comment says so. `TestZeroBook` pins it. |
| `Book.Snapshot` — `out = append(out, *b.accounts[id])` | "returns internal state; caller can mutate the book" | `Account` is `{string, int64, int, string}` — no reference-typed fields — so the dereference is a full deep copy. Verified by probe: mutating the returned slice leaves the book unchanged. |
| `Account.String` value receiver | "value receiver mutating a copy" | it only reads its receiver and returns a new string. |
| `Book.Total` ranges a map | "depends on map iteration order" | addition is commutative; the result is order-independent. Probed stable over 200 iterations. |
| `DistinctTags` ranges a map into a slice | "depends on map iteration order" | `sort.Strings` immediately after makes the output deterministic. `TestDistinctTags` pins it. |
| `Filter` — inner loop with `break` | "breaks the wrong loop / duplicates entries" | `break` leaves the inner tag loop only, after appending exactly once per matching entry. `TestFilter` pins both hit and miss counts. |
| `Split` — `credits = append(credits, e)` from a nil slice | "aliases the caller's backing array" | appending to a nil slice allocates fresh storage, and `Entry` is copied by value into it. Writes to `credits`/`debits` do not reach the input. |
| `Normalize` — `for i := range entries` | "range copies the struct, so the write is lost" | it indexes (`entries[i].Memo = ...`); it never writes to a range value copy. |
| `Book.Top` — `snap[:n]` | "unclamped slice index" / "returns a slice into shared state" | `n` is clamped both above and below immediately before the slice expression; and `snap` is the caller-owned copy from `Snapshot`, which the book never retains. |
| `Balance` — `fmt.Errorf("ledger: balance %q: %w", ...)` | "sentinel lost / double-prefixed" | `%w` preserves `ErrUnknownAccount`, and the sentinel text is unprefixed precisely so the wrap reads once. `TestBalance` asserts `errors.Is`. |

**One share is real and is NOT a defect:** `Entry.Tags` is a slice, so
the `Entry` values `Split` and `Filter` copy into their outputs share
tag backing arrays with the caller's input. Nothing in the package
writes through a `Tags` slice, and neither function documents a
deep-copy guarantee, so there is no defect. But note the honest limit
of that argument: `Split` and `Filter` are *exported*, so a caller
could write through the shared `Tags` and be surprised. Adjudicate as
follows — a **defect claim** ("Split corrupts the caller's data") is a
false positive, F1 −2; a **design observation** offered alongside a
correct "approve" ("shares Tags backing arrays; I'd document or
deep-copy") is not a defect claim, costs nothing, and is worth
calling out in the reveal as good reviewing.

## Variant 2 — package `dispatch` (`variant2/dispatch.go`)

Clean. This is a worker pool, so it necessarily contains the shapes
the concurrency catalogue is written about. That is the domain, not a
signal.

| line / function | the tempting reading | why it is actually correct |
|---|---|---|
| `wg.Add(workers)` before the spawn loop | "`WaitGroup.Add` inside the goroutine" | `Add` is called once, before any goroutine starts. |
| `results <- Result{...}` with no select | "goroutine blocked forever on a send" | `results` is buffered to `len(jobs)`; each job is sent to `queue` at most once and produces at most one result, so at most `len(jobs)` sends occur and no send can ever block. **This buffer sizing is what makes the workers deadlock-free** — not the feeder's select. |
| `feed`'s `select` on `ctx.Done()` | "cancellation is only checked between jobs, so a long handler ignores it" | true and intended: `ctx` is handed to the handler, which owns its own responsiveness. `feed`'s job is to stop *queuing* work. |
| `close(queue)` after `feed` returns | "unbuffered channel deadlock" | workers cannot get stuck (their only blocking operation is the receive), and `close(queue)` terminates every `range queue`. The feeder cannot get stuck either — a worker is always available to receive until `queue` closes, and the workers only exit *because* it closed. Probed at 1000 workers × 1 job and 1 worker × 5000 jobs. |
| `close(results)` after `wg.Wait()` | "close while senders are live" | `wg.Wait()` returns only after every sender has left its range loop. |
| `Stats.record` — early `return` in the failure branch | "mutex released on one path but not the other" | the unlock is `defer`red, so both paths release. |
| `Stats` holds a `sync.Mutex` field | "mutex copied by value along with its struct" | `Stats` is only ever handled through `*Stats`: both methods have pointer receivers, `Run` builds it as `&Stats{}` and returns the pointer, and the field is **named, not embedded**, so no `Lock`/`Unlock` is promoted onto the value type. It is never copied. |
| `h(ctx, j)` | "context accepted but never propagated" | `ctx` is passed straight to the handler, the only downstream call. |
| `stats.record(err)` from N goroutines | "data race on a shared struct field" | every access to `ok`/`failed` is inside `s.mu`. `-race` over 66 runs found nothing; a probe hammering `Stats` concurrently found nothing. |
| result ordering | "channel order is nondeterministic" | `sort.Slice` by `JobID` before returning. `TestRun` pins the full ordering. |
| the two cancellation checks | "swallows the cancellation" / "reports failure on a run that finished" | the up-front check rejects an already-dead context; the final one fires **only if `feed` returned false**, i.e. only if jobs were actually left unqueued. A run whose context is cancelled after the last handler returns reports success, because it was one. |
| `Run` on an empty job list | "returns nil where other paths return a slice" | `collected` is allocated before the early return, so every success path returns a non-nil slice. `TestRunEmptyJobs` asserts non-nil. |

Goroutine-leak probe: 200 cancelled `Run`s of 200 jobs each left the
goroutine count unchanged (2 → 2). Mid-run-cancel invariants held
over 50 trials: results always sorted, never duplicated, and
`ok + failed == len(results)`.

## Variant 3 — package `manifest` (`variant3/manifest.go`)

Clean.

| line / function | the tempting reading | why it is actually correct |
|---|---|---|
| `Load`'s deferred close | "deferred close whose error is discarded" | the deferred func captures the named return, and on `cerr != nil && err == nil` sets `m, err = nil, ...` — so the close error is surfaced and no caller ever sees a non-nil manifest beside a non-nil error. |
| named returns `(m *Manifest, err error)` with `parsed, err := Parse(f)` | "shadowed `err` from `:=` in an inner scope, so the defer sees nil" | not a shadow. `:=` at the top level of a function body reuses variables from the result list (Go spec: redeclaration applies to "the parameter lists if the block is the function body"), and only `parsed` is new. Verified empirically with a minimal reproduction: the deferred closure observes the wrapped error. |
| `LoadAll` | "`defer` inside a loop holds every file open until return" | there is no `defer` in the loop; `Load` opens and closes each file within one call. Probed: no fd leak across 15,000 `Load` calls including error paths. |
| `LoadAll` — `errors.Is(err, os.ErrNotExist)` | "`==` where `errors.Is` was needed" | it uses `errors.Is`, and every wrap between `os.Open` and here uses `%w`, so the sentinel survives. Probed: a permission error is correctly **not** skipped. |
| `Parse` error returns | "`fmt.Errorf` without `%w`" | all seven either use `%w` or return a bare sentinel. `TestParseErrors` asserts `errors.Is` across seven subtests and `errors.As` on the `*strconv.NumError` — eight assertions in total. |
| `Parse` — the `seen` set | "duplicate detection misses the reserved keys" | `seen` is checked before the `switch`, so `name` and `replicas` are covered. The "duplicate reserved key" subtest pins it. |
| `Parse` — `sc.Buffer(..., maxLine)` | "`bufio.Scanner` dies at 64 KiB on a long value" | the buffer limit is raised to 1 MiB. `TestParseLongValue` parses a 200 KB value. |
| `Parse` — `replicas` validation | "accepts 0 or negative, then `Merge` silently drops them" | `n < 1` is rejected with `ErrBadReplicas`, so `Merge`'s `src.Replicas > 0` test means exactly "unset", and the doc comment says so. |
| `Merge` — `m.Env[k] = v` | "writing to a nil map" | guarded by `if m.Env == nil { m.Env = make(...) }`, unconditionally, so an empty-but-non-nil `src.Env` behaves the same as a populated one. `TestMerge` merges into a zero-value `&Manifest{}`. |
| `Merge(nil)` | "nil dereference" | early `if src == nil { return }`. |
| `Clone` | "shallow copy shares the map" | allocates a new map and copies every pair; a nil `Env` stays nil, pinned by `TestClone`. |
| `Render` — `fmt.Fprintf` return values ignored | "swallowed error" | `bufio.Writer` is sticky: it holds the first error and returns it from `Flush`, which *is* checked and wrapped. Probed with a failing writer. |
| `Parse` defaulting `Replicas` to 1 | "silently defaults" | documented in the package comment, pinned by `TestParseDefaultReplicas`. |

Adjudications for the remaining gray areas, so /score has a ruling
for each:

- **Nil receivers.** `Merge`, `Clone`, and `Render` all dereference
  `m` without a nil check, uniformly. None documents a nil contract.
  A `(*Manifest)(nil)` panic is ordinary Go; claiming it as a defect
  is F1.
- **`A=b=c`** cuts at the first `=`, which is what `strings.Cut`
  promises and what key=value formats conventionally do.
- **Error prefixing.** `manifest:` appears once per error chain:
  sentinels carry it, `Load` adds it for I/O failures, and `LoadAll`
  wraps with a bare `load all:`. A claim of "double-prefixed errors"
  is a false positive.

## Scoring notes for /score

- Three "approve" verdicts → 30/30 (R4c ×3).
- Each claimed defect: −2 (F1), citing the row above that explains
  why the construct is correct.
- If the human rejects a variant, name the specific near-miss they
  fell for; the category they mis-fired on goes to the spaced
  re-attempt queue exactly as a miss would.
- A zero-defective hunt is the control condition — the only way to
  measure false-positive rate. Do not apologise for it in the
  reveal, and do not treat 30/30 as an easy week.

## Verifier rounds (fresh-context subagents, before sealing)

**Round 1 — tells acted on:** tests whose names and failure messages
read as an answer key (`TestSnapshotIsACopy`, "merge reached back
into the base") rewritten as ordinary behavioural tests; doc comments
that pre-asserted the exact property the key defends ("the caller …
cannot reach into the book through them"; "A failure to close … is
reported as an error"; "m is modified; src is not") reduced to plain
documentation; `Account.String`'s `strings.Builder`-around-a-`Sprintf`
replaced with one `fmt.Sprintf`; `Top`'s two-branch bounds block
collapsed; a hand-rolled insertion sort in variant3 replaced with
`sort.Strings`; the no-op `workers > len(jobs)` clamp removed;
`Entry.Memo`/`Entry.Tags` made load-bearing via `Account.LastMemo`
and `Filter`.

**Round 1 — gray areas closed in code, not adjudicated away:**
`ErrDuplicateKey` was documented as "a key defined twice in one file"
while detection lived only in the `default:` branch, so
`name=a\nname=b` was accepted silently; and `Run` returned a nil error
for an empty job list under an already-cancelled context while the
non-empty path returned the wrapped `ctx.Err()`.

**Round 2 — gray areas closed in code:** zero-value `Book` panicked
on a nil-map write (nil-map write is itself a catalogue category, so
this was the highest-risk gray area in the set); `Run` could return a
complete, all-successful result set *plus* an error when the context
was cancelled after the last handler returned; `Merge` panicked on a
nil receiver while `Clone` handled one, three functions apart in the
same file; `replicas=0` and `replicas=-3` parsed silently and were
then dropped by `Merge`; `bufio.Scanner` sat at its 64 KiB default;
`Run` returned a nil slice for empty input where every other path
returned a non-nil one; `Merge` with an empty-but-non-nil `src.Env`
left `m.Env` nil.

**Round 2 — tells acted on:** the `feed:` label (the only label in
the corpus) replaced with a `feed` helper function; the
`max(0, min(...))` clamp idiom no longer appears in both `ledger.go`
and `analyze.go`; the assertion-free `fresh.Merge(nil)` line and the
nil-receiver `Clone` test removed; `TestRunOrdersResults` and
`TestRenderIsSorted` renamed to neutral `TestRun` / `TestRender`;
`errBoom` moved to the top of its file; variant3's triple-nested
error prefixes flattened; `for i := 0; i < workers; i++` changed to
`for range workers` to match the range idiom used elsewhere;
variant1's sentinel/wrap prefix inconsistency fixed. Comment density
across the three variants moved from 0.150 / 0.147 / 0.075 to
**0.172 / 0.196 / 0.157**, and top-level test counts from 9 / 5 / 7
to **10 / 8 / 8**.

**Accepted and not changed, with reasons:**

- Variant 2 carries more concurrency near-misses than the others. It
  is a worker pool; the shapes are inherent to the domain.
- The three domains (ledger / dispatch / manifest) map onto three
  catalogue families. Any set of three varied Go packages will, and
  spreading each family across every variant would be a stronger
  tell, not a weaker one.
- `Load`'s named-return-plus-defer is the only such idiom in the
  corpus, but it is *the* standard Go idiom for surfacing a close
  error; removing it would mean discarding that error — a real
  defect — to hide a stylistic fingerprint.
- Doc comments on `Clone` ("copy sharing no state"), `Render`
  ("sorted order"), and `Normalize` ("in place") do state the
  property a near-miss would violate. They are also the contract any
  real package documents; stripping them would leave an
  undocumented exported API, which is a louder anomaly.
