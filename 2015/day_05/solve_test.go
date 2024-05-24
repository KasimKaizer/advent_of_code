// nolint: all
package day05_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_05"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       parse.Opener
	Expected    int
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewTextOpener("ugknbfddgicrmopn"), // add example input here.
		1,                                       // add example expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("aaa"), // add example input here.
		1,                          // add example expected here.
	},
	{
		"Third example case",
		parse.NewTextOpener("jchzalrnumimnmhp"), // add example input here.
		0,                                       // add example expected here.
	},

	{
		"Fourth example case",
		parse.NewTextOpener("haegwjzuvuyypxyu"), // add example input here.
		0,                                       // add example expected here.
	},
	{
		"Fifth example case",
		parse.NewTextOpener("dvszwmarrgswjxmb"), // add example input here.
		0,                                       // add example expected here.
	},
	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		258,                              // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		parse.NewTextOpener("qjhvhtzxzqqjkmpb"), // add example input here.
		1,                                       // add example expected here.
	},
	{
		"Second example case",
		parse.NewTextOpener("xxyxx"), // add example input here.
		1,                            // add example expected here.
	},

	{
		"Third example case",
		parse.NewTextOpener("uurcxstgmygtbstg"), // add example input here.
		0,                                       // add example expected here.
	},
	{
		"Fourth example case",
		parse.NewTextOpener("ieodomkazucvgmuy"), // add example input here.
		0,                                       // add example expected here.
	},
	{
		"Problem case",
		parse.NewFileOpener("input.txt"), // add actual test input here.
		53,                               // add actual test expected here.
	},
}

func runTests(t *testing.T, ops func([]string) int, funcName string, tests []tests) {
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

func runBenchmark(b *testing.B, ops func([]string) int, test []tests) {
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
			for range b.N {
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
