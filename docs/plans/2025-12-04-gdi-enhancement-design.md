# GDI Enhancement Design

## Overview

Enhance `go-day-init` (gdi) to reduce friction when starting new AoC days by:
- Supporting flexible invocation from anywhere in the repo
- Auto-creating directory structure
- Fetching all example code blocks from the puzzle page

## CLI Interface

```
gdi              # in day folder: init + input (current behavior)
gdi init         # in day folder: just files
gdi input        # in day folder: just input

gdi 4            # from anywhere: create day04 for current year
gdi 2025 4       # from anywhere: create 2025/go/day04/
gdi 2025 4 init  # just create files, no downloads
```

## File Structure

After running `gdi 2025 4`:

```
2025/
├── go/
│   └── day04/
│       ├── main.go          # template with correct input path
│       ├── main_test.go     # template with first example pre-filled
│       └── examples.txt     # all code blocks from puzzle page
└── inputs/
    └── day04.txt            # downloaded puzzle input
```

## Examples File Format

```
--- Example 1 ---
first code block content here

--- Example 2 ---
second code block content here
```

## Implementation Details

### Repo Root Detection

Walk up from cwd looking for directory containing:
- A year folder matching pattern `20\d{2}`
- A `.git` directory

### Example Extraction

- Fetch puzzle page: `https://adventofcode.com/{year}/day/{day}`
- Parse HTML for all `<pre><code>` blocks within the puzzle description
- Write to `examples.txt` with delimiter format
- Pre-populate `main_test.go` with first example block

### Argument Parsing

| Args | Behavior |
|------|----------|
| (none) | Detect year/day from cwd, run init + input |
| `init` | Detect from cwd, files only |
| `input` | Detect from cwd, download only |
| `N` | Day N, infer year from current date, full setup |
| `YYYY N` | Explicit year and day, full setup |
| `YYYY N init` | Explicit year and day, files only |
| `YYYY N input` | Explicit year and day, download only |

### Environment Variables

- `AOC_SESSION`: Required for downloading input and fetching puzzle page
