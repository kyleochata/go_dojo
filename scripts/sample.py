#!/usr/bin/env python3
"""Roll the hunt dice with real randomness (A3).

An LLM asked to "sample 0 (~20%) / 1 (~50%) / 2 (~30%)" is a biased
die. This script is not. It writes the outcome straight into sealed/
so it never appears in a session transcript; the /author session
reads it from there and never restates it.

Usage:
    python3 scripts/sample.py <week> [--variants {3,4,5}]

    --variants   fix the variant count (use 3 on flash-deck days);
                 default: uniform random 3-5.

Writes: sealed/<week>/sampling.json
Distributions (edit only at a gate review; audit realized
frequencies against these from post-reveal log data):
    defective variants per hunt : 0 (20%), 1 (50%), 2 (30%)
    defects per defective variant: 1 (85%), 2 (15%)
    defect placement            : low-salience (60%), point-of-interest (40%)
"""
import argparse
import json
import os
import random
import sys

def main() -> int:
    p = argparse.ArgumentParser()
    p.add_argument("week", help="ISO week folder name, e.g. 2026-w32")
    p.add_argument("--variants", type=int, choices=[3, 4, 5], default=None)
    a = p.parse_args()

    sr = random.SystemRandom()
    n = a.variants if a.variants else sr.choice([3, 4, 5])
    defective_count = min(sr.choices([0, 1, 2], weights=[20, 50, 30])[0], n)
    which = sorted(sr.sample(range(1, n + 1), defective_count))

    plan = []
    for v in which:
        ndef = sr.choices([1, 2], weights=[85, 15])[0]
        placements = [
            sr.choices(["low-salience", "point-of-interest"], weights=[60, 40])[0]
            for _ in range(ndef)
        ]
        plan.append({"variant": v, "defects": ndef, "placement": placements})

    out = {
        "week": a.week,
        "variants": n,
        "defective_count": defective_count,
        "defective": plan,
    }

    d = os.path.join("sealed", a.week)
    os.makedirs(d, exist_ok=True)
    path = os.path.join(d, "sampling.json")
    if os.path.exists(path):
        print(f"refusing to overwrite existing {path} — delete it "
              f"deliberately if you mean to re-roll", file=sys.stderr)
        return 1
    with open(path, "w") as f:
        json.dump(out, f, indent=2)

    # The outcome stays in the file, not the transcript.
    print(f"sampled -> {path}")
    return 0

if __name__ == "__main__":
    sys.exit(main())
