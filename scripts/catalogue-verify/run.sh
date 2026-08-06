#!/usr/bin/env bash
# Re-verifies every catalogue.md entry against the currently installed
# Go toolchain. Run this as the Sunday version-pin check (see
# catalogue.md header) whenever `go version` changes.
#
# Each entry is a self-contained main.go that prints one of:
#   BUG REPRODUCED   - defect still manifests as catalogue.md describes
#   STALE            - toolchain has neutralized this defect; catalogue.md
#                       needs a reword or removal, same as the loop-var
#                       precedent (Go 1.22) and the time.After precedent
#                       (Go 1.23)
#   INCONCLUSIVE      - runner-level issue (timing-sensitive race, etc.),
#                       needs a human look, not necessarily stale
set -uo pipefail
cd "$(dirname "$0")"

pass=0
stale=0
inconclusive=0

for d in */; do
  name="${d%/}"
  [ -f "$name/main.go" ] || continue
  out=$(cd "$name" && go run main.go 2>&1)
  if echo "$out" | grep -q "STALE"; then
    status="STALE"
    stale=$((stale+1))
  elif echo "$out" | grep -q "BUG REPRODUCED"; then
    status="PASS"
    pass=$((pass+1))
  else
    status="INCONCLUSIVE"
    inconclusive=$((inconclusive+1))
  fi
  printf "%-38s %s\n" "$name" "$status"
  if [ "$status" != "PASS" ]; then
    echo "$out" | sed 's/^/    /'
  fi
done

echo
echo "go version: $(go version)"
echo "pass=$pass stale=$stale inconclusive=$inconclusive"
if [ "$stale" -gt 0 ] || [ "$inconclusive" -gt 0 ]; then
  echo "ACTION NEEDED: update catalogue.md and its pinned Go version line for any STALE/INCONCLUSIVE entry above."
  exit 1
fi
