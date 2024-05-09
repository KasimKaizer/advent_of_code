// nolint: all
package day02_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2017/day_02"
)

type tests struct {
	Description string
	Input       string
	Expected    int
}

var testCasesOne = []tests{
	{
		"First example case",
		"base.txt",
		18,
	},

	{
		"Problem case",
		"input.txt",
		32121,
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		"base2.txt",
		9,
	},

	{
		"Problem case",
		"input.txt",
		197,
	},
}

func parseInput(path string) ([][]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var output [][]int
	for scanner.Scan() {
		var row []int
		splitRow := strings.Fields(scanner.Text())
		for _, numChar := range splitRow {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		output = append(output, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func runTests(t *testing.T, ops func([][]int) int, funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			input, err := parseInput(tc.Input)
			if err != nil {
				t.Error(err)
			}
			got := ops(input)
			if got != tc.Expected {
				t.Fatalf("%s(%v) = %v, expected: %v", funcName, tc.Input, got, tc.Expected)
			}
		})
	}
}

func TestSolveOne(t *testing.T) {
	runTests(t, SolveOne, "SolveOne", testCasesOne)
}

func TestSolveTwo(t *testing.T) {
	runTests(t, SolveTwo, "SolveTwo", testCasesTwo)
}

func runBenchmark(b *testing.B, ops func([][]int) int, test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		b.Run(tc.Description, func(b *testing.B) {
			input, err := parseInput(tc.Input)
			if err != nil {
				b.Error(err)
			}
			for i := 0; i < b.N; i++ {
				ops(input)
			}
		})
	}
}

func BenchmarkSolveOne(b *testing.B) {
	runBenchmark(b, SolveOne, testCasesOne)
}

func BenchmarkSolveTwo(b *testing.B) {
	runBenchmark(b, SolveTwo, testCasesTwo)
}
