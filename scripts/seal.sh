#!/usr/bin/env bash
# Seal a week's answer keys (A3): even a casual `cat` now fails.
# Run AFTER the Sunday commit — git cannot read sealed files.
set -euo pipefail
week="${1:?usage: scripts/seal.sh <week>   e.g. 2026-w32}"
dir="sealed/$week"
[ -d "$dir" ] || { echo "no such folder: $dir" >&2; exit 1; }

# Record which files are executable before sealing. `unseal.sh` restores mode
# 000 with `chmod u+X`, which only grants +x where it already exists — since
# sealing zeroes every bit, that always resolves false and +x is lost forever
# without this manifest. Manifest lives outside $dir since it must be
# readable while the dir itself is mode 000.
manifest_dir="sealed/.manifest"
mkdir -p "$manifest_dir"
find "$dir" -type f -perm -u+x -print > "$manifest_dir/$week.exec"

# -depth: seal children before the folder itself. Plain `chmod -R` strips the
# folder's own traverse bit first and then cannot descend into it.
chmod u+rwx "$dir"
find "$dir" -depth -exec chmod a-rwx {} +
echo "$dir sealed (mode 000). Unseal only for /score or the Sunday verifier pass."
