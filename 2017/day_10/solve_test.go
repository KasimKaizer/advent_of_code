// nolint: all
package day10_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2017/day_10"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       parse.Opener
	limit       int
	Expected    any
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewTextOpener("3,4,1,5"), // add example input here.
		5,
		12, // add example expected here.
	},
	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		256,                              // add actual test expected here.
		826,
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.NewTextOpener(""), // add example input here.
		256,
		"a2582a3a0e66e6e86e3812dcb672a272", // add example expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("AoC 2017"), // add example input here.
		256,
		"33efeb34ea91902bb2f59c9920caa6cd", // add example expected here.
	},
	{
		"Third example case",
		parse.NewTextOpener("1,2,3"), // add example input here.
		256,
		"3efbe78a8d82f29979031a4aa0b16a9d", // add example expected here.
	},
	{
		"Forth example case",
		parse.NewTextOpener("1,2,4"), // add example input here.
		256,
		"63960835bcdc130f0b66d7ff4f6a5a8e", // add example expected here.
	},
	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		256,                              // add actual test expected here.
		"d067d3f14d07e09c2e7308c3926605c4",
	},
}

func runTests(t *testing.T, ops func(string, int) (any, error), funcName string, tests []tests) {
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
			got, err := ops(data, tc.limit)
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

func runBenchmark(b *testing.B, ops func(string, int) (any, error), test []tests) {
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
				ops(data, tc.limit)
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
