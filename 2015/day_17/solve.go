package day17

import "math"

func SolveOne(input []int, max int) (int, error) {
	var comb int
	travel(input, max, 0, 0, func(_ int) {
		comb++
	})
	return comb, nil
}

func SolveTwo(input []int, max int) (int, error) {
	minCons := math.MaxInt
	combPerCon := make(map[int]int)
	travel(input, max, 0, 0, func(count int) {
		minCons = min(minCons, count)
		combPerCon[count]++
	})
	return combPerCon[minCons], nil
}

func travel(containers []int, left, idx, count int, ops func(int)) {
	if left == 0 {
		ops(count)
		return
	}

	if left < 0 || idx >= len(containers) {
		return
	}

	// use a container
	travel(containers, (left - containers[idx]), (idx + 1), count+1, ops)
	// don't use a container
	travel(containers, left, (idx + 1), count, ops)
}
