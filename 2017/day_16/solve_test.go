// nolint: all
package day16_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2017/day_16"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Size        int
	Input       parse.Opener
	Expected    string
}

var testCasesOne = []tests{
	{
		"First example case",
		5,
		parse.NewTextOpener("s1,x3/4,pe/b"), // add example input here.
		"baedc",                             // add example expected here.
	},
	// {
	// 	"Problem case",
	// 	16,
	// 	parse.NewFileOpener("input.txt"), // add example input here.
	// 	"",                               // add example expected here.
	// },
}

var testCasesTwo = []tests{
	// {
	// 	"Problem case",
	// 	16,
	// 	parse.NewFileOpener("input.txt"), // add example input here.
	// 	"",               // add example expected here.
	// },
}

func runTests(t *testing.T, ops func(int, []string) (string, error), funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			f, err := tc.Input.Open()
			if err != nil {
				t.Error(err)
			}
			defer f.Close()
			data, err := parse.ToSplitString(f, ",")
			if err != nil {
				t.Error(err)
			}
			got, err := ops(tc.Size, data)
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

func runBenchmark(b *testing.B, ops func(int, []string) (string, error), test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		f, err := tc.Input.Open()
		if err != nil {
			b.Error(err)
		}
		data, err := parse.ToSplitString(f, ",")
		if err != nil {
			b.Error(err)
		}
		b.Run(tc.Description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ops(tc.Size, data)
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
