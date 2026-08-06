// Package manifest reads simple key=value deployment manifests.
//
// A manifest is one key=value pair per line. Blank lines and lines
// whose first non-space character is '#' are ignored. Two keys are
// reserved: "name" is required, and "replicas" must be a positive
// integer and defaults to 1. Every other key becomes an environment
// entry. No key may appear twice.
package manifest

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// maxLine bounds a single manifest line, well above the scanner's
// 64 KiB default so that long environment values parse.
const maxLine = 1 << 20

var (
	// ErrMalformed marks a line that is not a key=value pair.
	ErrMalformed = errors.New("manifest: malformed line")
	// ErrDuplicateKey marks a key defined twice in one file.
	ErrDuplicateKey = errors.New("manifest: duplicate key")
	// ErrNoName marks a manifest with no name key.
	ErrNoName = errors.New("manifest: missing name")
	// ErrBadReplicas marks a replicas value that is not a positive
	// integer.
	ErrBadReplicas = errors.New("manifest: replicas must be positive")
)

// Manifest is one parsed deployment description.
type Manifest struct {
	Name     string
	Replicas int
	Env      map[string]string
}

// Parse reads a manifest from r.
func Parse(r io.Reader) (*Manifest, error) {
	m := &Manifest{Replicas: 1, Env: make(map[string]string)}
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, bufio.MaxScanTokenSize), maxLine)
	seen := make(map[string]struct{})
	lineNo := 0
	for sc.Scan() {
		lineNo++
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			return nil, fmt.Errorf("line %d: %q: %w", lineNo, line, ErrMalformed)
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" {
			return nil, fmt.Errorf("line %d: %q: %w", lineNo, line, ErrMalformed)
		}
		if _, exists := seen[key]; exists {
			return nil, fmt.Errorf("line %d: %q: %w", lineNo, key, ErrDuplicateKey)
		}
		seen[key] = struct{}{}
		switch key {
		case "name":
			m.Name = value
		case "replicas":
			n, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("line %d: replicas %q: %w", lineNo, value, err)
			}
			if n < 1 {
				return nil, fmt.Errorf("line %d: replicas %d: %w", lineNo, n, ErrBadReplicas)
			}
			m.Replicas = n
		default:
			m.Env[key] = value
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("manifest: read: %w", err)
	}
	if m.Name == "" {
		return nil, ErrNoName
	}
	return m, nil
}

// Load parses the manifest at path.
func Load(path string) (m *Manifest, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("manifest: open %s: %w", path, err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			m, err = nil, fmt.Errorf("manifest: close %s: %w", path, cerr)
		}
	}()

	parsed, err := Parse(f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return parsed, nil
}

// LoadAll loads every path in order, skipping ones that do not exist.
func LoadAll(paths []string) ([]*Manifest, error) {
	loaded := make([]*Manifest, 0, len(paths))
	for _, p := range paths {
		m, err := Load(p)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, fmt.Errorf("load all: %w", err)
		}
		loaded = append(loaded, m)
	}
	return loaded, nil
}

// Merge layers the set fields of src on top of m. A zero Replicas or
// an empty Name in src means "unset" and leaves m's value in place.
func (m *Manifest) Merge(src *Manifest) {
	if src == nil {
		return
	}
	if src.Name != "" {
		m.Name = src.Name
	}
	if src.Replicas > 0 {
		m.Replicas = src.Replicas
	}
	if m.Env == nil {
		m.Env = make(map[string]string, len(src.Env))
	}
	for k, v := range src.Env {
		m.Env[k] = v
	}
}

// Clone returns a copy of m sharing no state with it.
func (m *Manifest) Clone() *Manifest {
	cp := &Manifest{Name: m.Name, Replicas: m.Replicas}
	if m.Env != nil {
		cp.Env = make(map[string]string, len(m.Env))
		for k, v := range m.Env {
			cp.Env[k] = v
		}
	}
	return cp
}

// Render writes the manifest to w, reserved keys first and the
// environment keys after them in sorted order.
func (m *Manifest) Render(w io.Writer) error {
	bw := bufio.NewWriter(w)
	fmt.Fprintf(bw, "name=%s\n", m.Name)
	fmt.Fprintf(bw, "replicas=%d\n", m.Replicas)
	keys := make([]string, 0, len(m.Env))
	for k := range m.Env {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(bw, "%s=%s\n", k, m.Env[k])
	}
	if err := bw.Flush(); err != nil {
		return fmt.Errorf("manifest: render: %w", err)
	}
	return nil
}
