// nolint: all
package day03_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2017/day_03"
)

type tests struct {
	Description string
	Input       int
	Expected    int
}

var testCasesOne = []tests{
	{
		"First example case",
		1, // add example input here.
		0, // add example expected here.
	},
	{
		"Second example case",
		12, // add example input here.
		3,  // add example expected here.
	},

	{
		"Third example case",
		23, // add example input here.
		2,  // add example expected here.
	},

	{
		"fourth example case",
		1024, // add example input here.
		31,   // add example expected here.
	},

	{
		"Problem case",
		325489, // add actual test input here.
		552,    // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		2, // add example input here.
		4, // add example expected here.
	},
	{
		"Second example case",
		11, // add example input here.
		23, // add example expected here.
	},

	{
		"Third example case",
		880, // add example input here.
		931, // add example expected here.
	},
	{
		"Problem case",
		325489, // add actual test input here.
		330785, // add actual test expected here.
	},
}

func runTests(t *testing.T, ops func(int) int, funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			got := ops(tc.Input)
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

func runBenchmark(b *testing.B, ops func(int) int, test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		b.Run(tc.Description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ops(tc.Input)
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
