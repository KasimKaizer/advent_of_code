# Advent of Code Solutions in Go

This repository contains my solutions to the Advent of Code challenges implemented in Go.

The CLI that scaffolds days, downloads input, and submits answers lives in a separate repo: [KasimKaizer/aoc-cli](https://github.com/KasimKaizer/aoc-cli). Shared input helpers are imported from `github.com/KasimKaizer/aoc-cli/parse`.

## Project Structure

The repository is organized by year and day:

```text
advent_of_code/
├── 2015/
│   ├── day_01/
│   │   ├── solve.go
│   │   ├── solve_test.go
│   │   └── ...
│   └── ...
├── go.mod
├── go.sum
├── README.md
└── ...
```

- YYYY: Year of the challenge
  - day_XX: Directory for each day's challenge
    - solve.go: Solution implementation
    - solve_test.go: Test cases for the solution

## Scaffolding a new day

Install and run [aoc-cli](https://github.com/KasimKaizer/aoc-cli):

```bash
go install github.com/KasimKaizer/aoc-cli/cmd/aoc@latest
export AOC_PATH=/path/to/advent_of_code
export AOC_SESSION=your_session_cookie
aoc -y 2024 -d 7
```

## Running the Solutions

Navigate to the appropriate year and day directory, then:

For part one:

```bash
go test -v . -run TestSolveOne
```

For part two:

```bash
go test -v . -run TestSolveTwo
```

## About Advent of Code

Advent of Code is an annual coding event created by Eric Wastl. Each year, starting on December 1st, a new programming puzzle is released each day until December 25th. These puzzles can be solved in any programming language and are designed to be fun and educational. For more information, visit adventofcode.com.

## License

This project is open source and available under the MIT License.

## Acknowledgements

Thanks to the Advent of Code community for sharing ideas and approaches. Special thanks to Eric Wastl for creating and maintaining this fantastic coding challenge.
