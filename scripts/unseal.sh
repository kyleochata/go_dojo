#!/usr/bin/env bash
# Unseal a week's answer keys — a deliberate act, not a tab-complete.
# Legitimate reasons: /score, or the Sunday fresh-context verifier pass.
set -euo pipefail
week="${1:?usage: scripts/unseal.sh <week>   e.g. 2026-w32}"
dir="sealed/$week"
[ -d "$dir" ] || { echo "no such folder: $dir" >&2; exit 1; }
chmod -R u+rwX "$dir"
echo "$dir unsealed. Re-seal with scripts/seal.sh $week if scoring hasn't happened yet."
