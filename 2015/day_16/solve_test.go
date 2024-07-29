// nolint: all
package day16_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_16"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	ToSearch    parse.Opener
	Input       parse.Opener
	Expected    int
}

var testCasesOne = []tests{
	{
		"Problem case",
		parse.NewFileOpener("base_1.txt"),
		parse.NewFileOpener("input.txt"), // add actual test input here.
		103,                              // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"Problem case",
		parse.NewFileOpener("base_1.txt"),
		parse.NewFileOpener("input.txt"), // add actual test input here.
		405,                              // add actual test expected here.
	},
}

func runTests(t *testing.T, ops func([]string, []string) (int, error), funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			f, err := tc.Input.Open()
			if err != nil {
				t.Error(err)
			}
			defer f.Close()
			fc, err := tc.ToSearch.Open()
			if err != nil {
				t.Error(err)
			}
			defer fc.Close()
			inputData, err := parse.ToStringSlice(f)
			if err != nil {
				t.Error(err)
			}
			searchData, err := parse.ToStringSlice(fc)
			if err != nil {
				t.Error(err)
			}
			got, err := ops(searchData, inputData)
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

func runBenchmark(b *testing.B, ops func([]string, []string) (int, error), test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		f, err := tc.Input.Open()
		if err != nil {
			b.Error(err)
		}
		defer f.Close()
		fc, err := tc.ToSearch.Open()
		if err != nil {
			b.Error(err)
		}
		defer fc.Close()
		inputData, err := parse.ToStringSlice(f)
		if err != nil {
			b.Error(err)
		}
		searchData, err := parse.ToStringSlice(fc)
		if err != nil {
			b.Error(err)
		}
		b.Run(tc.Description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ops(searchData, inputData)
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
