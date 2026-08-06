package manifest

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

const good = `# deployment
name = api
replicas = 3

REGION = us-east-1
LOG_LEVEL = debug
`

func TestParse(t *testing.T) {
	m, err := Parse(strings.NewReader(good))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if m.Name != "api" || m.Replicas != 3 {
		t.Errorf("got name=%q replicas=%d", m.Name, m.Replicas)
	}
	if m.Env["REGION"] != "us-east-1" || m.Env["LOG_LEVEL"] != "debug" {
		t.Errorf("env = %v", m.Env)
	}
	if len(m.Env) != 2 {
		t.Errorf("len(env) = %d, want 2", len(m.Env))
	}
}

func TestParseDefaultReplicas(t *testing.T) {
	m, err := Parse(strings.NewReader("name=api\n"))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if m.Replicas != 1 {
		t.Errorf("Replicas = %d, want 1", m.Replicas)
	}
}

func TestParseErrors(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want error
	}{
		{"no separator", "name=api\njust-a-word\n", ErrMalformed},
		{"empty key", "name=api\n = value\n", ErrMalformed},
		{"duplicate env key", "name=api\nA=1\nA=2\n", ErrDuplicateKey},
		{"duplicate reserved key", "name=api\nname=web\n", ErrDuplicateKey},
		{"no name", "A=1\n", ErrNoName},
		{"zero replicas", "name=api\nreplicas=0\n", ErrBadReplicas},
		{"negative replicas", "name=api\nreplicas=-3\n", ErrBadReplicas},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := Parse(strings.NewReader(tc.in)); !errors.Is(err, tc.want) {
				t.Errorf("err = %v, want %v", err, tc.want)
			}
		})
	}

	_, err := Parse(strings.NewReader("name=api\nreplicas=many\n"))
	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Errorf("replicas err = %v, want a *strconv.NumError", err)
	}
}

func TestParseLongValue(t *testing.T) {
	value := strings.Repeat("x", 200_000)
	m, err := Parse(strings.NewReader("name=api\nBLOB=" + value + "\n"))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if m.Env["BLOB"] != value {
		t.Errorf("BLOB length = %d, want %d", len(m.Env["BLOB"]), len(value))
	}
}

func TestLoadAll(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a.mf")
	if err := os.WriteFile(path, []byte(good), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := LoadAll([]string{path, filepath.Join(dir, "gone.mf")})
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}
	if len(got) != 1 || got[0].Name != "api" {
		t.Fatalf("LoadAll = %v", got)
	}

	bad := filepath.Join(dir, "bad.mf")
	if err := os.WriteFile(bad, []byte("oops\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadAll([]string{bad}); !errors.Is(err, ErrMalformed) {
		t.Errorf("LoadAll(bad) err = %v, want ErrMalformed", err)
	}
}

func TestMerge(t *testing.T) {
	base, err := Parse(strings.NewReader(good))
	if err != nil {
		t.Fatal(err)
	}
	overlay := &Manifest{Replicas: 5, Env: map[string]string{"LOG_LEVEL": "info", "EXTRA": "1"}}

	cp := base.Clone()
	cp.Merge(overlay)

	if cp.Name != "api" || cp.Replicas != 5 {
		t.Errorf("name=%q replicas=%d", cp.Name, cp.Replicas)
	}
	if cp.Env["LOG_LEVEL"] != "info" || cp.Env["EXTRA"] != "1" || cp.Env["REGION"] != "us-east-1" {
		t.Errorf("env = %v", cp.Env)
	}

	fresh := &Manifest{}
	fresh.Merge(overlay)
	if fresh.Env["EXTRA"] != "1" || fresh.Replicas != 5 {
		t.Errorf("fresh = %+v", fresh)
	}
}

func TestClone(t *testing.T) {
	base, err := Parse(strings.NewReader(good))
	if err != nil {
		t.Fatal(err)
	}
	cp := base.Clone()
	if cp.Name != base.Name || cp.Replicas != base.Replicas || len(cp.Env) != len(base.Env) {
		t.Errorf("clone = %+v, base = %+v", cp, base)
	}
	if got := (&Manifest{}).Clone(); got.Env != nil {
		t.Errorf("clone of an empty manifest has env %v", got.Env)
	}
}

func TestRender(t *testing.T) {
	m := &Manifest{Name: "api", Replicas: 2, Env: map[string]string{"Z": "1", "A": "2", "M": "3"}}
	var sb strings.Builder
	if err := m.Render(&sb); err != nil {
		t.Fatalf("Render: %v", err)
	}
	want := "name=api\nreplicas=2\nA=2\nM=3\nZ=1\n"
	if sb.String() != want {
		t.Errorf("Render() =\n%q\nwant\n%q", sb.String(), want)
	}

	round, err := Parse(strings.NewReader(sb.String()))
	if err != nil {
		t.Fatalf("re-parse: %v", err)
	}
	if round.Name != m.Name || round.Replicas != m.Replicas || len(round.Env) != len(m.Env) {
		t.Errorf("round trip = %+v, want %+v", round, m)
	}
}
