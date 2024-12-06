package day01

import (
	"slices"
)

func SolveOne(input [][]int) (int, error) {
	l, r := parseInput(input)
	var accum int
	for i := range l {
		diff := (r[i] - l[i])
		accum += max(diff, -diff)
	}
	return accum, nil
}

func SolveTwo(input [][]int) (int, error) {
	l, r := parseInput(input)
	// reduce
	rgt := make(map[int]int)
	for _, num := range r {
		rgt[num]++
	}
	var accum int
	for _, num := range l {
		accum += (num * rgt[num])
	}
	return accum, nil
}

func parseInput(in [][]int) ([]int, []int) {
	var first, second []int
	for _, data := range in {
		first = append(first, data[0])
		second = append(second, data[1])
	}
	slices.Sort(first)
	slices.Sort(second)
	return first, second
}
