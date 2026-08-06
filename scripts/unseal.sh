#!/usr/bin/env bash
# Unseal a week's answer keys — a deliberate act, not a tab-complete.
# Legitimate reasons: /score, or the Sunday fresh-context verifier pass.
set -euo pipefail
week="${1:?usage: scripts/unseal.sh <week>   e.g. 2026-w32}"
dir="sealed/$week"
[ -d "$dir" ] || { echo "no such folder: $dir" >&2; exit 1; }

read -r -p "Unseal $dir? Only for /score or the Sunday verifier pass. [y/N] " confirm
case "$confirm" in
    y|Y|yes|YES) ;;
    *) echo "aborted." >&2; exit 1 ;;
esac

chmod -R u+rwX "$dir"

# Restore +x on files that had it before sealing (see seal.sh). `u+X` alone
# can't do this: it only grants +x where the bit already exists, and sealing
# zeroed every bit first.
manifest="sealed/.manifest/$week.exec"
if [ -f "$manifest" ]; then
    while IFS= read -r f; do
        [ -f "$f" ] && chmod u+x "$f"
    done < "$manifest"
fi

echo "$dir unsealed. Re-seal with scripts/seal.sh $week if scoring hasn't happened yet."
