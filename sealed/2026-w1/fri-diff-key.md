# SEALED — 2026-w1 Friday key (explain-before-accept)

Open only after all four sections are written (CLAUDE.md rule 3).

Subject: `worksheets/2026-w1/fri-diff/change.diff` applied to
`worksheets/2026-w1/fri-diff/sessions/store.go`.

Verified on go1.26.5: the diff applies cleanly with `git apply`, the
result is `gofmt`-clean, builds, and **`go vet` reports nothing**.
Both defects below were reproduced this session.

## Which threshold clause fired

The concurrency clause: the diff adds a goroutine (`sweepLoop`), a
new channel (`stop`), and code that reads and writes shared state
under — and outside — a mutex. It also adds six error paths. It
would take all four sections at any line count. Naming *the clause*,
not just "it's big", is the scored half of that line.

## §1 — what the change does

There is no honest one-sentence answer, and that is the finding. The
diff does **four** separable things:

1. adds a TTL and an `ErrExpired` sentinel to reads,
2. adds a background sweeper goroutine plus `Close` to stop it,
3. changes `Get`'s signature from `(Session, error)` to
   `(*Session, error)` — a breaking API change, and the source of
   both defects,
4. adds `Persist`/`Restore` (JSON file I/O).

Full marks for "it does four things, here they are"; half for
naming three; a smooth one-sentence summary that papers over the mix
scores zero on this section even if it reads well.

## §2 — the error paths (reference table)

| # | what produces it | what the code does | what the caller can distinguish |
|---|---|---|---|
| 1 | `Get`: token absent | returns `ErrNotFound` | yes — sentinel, comparable with `errors.Is` |
| 2 | `Get`: session older than ttl | returns `ErrExpired` | yes — new sentinel, and correctly distinct from #1 |
| 3 | `Touch`: propagates `Get`'s error | returns it unwrapped | yes, but the caller cannot tell *which call* failed; acceptable here since `Touch` has one failure source |
| 4 | `Persist`: `os.Create` fails | wrapped with `%w` | yes — `os.ErrPermission` etc. survive |
| 5 | `Persist`: `json.Encoder.Encode` fails | wrapped with `%w` | partially — **see defect 2**: the deferred `f.Close()` error is discarded, so a write that fails at flush time is invisible |
| 6 | `Restore`: `os.ReadFile` fails | wrapped with **`%v`** | **no** — see defect 1 |
| 7 | `Restore`: `json.Unmarshal` fails | wrapped with `%w` | yes |
| 8 | `Sweep`, `Close`, `Delete`, `Len`, `Put` | no error paths at all | worth writing down: `Close` on an already-closed store is safe via `sync.Once`, and `Close` on a store from `New()` (nil `stop`) is guarded — both correct |

Rows 6 and 5 are the scored ones. A table that lists all eight rows
but marks column four "yes" for row 6 has missed the point of the
section.

## The defects

### Defect 1 — `Restore` wraps with `%v`, destroying the sentinel

`fmt.Errorf("sessions: restore %s: %v", path, err)` — every other new
error in the diff uses `%w`. This one does not.

Consequence: a caller doing
`if errors.Is(err, os.ErrNotExist) { /* first boot, start empty */ }`
gets `false` and treats a missing snapshot file as a hard failure.
The error message still *reads* correctly, which is why this survives
review.

Reproduction (go1.26.5, confirmed this session):

```go
err := New().Restore("/nonexistent/path/x.json")
// err = "sessions: restore /nonexistent/path/x.json: open
//        /nonexistent/path/x.json: no such file or directory"
// errors.Is(err, os.ErrNotExist) == false
```

Fix: `%w`.

### Defect 2 — `Get` now returns the stored pointer; `Touch` mutates it outside the lock

This is the one to find, and it is the answer to §4.

`Get` takes `s.mu.RLock()`, then returns `sess` — the `*Session` that
lives in the map — and releases the lock on return. The pointer
outlives the lock. `Touch` then does:

```go
sess, err := s.Get(token)   // RLock taken and released
if err != nil { return err }
sess.Created = time.Now()   // unsynchronized write to shared state
```

Consequence: an unsynchronized write to `sess.Created` racing against
the sweeper goroutine's read of the same field in `Sweep` (which
holds the write lock, but the writer does not). Two concurrent
`Touch` calls race with each other as well. Under the race detector
this fires immediately; in production it is a torn/stale timestamp,
so a session is swept while in active use, or never swept at all.

It is not only `Touch`: **any** caller of the new `Get` now holds a
pointer into the store's protected state and can mutate `Scopes` or
`UserID` with no lock at all. The pre-image deliberately returned a
`Session` value with a copied `Scopes` slice; the diff silently
deletes that protection. The `Put` method still copies on the way
in — the asymmetry is the tell, and it is the kind of tell that only
shows up if you read the *pre-image* alongside the diff.

Reproduction (go1.26.5, confirmed this session). Apply the diff, set
`ttl` on a store, then run 500 `Touch("a")` calls concurrently with
500 `Sweep()` calls under `go test -race`:

```go
s := New()
s.ttl = time.Hour
s.Put(Session{Token: "a", UserID: "u", Created: time.Now()})
// goroutine 1: for i := 0; i < 500; i++ { _ = s.Touch("a") }
// goroutine 2: for i := 0; i < 500; i++ { s.Sweep() }
```

`WARNING: DATA RACE` — the two ends are, in the applied post-image,
**store.go:121** (`sess.Created = time.Now()` in `Touch`, unlocked)
and **store.go:75** (`sess.Created.Before(cutoff)` in `Sweep`, under
the write lock). `Get`'s own read at store.go:109 races the same way.

Fix: either keep returning a value copy and give `Touch` its own
locked implementation, or document `Get`'s result as read-only and
have `Touch` take `s.mu.Lock()` and mutate the map entry directly.

### Also true, but not defects

- `Persist` holds `RLock` across file I/O — every reader blocks for
  the duration of a disk write. A real latency concern, worth
  writing down under §3 or §4, but it is correct code. Claiming it as
  a bug is a false positive.
- `Persist` serializes `[]*Session` taken from a map, so the JSON
  order varies run to run. Harmless here (`Restore` rebuilds by
  token) but it makes the file un-diffable.
- `Restore` does not respect `ttl`: it loads expired sessions, which
  the next `Sweep` or `Get` then rejects. Self-correcting, so a
  design note, not a defect.
- `NewWithTTL` starts a goroutine that outlives the constructor. Not
  a leak — `Close` exists and `sync.Once` makes it idempotent — but
  a caller who forgets `Close` leaks the ticker goroutine for the
  process lifetime. Legitimate §4 material.
- `f.Close()` on the `os.Create` path discards its error. On most
  filesystems a failed close after a successful encode means the
  data did not reach the disk. Listing this in §2 row 5 is worth
  full credit; calling it *the* most-likely-wrong line is a
  defensible §4 answer, though defect 2 is the stronger one.

## §3 — copies and shares (reference)

| place | copy or share | intended |
|---|---|---|
| `Put`: `cp := sess` then `cp.Scopes = append([]string(nil), ...)` | copy, deep on `Scopes` | yes — unchanged from the pre-image |
| `Get` (new): `return sess, nil` | **share** — pointer into the map | **no** — this is defect 2 |
| `Sweep`: `for token, sess := range s.byID` | share (the values are pointers) | yes, and safe: it is under `Lock` |
| `Persist`: `all = append(all, sess)` | share, then serialized under `RLock` | yes, but see the latency note |
| `Restore`: `s.byID[sess.Token] = sess` | share — the decoded pointers become the store's state | yes; they are freshly decoded and unreachable elsewhere |
| `Touch`: `sess.Created = time.Now()` | write through a shared pointer | **no** — defect 2 |

A §3 table that names the `Get` row as a share **and** marks it
unintended has effectively found defect 2 even if §4 names a
different line. Score it as found.

## Scoring (rubric S-lines)

- S1 (+3): both defects named. One of two → +1.
- S2 (+3): the *mechanism* for each — "pointer outlives the RLock",
  "%w vs %v breaks errors.Is" — not just "there's a race".
- S3 (+2): boundary cases addressed — `Close` on a `New()` store
  (nil channel), double `Close`, `Restore` of a missing file,
  `Sweep` with `ttl == 0`.
- S4 (+2): the §2 table covers all eight error paths with column
  four filled honestly.
- Deduct for false positives on the "also true, but not defects"
  list, same spirit as F1.
