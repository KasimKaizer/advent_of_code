// nolint: all
package day10_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_10"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       parse.Opener
	Rep         int
	Expected    int
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewTextOpener("1"), // add example input here.
		1,
		2, // add example expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("21"), // add actual test input here.
		2,
		6, // add actual test expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("111221"), // add actual test input here.
		1,
		6, // add actual test expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("1"), // add actual test input here.
		5,
		6, // add actual test expected here.
	},
	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		40,
		492982, // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.NewTextOpener("1"), // add example input here.
		1,
		2, // add example expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("21"), // add actual test input here.
		2,
		6, // add actual test expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("111221"), // add actual test input here.
		1,
		6, // add actual test expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("1"), // add actual test input here.
		5,
		6, // add actual test expected here.
	},
	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		50,
		6989950, // add actual test expected here.
	},
}

func runTests(t *testing.T, ops func(string, int) int, funcName string, tests []tests) {
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
			got := ops(data, tc.Rep)
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

func runBenchmark(b *testing.B, ops func(string, int) int, test []tests) {
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
				ops(data, tc.Rep)
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
