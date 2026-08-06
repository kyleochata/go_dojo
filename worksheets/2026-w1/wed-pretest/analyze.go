// Package bigo holds the Wednesday pre-test subjects. Every function
// here is correct — the drill is about cost, not correctness.
package bigo

import (
	"sort"
	"strings"
)

// A. n = len(xs)
func hasDuplicate(xs []int) bool {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[i] == xs[j] {
				return true
			}
		}
	}
	return false
}

// B. n = len(xs)
func hasDuplicateSeen(xs []int) bool {
	seen := make(map[int]struct{})
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}

// C. n = len(names), and assume every name has the same length L.
func report(names []string) string {
	out := ""
	for _, n := range names {
		out += n + "\n"
	}
	return out
}

// D. n = len(names), same assumption as C.
func reportBuilder(names []string) string {
	var sb strings.Builder
	for _, n := range names {
		sb.WriteString(n)
		sb.WriteString("\n")
	}
	return sb.String()
}

// E. n = len(xs)
func topK(xs []int, k int) []int {
	cp := make([]int, len(xs))
	copy(cp, xs)
	sort.Sort(sort.Reverse(sort.IntSlice(cp)))
	if k < 0 {
		k = 0
	}
	if k > len(cp) {
		k = len(cp)
	}
	return cp[:k]
}

// F. n = len(sorted)
func contains(sorted []int, v int) bool {
	lo, hi := 0, len(sorted)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case sorted[mid] == v:
			return true
		case sorted[mid] < v:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return false
}

// G. g = len(groups), and every group has n elements.
func anyGroupHasDuplicate(groups [][]int) bool {
	for _, g := range groups {
		if hasDuplicate(g) {
			return true
		}
	}
	return false
}

// H. n = len(xs). What does this cost, and why is the answer not
// "O(n^2) because append may reallocate"?
func evens(xs []int) []int {
	var out []int
	for _, x := range xs {
		if x%2 == 0 {
			out = append(out, x)
		}
	}
	return out
}
