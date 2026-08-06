# Idiom 03 — `sort.Slice` with a custom comparator

## Canonical

```go
package idiom

import "sort"

type person struct {
	name string
	age  int
}

// Multi-key: age ascending, then name ascending as the tiebreak.
func sortPeople(ps []person) {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].age != ps[j].age {
			return ps[i].age < ps[j].age
		}
		return ps[i].name < ps[j].name
	})
}
```

Modern equivalent — `slices.SortFunc` takes elements, not indices,
and its comparator returns an int like `strings.Compare`:

```go
package idiom

import (
	"cmp"
	"slices"
)

func sortPeopleModern(ps []person) {
	slices.SortFunc(ps, func(a, b person) int {
		if c := cmp.Compare(a.age, b.age); c != 0 {
			return c
		}
		return cmp.Compare(a.name, b.name)
	})
}
```

## Type it cold

No docs, no scrolling up. Compile it. Type both forms.

```go

```

- [ ] Attempt 1 — clean / stumbled / failed
- [ ] Attempt 2 — clean / stumbled / failed

## Recall prompts

1. The comparator is `less(i, j) bool`. What exactly must it return
   for equal elements, and what breaks if you return `true` there?
2. `sort.Slice` vs `sort.SliceStable` — when does the difference
   actually change your output? Tie it to the multi-key example.
3. Write the descending version of both forms without adding a
   negation to the wrong place. Which argument order flips?

## Card state

Automatic count: ☐ ☐  (two clean cold runs → spaced maintenance)
