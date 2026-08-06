# Friday target diff — week 2026-w1

- `sessions/store.go` is the **pre-image**: the package as it stands today.
- `change.diff` is the change under review, produced by an agent asked
  to "add TTL expiry and disk persistence to the session store."

The diff touches mutexes, a goroutine, and error handling, so the
standing threshold puts it in the **all four sections** bucket
regardless of its size.

Do not apply it until sections 1–4 are written. To apply afterwards:

```
cd worksheets/2026-w1/fri-diff
git apply change.diff       # or: patch -p1 < change.diff
go build ./... && go vet ./...
```

`git checkout -- sessions/store.go` puts it back.
