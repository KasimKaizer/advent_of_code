# Advent of Code Solutions in Go

This repository contains my solutions to the Advent of Code challenges implemented in Go. Advent of Code is an annual set of Christmas-themed programming puzzles that can be solved in any programming language.

## Project Structure

The repository is organized by year and day:

```text
advent_of_code/
├── 2015/
│   ├── day_01/
│   │   ├── input.txt
│   │   ├── solve.go
│   │   ├── solve_test.go
│   │   └── ...
│   ├── day02/
│   │   ├── input.txt
│   │   ├── solve.go
│   │   ├── solve_test.go
│   │   └── ...
│   └── ...
├── cmd
│  └── cli
│     ├── main.go
│     └── templates
├── go.mod
├── go.sum
├── pkg
│  └── parse
├── README.md
└── ...
```
- cmd/cli: Contains the CLI tool for managing Advent of Code tasks
- pkg: Common utilities and shared code
- YYYY: Year of the challenge
  - day_XX: Directory for each day's challenge
  - solve.go: Solution implementation
  - solve_test.go: Test cases for the solution
  - README.md: Problem specification
  - input.txt: Input data for the challenge

## CLI Tool

The CLI tool in the cmd/cli folder provides several useful features to streamline your Advent of Code workflow:

- *Create stub files*: Generates solve.go and solve_test.go files with basic structure for new challenges.
- *Download problem specifications*: Fetches the challenge description and saves it as a README.md file.
- *Download input data*: Retrieves the input data for the challenge and saves it as input.txt.
- *Extract base cases*: Automatically identifies and creates files for any example cases found in the problem description.

## Usage

To use the CLI tool, navigate to the cmd/cli directory and run:

```bash
go run main.go [command] [options]
```

For detailed usage instructions, use the --help flag.

## Running the Solutions

To run a solution for a specific day:

Navigate to the appropriate year and day directory
Run the solution file with Go:

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
