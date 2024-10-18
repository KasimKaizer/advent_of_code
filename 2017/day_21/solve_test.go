// nolint: all
package day21_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2017/day_21"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       parse.Opener
	Count       int
	Expected    int
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewFileOpener("base_6.txt"), // add example input here.
		2,
		12, // add example expected here.
	},

	// {
	// 	"Problem case",
	// 	parse.NewFileOpener("input.txt"), // add actual test input here.
	// 	5,
	// 	0, // add actual test expected here.
	// },
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.NewFileOpener("base_6.txt"), // add example input here.
		2,
		12, // add example expected here.
	},

	// {
	// 	"Problem case",
	// 	parse.NewFileOpener("input.txt"), // add actual test input here.
	// 	18,
	// 	0, // add actual test expected here.
	// },
}

func runTests(t *testing.T, ops func([]string, int) (int, error), funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			f, err := tc.Input.Open()
			if err != nil {
				t.Error(err)
			}
			defer f.Close()
			data, err := parse.ToStringSlice(f)
			if err != nil {
				t.Error(err)
			}
			got, err := ops(data, tc.Count)
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

func runBenchmark(b *testing.B, ops func([]string, int) (int, error), test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		f, err := tc.Input.Open()
		if err != nil {
			b.Error(err)
		}
		defer f.Close()
		data, err := parse.ToStringSlice(f)
		if err != nil {
			b.Error(err)
		}
		b.Run(tc.Description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ops(data, tc.Count)
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
