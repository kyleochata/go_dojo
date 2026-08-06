// Package ledger aggregates posted entries into per-account balances.
package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// ErrUnknownAccount is returned for lookups of accounts that have
// never been posted to.
var ErrUnknownAccount = errors.New("unknown account")

// Entry is a single posting against an account. Amount is in minor
// units; negative amounts are debits.
type Entry struct {
	Account string
	Amount  int64
	Memo    string
	Tags    []string
}

// Account is the running state of one account in the book.
type Account struct {
	ID       string
	Balance  int64
	Postings int
	LastMemo string
}

// String renders the account for logs and CLI output.
func (a Account) String() string {
	return fmt.Sprintf("%s %d (%d postings)", a.ID, a.Balance, a.Postings)
}

// Book holds the accounts seen so far and the order they first
// appeared in. The zero value is ready to use.
type Book struct {
	accounts map[string]*Account
	order    []string
}

// NewBook returns an empty book.
func NewBook() *Book {
	return &Book{accounts: make(map[string]*Account)}
}

// Post applies one entry, creating the account on first sight.
func (b *Book) Post(e Entry) {
	if b.accounts == nil {
		b.accounts = make(map[string]*Account)
	}
	acct, ok := b.accounts[e.Account]
	if !ok {
		acct = &Account{ID: e.Account}
		b.accounts[e.Account] = acct
		b.order = append(b.order, e.Account)
	}
	acct.Balance += e.Amount
	acct.Postings++
	acct.LastMemo = e.Memo
}

// PostAll applies every entry in order.
func (b *Book) PostAll(entries []Entry) {
	for _, e := range entries {
		b.Post(e)
	}
}

// Balance reports the current balance of one account.
func (b *Book) Balance(id string) (int64, error) {
	acct, ok := b.accounts[id]
	if !ok {
		return 0, fmt.Errorf("ledger: balance %q: %w", id, ErrUnknownAccount)
	}
	return acct.Balance, nil
}

// Snapshot returns the accounts in first-posting order.
func (b *Book) Snapshot() []Account {
	out := make([]Account, 0, len(b.order))
	for _, id := range b.order {
		out = append(out, *b.accounts[id])
	}
	return out
}

// Total sums every balance in the book.
func (b *Book) Total() int64 {
	var total int64
	for _, acct := range b.accounts {
		total += acct.Balance
	}
	return total
}

// Top returns the n accounts with the largest balances, ties broken
// by ID.
func (b *Book) Top(n int) []Account {
	snap := b.Snapshot()
	sort.Slice(snap, func(i, j int) bool {
		if snap[i].Balance != snap[j].Balance {
			return snap[i].Balance > snap[j].Balance
		}
		return snap[i].ID < snap[j].ID
	})
	n = max(0, min(n, len(snap)))
	return snap[:n]
}

// Split partitions entries into credits (Amount >= 0) and debits.
func Split(entries []Entry) (credits, debits []Entry) {
	for _, e := range entries {
		if e.Amount >= 0 {
			credits = append(credits, e)
		} else {
			debits = append(debits, e)
		}
	}
	return credits, debits
}

// Normalize trims surrounding whitespace from every memo, in place.
func Normalize(entries []Entry) {
	for i := range entries {
		entries[i].Memo = strings.TrimSpace(entries[i].Memo)
	}
}

// Filter returns the entries carrying the given tag.
func Filter(entries []Entry, tag string) []Entry {
	var out []Entry
	for _, e := range entries {
		for _, t := range e.Tags {
			if t == tag {
				out = append(out, e)
				break
			}
		}
	}
	return out
}

// DistinctTags lists every tag used across entries, sorted.
func DistinctTags(entries []Entry) []string {
	seen := make(map[string]struct{})
	for _, e := range entries {
		for _, t := range e.Tags {
			seen[t] = struct{}{}
		}
	}
	tags := make([]string, 0, len(seen))
	for t := range seen {
		tags = append(tags, t)
	}
	sort.Strings(tags)
	return tags
}
