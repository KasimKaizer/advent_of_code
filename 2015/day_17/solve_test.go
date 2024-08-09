// nolint: all
package day17_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_17"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       parse.Opener
	Max         int
	Expected    any
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewTextOpener("20\n15\n10\n5\n5"), // add example input here.
		25,
		4,
	},
	{
		"Main case",
		parse.NewFileOpener("input.txt"), // add example input here.
		150,
		654,
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.NewTextOpener("20\n15\n10\n5\n5"), // add example input here.
		25,
		3,
	},
	{
		"Main case",
		parse.NewFileOpener("input.txt"), // add example input here.
		150,
		57,
	},
}

func runTests(t *testing.T, ops func([]int, int) (int, error), funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			f, err := tc.Input.Open()
			if err != nil {
				t.Error(err)
			}
			defer f.Close()
			data, err := parse.ToIntSlice(f)
			if err != nil {
				t.Error(err)
			}
			got, err := ops(data, tc.Max)
			if err != nil {
				t.Error(err)
			}
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

func runBenchmark(b *testing.B, ops func([]int, int) (int, error), test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		f, err := tc.Input.Open()
		if err != nil {
			b.Error(err)
		}
		defer f.Close()
		data, err := parse.ToIntSlice(f)
		if err != nil {
			b.Error(err)
		}
		b.Run(tc.Description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ops(data, tc.Max)
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
