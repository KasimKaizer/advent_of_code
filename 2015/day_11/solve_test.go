// nolint: all
package day11_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_11"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       parse.Opener
	Expected    string
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewTextOpener("abcdefgh"), // add example input here.
		"abcdffaa",                      // add example expected here.
	},

	{
		"Second example case",
		parse.NewTextOpener("ghijklmn"), // add actual test input here.
		"ghjaabcc",                      // add actual test expected here.
	},

	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		"cqjxxyzz",                       // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.NewTextOpener("abcdefgh"), // add example input here.
		"abcdffaa",                      // add example expected here.
	},

	{
		"Second example case",
		parse.NewTextOpener("ghijklmn"), // add actual test input here.
		"ghjaabcc",                      // add actual test expected here.
	},

	{
		"Problem case",
		parse.NewTextOpener("cqjxxyzz"), // add actual test input here.
		"cqkaabcc",                      // add actual test expected here.
	},
}

func runTests(t *testing.T, ops func(string) string, funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			f, err := tc.Input.Open()
			if err != nil {
				t.Error(err)
			}
			defer f.Close()
			data, err := parse.ToString(f)
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

func runBenchmark(b *testing.B, ops func(string) string, test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		f, err := tc.Input.Open()
		if err != nil {
			b.Error(err)
		}
		defer f.Close()
		data, err := parse.ToString(f)
		if err != nil {
			b.Error(err)
		}
		b.Run(tc.Description, func(b *testing.B) {
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
