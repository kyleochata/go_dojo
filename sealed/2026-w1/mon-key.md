# SEALED — 2026-w1 Monday reference (LRU cache)

Do not open outside /score, the Sunday verifier pass, or a rung-1
worked-example request that the human has explicitly asked for.
Rungs 2–4 open this only AFTER the attempt (CLAUDE.md rule 3).

Verified on go1.26.5: `gofmt -l` clean, `go vet` clean, four tests
pass.

## Reference implementation

```go
// Package lru is a fixed-capacity least-recently-used cache.
package lru

import "container/list"

type entry[K comparable, V any] struct {
	key K
	val V
}

// Cache is not safe for concurrent use.
type Cache[K comparable, V any] struct {
	cap   int
	ll    *list.List          // front = most recently used
	items map[K]*list.Element // key -> element holding *entry
}

// New returns a cache holding at most capacity entries. capacity
// must be positive.
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("lru: capacity must be positive")
	}
	return &Cache[K, V]{
		cap:   capacity,
		ll:    list.New(),
		items: make(map[K]*list.Element, capacity),
	}
}

func (c *Cache[K, V]) Len() int { return c.ll.Len() }

// Get returns the value for key and promotes it to most-recently-used.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	el, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.ll.MoveToFront(el)
	return el.Value.(*entry[K, V]).val, true
}

// Put inserts or updates key, promoting it, and evicts the
// least-recently-used entry if that pushed the cache over capacity.
func (c *Cache[K, V]) Put(key K, val V) {
	if el, ok := c.items[key]; ok {
		el.Value.(*entry[K, V]).val = val
		c.ll.MoveToFront(el)
		return
	}
	el := c.ll.PushFront(&entry[K, V]{key: key, val: val})
	c.items[key] = el
	if c.ll.Len() > c.cap {
		c.evictOldest()
	}
}

func (c *Cache[K, V]) evictOldest() {
	el := c.ll.Back()
	if el == nil {
		return
	}
	c.ll.Remove(el)
	delete(c.items, el.Value.(*entry[K, V]).key) // both structures, or it leaks
}
```

## Acceptance criteria, mapped to the code

| criterion | where it is satisfied | the usual failure |
|---|---|---|
| capacity <= 0 handled | `New` panics with a documented message | silently accepting 0, so the first `Put` evicts the thing it just inserted |
| `Get` O(1) | map lookup + `MoveToFront` (pointer splice) | scanning a slice to find the key |
| `Put` O(1) | map lookup, `PushFront`, `Back`, `Remove` are all O(1) | scanning to find the LRU entry — correct, passes every test, and O(n). This is the Wednesday Big-O anchor. |
| `Get` promotes on hit | `c.ll.MoveToFront(el)` | reading through the map without touching the list, so nothing is ever promoted and the cache degrades to FIFO |
| `Put` on existing key updates and promotes without growing | the `if el, ok := c.items[key]; ok` branch returns early | pushing a second element for the same key: the map points at the new one, the old one lingers in the list forever |
| eviction removes from BOTH structures | `evictOldest` does `ll.Remove` then `delete(c.items, ...)` | removing from the list only. Nothing observable breaks; the map grows without bound. A slow leak, not a wrong answer — this is the criterion most worth checking on a rung-4 comparison. |

## Rung-4 comparison prompts

If Monday runs as race-the-agent, compare on exactly these and write
one sentence per divergence:

1. Where does the key live? (An implementation that stores only the
   value in the list cannot evict — it has no way back to the key.)
2. What happens on `Put` of an existing key at capacity?
3. Is `Get`'s promotion inside or outside the map lookup?
4. Does `Len()` read the list or the map, and would the two ever
   disagree?
