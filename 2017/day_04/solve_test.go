// nolint: all
package day04_test

import (
	"os"
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2017/day_04"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Input       string
	Expected    int
}

var testCasesOne = []tests{
	// {
	// 	"First example case",
	// 	"base_1.txt", // add example input here.
	//     "", // add example expected here.
	// },

	{
		"Problem case",
		"input.txt", // add actual test input here.
		386,         // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	// {
	// 	"First example case",
	// 	"base_1.txt", // add example input here.
	// 	0,            // add example expected here.
	// },

	{
		"Problem case",
		"input.txt", // add actual test input here.
		208,         // add actual test  expected here.
	},
}

func runTests(t *testing.T, ops func([]string) int, funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			f, err := os.Open(tc.Input)
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
		b.Run(tc.Description, func(b *testing.B) {
			f, err := os.Open(tc.Input)
			if err != nil {
				b.Error(err)
			}
			defer f.Close()
			data, err := parse.ToStringSlice(f)
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
