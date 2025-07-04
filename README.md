# DailyCodingProblems

A collection of daily coding challenges implemented in Go. 
Each problem is organized by day and includes a runnable solution, sometimes test cases, and optional benchmarking.

## Project Structure

```
DailyCodingProblems/
├── main.go               # CLI dispatcher: run any day's solution
├── scaffold.go           # Tool to generate new day folders
├── day01/
│   ├── run.go            # Entry point for Day 01
│   ├── solution.go       # Core logic
│   └── solution_test.go  # Unit tests
├── day02/
│   └── ...
```

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/DailyCodingProblems.git
cd DailyCodingProblems
go mod tidy
```

### 2. Run a specific day's solution

```bash
go run . --day=01
```

### 3. Scaffold a new day

```bash
go run scaffold.go --day=03
```

This creates a new folder `day03/` with boilerplate files and auto-registers it in `main.go`.

## Features

- Clean, idiomatic Go
- Recursive and iterative solutions
- Unit tests for every problem
- Performance benchmarks (time + memory)

## Learning Goals

- Deepen understanding of recursion, dynamic programming , and Go internals
- Practice writing testable, modular code
- Explore performance trade-offs (e.g. naive vs optimized recursion)
- Build tools to support long-term learning

## Contributing

This is a personal learning project, but feel free to fork, suggest improvements, or use the scaffold for your own daily practice.


## Inspired By

- [Daily Coding Problem](https://www.dailycodingproblem.com/)
- [LeetCode](https://leetcode.com/)
- [Go Proverbs](https://go-proverbs.github.io/)


## Author

Built with curiosity and Go by [@adjanour](https://github.com/adjanour)
