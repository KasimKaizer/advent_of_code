// nolint: all
package day22_test

import (
	"testing"

	. "github.com/KasimKaizer/advent_of_code/2015/day_22"
	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type tests struct {
	Description string
	Player      parse.Opener
	Enemy       parse.Opener
	Expected    any
}

var testCasesOne = []tests{
	{
		"First example case",
		parse.NewTextOpener("Hit Points: 10\nMana: 250"), // add example input here.
		parse.NewTextOpener("Hit Points: 14\nDamage: 8"), // add example input here.
		641, // add example expected here.
	},

	// {
	// 	"Problem case",
	// 	parse.NewTextOpener("Hit Points: 50\nMana: 500"), // add actual test input here.
	// 	parse.NewFileOpener("input.txt"),
	// 	0, // add actual test expected here.
	// },
}

var testCasesTwo = []tests{
	// {
	// 	"Problem case",
	// 	parse.NewTextOpener("Hit Points: 50\nMana: 500"), // add actual test input here.
	// 	parse.NewFileOpener("input.txt"),
	// 	0, // add actual test expected here.
	// },
}

func runTests(t *testing.T, ops func(string, string) (int, error), funcName string, tests []tests) {
	for _, tc := range tests {
		t.Run(tc.Description, func(t *testing.T) {
			p, err := tc.Player.Open()
			if err != nil {
				t.Error(err)
			}
			defer p.Close()
			e, err := tc.Enemy.Open()
			if err != nil {
				t.Error(err)
			}
			defer e.Close()
			player, err := parse.ToString(p)
			if err != nil {
				t.Error(err)
			}
			enemy, err := parse.ToString(e)
			if err != nil {
				t.Error(err)
			}
			got, err := ops(player, enemy)
			if err != nil {
				t.Error(err)
			}
			if got != tc.Expected {
				t.Fatalf("%s(%v, %v) = %v, expected: %v", funcName, tc.Player, tc.Enemy, got, tc.Expected)
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

func runBenchmark(b *testing.B, ops func(string, string) (int, error), test []tests) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, tc := range test {
		p, err := tc.Player.Open()
		if err != nil {
			b.Error(err)
		}
		defer p.Close()
		e, err := tc.Enemy.Open()
		if err != nil {
			b.Error(err)
		}
		defer e.Close()
		player, err := parse.ToString(p)
		if err != nil {
			b.Error(err)
		}
		enemy, err := parse.ToString(e)
		if err != nil {
			b.Error(err)
		}
		b.Run(tc.Description, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ops(player, enemy)
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
