# /// script
# requires-python = ">=3.12"
# dependencies = ["numpy", "scipy"]
# ///

import re

import numpy as np
from scipy.optimize import linprog

total = 0


with open("../../inputs/day10.txt") as f:
    for i, line in enumerate(f):
        line = line.strip()
        if not line:
            continue

        buttons = re.findall(r"\(([^)]+)\)", line)
        buttons = [[int(x) for x in b.split(",")] for b in buttons]

        joltage = re.search(r"\{([^}]+)\}", line)
        joltage = [int(x) for x in joltage.group(1).split(",")]

        num_counters = len(joltage)
        num_buttons = len(buttons)

        a = np.zeros((num_counters, num_buttons))
        for btn_idx, btn in enumerate(buttons):
            for counter_idx in btn:
                a[counter_idx, btn_idx] = 1

        b = np.array(joltage)
        c = np.ones(num_buttons)

        result = linprog(
            c, A_eq=a, b_eq=b, bounds=(0, None), method="highs", integrality=1
        )

        if not result.success:
            print(f"Line {i}: FAILED - {result.message}")
        else:
            total += int(result.fun)

print(f"Total: {total}")
