#!/usr/bin/env bash
# Seal a week's answer keys (A3): even a casual `cat` now fails.
# Run AFTER the Sunday commit — git cannot read sealed files.
set -euo pipefail
week="${1:?usage: scripts/seal.sh <week>   e.g. 2026-w32}"
dir="sealed/$week"
[ -d "$dir" ] || { echo "no such folder: $dir" >&2; exit 1; }
chmod -R a-rwx "$dir"
echo "$dir sealed (mode 000). Unseal only for /score or the Sunday verifier pass."
