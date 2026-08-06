package ledger

import (
	"errors"
	"testing"
)

func sample() []Entry {
	return []Entry{
		{Account: "cash", Amount: 500, Memo: " opening ", Tags: []string{"open"}},
		{Account: "fees", Amount: -120, Memo: "wire", Tags: []string{"bank", "open"}},
		{Account: "cash", Amount: -75, Memo: "coffee"},
	}
}

func TestBalance(t *testing.T) {
	b := NewBook()
	b.PostAll(sample())

	got, err := b.Balance("cash")
	if err != nil {
		t.Fatalf("Balance(cash): %v", err)
	}
	if got != 425 {
		t.Errorf("cash = %d, want 425", got)
	}
	if _, err := b.Balance("nope"); !errors.Is(err, ErrUnknownAccount) {
		t.Errorf("Balance(nope) err = %v, want ErrUnknownAccount", err)
	}
}

func TestTotal(t *testing.T) {
	b := NewBook()
	b.PostAll(sample())
	if got := b.Total(); got != 305 {
		t.Errorf("Total() = %d, want 305", got)
	}
}

func TestSnapshot(t *testing.T) {
	b := NewBook()
	b.PostAll(sample())

	snap := b.Snapshot()
	if len(snap) != 2 {
		t.Fatalf("len = %d, want 2", len(snap))
	}
	if snap[0].ID != "cash" || snap[1].ID != "fees" {
		t.Errorf("order = %q, %q", snap[0].ID, snap[1].ID)
	}
	if snap[0].Postings != 2 || snap[0].LastMemo != "coffee" {
		t.Errorf("cash = %+v", snap[0])
	}
}

func TestNormalize(t *testing.T) {
	entries := sample()
	Normalize(entries)
	if entries[0].Memo != "opening" {
		t.Errorf("memo = %q, want %q", entries[0].Memo, "opening")
	}
}

func TestSplit(t *testing.T) {
	credits, debits := Split(sample())
	if len(credits) != 1 || len(debits) != 2 {
		t.Fatalf("split = %d credits, %d debits; want 1, 2", len(credits), len(debits))
	}
	if credits[0].Account != "cash" || debits[0].Account != "fees" {
		t.Errorf("credits[0]=%q debits[0]=%q", credits[0].Account, debits[0].Account)
	}
}

func TestFilter(t *testing.T) {
	got := Filter(sample(), "open")
	if len(got) != 2 {
		t.Fatalf("Filter(open) = %d entries, want 2", len(got))
	}
	if len(Filter(sample(), "missing")) != 0 {
		t.Errorf("Filter(missing) should be empty")
	}
}

func TestDistinctTags(t *testing.T) {
	tags := DistinctTags(sample())
	want := []string{"bank", "open"}
	if len(tags) != len(want) {
		t.Fatalf("tags = %v, want %v", tags, want)
	}
	for i := range want {
		if tags[i] != want[i] {
			t.Fatalf("tags = %v, want %v", tags, want)
		}
	}
}

func TestTop(t *testing.T) {
	b := NewBook()
	b.PostAll(sample())
	top := b.Top(1)
	if len(top) != 1 || top[0].ID != "cash" {
		t.Fatalf("Top(1) = %v", top)
	}
	if got := b.Top(10); len(got) != 2 {
		t.Errorf("Top(10) = %d accounts, want 2", len(got))
	}
}

func TestZeroBook(t *testing.T) {
	var b Book
	b.Post(Entry{Account: "cash", Amount: 10})
	if got, err := b.Balance("cash"); err != nil || got != 10 {
		t.Errorf("Balance(cash) = %d, %v; want 10, nil", got, err)
	}
}

func TestAccountString(t *testing.T) {
	a := Account{ID: "cash", Balance: 425, Postings: 2}
	if got, want := a.String(), "cash 425 (2 postings)"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
