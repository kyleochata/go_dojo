# Idiom 04 — BFS queue on a slice

## Canonical

```go
package idiom

// bfs returns nodes in visit order from start.
func bfs(adj map[int][]int, start int) []int {
	seen := map[int]struct{}{start: {}}
	queue := []int{start}
	var order []int

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)

		for _, next := range adj[cur] {
			if _, ok := seen[next]; ok {
				continue
			}
			seen[next] = struct{}{} // mark on enqueue, not on dequeue
			queue = append(queue, next)
		}
	}
	return order
}

// bfsLevels is the same walk, batched one hop at a time — the form
// you need whenever the answer is a distance or a depth.
func bfsLevels(adj map[int][]int, start int) [][]int {
	seen := map[int]struct{}{start: {}}
	frontier := []int{start}
	var levels [][]int

	for len(frontier) > 0 {
		levels = append(levels, frontier)
		var next []int
		for _, cur := range frontier {
			for _, n := range adj[cur] {
				if _, ok := seen[n]; ok {
					continue
				}
				seen[n] = struct{}{}
				next = append(next, n)
			}
		}
		frontier = next
	}
	return levels
}
```

## Type it cold

No docs, no scrolling up. Compile it.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. Marking `seen` on enqueue vs. on dequeue: construct a graph where
   the dequeue version visits a node twice. How bad does it get?
2. `queue = queue[1:]` never shrinks the backing array. What is the
   memory behaviour over a long walk, and when would you care enough
   to use an index cursor or `container/list` instead?
3. In `bfsLevels`, why build a fresh `next` slice each round instead
   of appending onto `frontier` and tracking a length? What is the
   one-liner that turns levels into a distance map?

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
