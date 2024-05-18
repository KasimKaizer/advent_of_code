// nolint: all
package day01_test

import (
	"io"
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_01"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       io.ReadCloser
	Expected    int
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.StrToReadCloser("(())"), // add example input here.
		0,                             // add example expected here.
	},
	{
		"Second example case",
		parse.StrToReadCloser("(()(()("), // add example input here.
		3,                                // add example expected here.
	},
	{
		"Problem case",
		parse.NoErrOpen("input.txt"), // add actual test input here.
		138,                          // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.StrToReadCloser(")"), // add example input here.
		1,                          // add example expected here.
	},
	{
		"Second example case",
		parse.StrToReadCloser("()())"), // add example input here.
		5,                              // add example expected here.
	},
	{
		"Problem case",
		parse.NoErrOpen("input.txt"), // add actual test input here.
		1771,                         // add actual test expected here.
	},
}

func runTests(t *testing.T, ops func(string) int, funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			defer tc.Input.Close()
			data, err := parse.ToString(tc.Input)
			if err != nil {
				t.Error(err)
			}
			got := ops(data)
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

func runBenchmark(b *testing.B, ops func(string) int, test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		b.Run(tc.Description, func(b *testing.B) {
			// defer tc.Input.Close()
			data, err := parse.ToString(tc.Input)
			if err != nil {
				b.Error(err)
			}
			for i := 0; i < b.N; i++ {
				ops(data)
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
