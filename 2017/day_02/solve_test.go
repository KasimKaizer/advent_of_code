// nolint: all
package day02_test

import (
	"testing"

	. "advent_of_code/2017/day_02"
)

type tests struct {
	Description string
	Input       any
	Expected    any
}

var testCasesOne = []tests{
	{
		"First example case",
		"", // add example input here.
		"", // add example expected here.
	},

	{
		"Problem case",
		"", // add actual test input here.
		"", // add actual test expected here.
	},
}

var testCasesTwo = []tests{
	{
		"First example case",
		"", // add example input here.
		"", // add example expected here.
	},

	{
		"Problem case",
		"", // add actual test input here.
		"", // add actual test  expected here.
	},
}

func runTests(t *testing.T, ops func(any) any, funcName string, tests []tests) {
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
