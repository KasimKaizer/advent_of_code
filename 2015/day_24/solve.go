package day24

import (
	"math"
)

func SolveOne(input []int) (int, error) {
	return solve(input, 3), nil
}

func SolveTwo(input []int) (int, error) {
	return solve(input, 4), nil
}

func solve(input []int, divBy int) int {
	limit := sumSlc(input) / divBy
	numPkgs, smallQE := math.MaxInt32, 0
	sortPrsts(input, limit, 0, 0, 1, func(num, qe int) {
		if num > numPkgs || qe < 0 {
			return
		}
		if num == numPkgs {
			smallQE = min(smallQE, qe)
			return
		}
		numPkgs, smallQE = num, qe
	})
	return smallQE
}

func sumSlc(in []int) int {
	var rollCount int
	for _, num := range in {
		rollCount += num
	}
	return rollCount
}

func sortPrsts(presents []int, left, idx, count, qe int, ops func(int, int)) {
	if left == 0 {
		ops(count, qe)
		return
	}

	if left < 0 || idx >= len(presents) {
		return
	}

	// use a present
	sortPrsts(presents, (left - presents[idx]), (idx + 1), count+1, qe*presents[idx], ops)
	// don't use a present
	sortPrsts(presents, left, (idx + 1), count, qe, ops)
}
