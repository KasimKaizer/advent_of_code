package day05

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type linked map[int]map[int]struct{}

func SolveOne(input []string) (int, error) {
	l, m, err := parseInput(input)
	if err != nil {
		return 0, nil
	}
	var count int
	sortRows(l, m, func(shift, midPoint int) {
		if shift == 0 {
			count += midPoint
		}
	})
	return count, nil
}

func SolveTwo(input []string) (int, error) {
	l, m, err := parseInput(input)
	if err != nil {
		return 0, nil
	}
	var count int
	sortRows(l, m, func(shift, midPoint int) {
		if shift != 0 {
			count += midPoint
		}
	})
	return count, nil
}

func parseInput(in []string) (linked, [][]int, error) {
	// split the input pased on the empty line
	// parse the first part or the lookup part as the mapofmap, where the second value is the first map value
	// parse the second part of the input as a matrix
	sep := slices.Index(in, "")
	lkup, err := createlkup(in[:sep])
	if err != nil {
		return nil, nil, err
	}
	matrix, err := createMatrix(in[sep+1:])
	if err != nil {
		return nil, nil, err
	}
	return lkup, matrix, nil
}

// create the map to map
func createlkup(in []string) (linked, error) {
	lkup := make(linked)
	for _, l := range in {
		splitL := strings.Split(l, "|")
		num1, err := strconv.Atoi(splitL[0])
		num2, err2 := strconv.Atoi(splitL[1])
		if err != nil || err2 != nil {
			return nil, *cmp.Or(&err, &err2)
		}
		if _, ok := lkup[num2]; !ok {
			lkup[num2] = make(map[int]struct{})
		}
		lkup[num2][num1] = struct{}{}
	}
	return lkup, nil
}

func createMatrix(in []string) ([][]int, error) {
	var matrix [][]int
	for _, r := range in {
		sRow := strings.Split(r, ",")
		var vals []int
		for _, numChar := range sRow {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			vals = append(vals, num)
		}
		matrix = append(matrix, vals)
	}
	return matrix, nil
}

// use the swap strat, where we swap the two numbers if the one is smaller then the other.
// go through each row of the matrix and see if there is any swaps required.
// i could probably use the in built algo to do this, but wanna try to implement something myself.
func sortRows(lkup linked, matx [][]int, f func(shift, midPoint int)) {
	for _, r := range matx {
		var shift int
	rLoop:
		for i := 0; i < len(r); {
			for j := i + 1; j < len(r); j++ {
				if _, ok := lkup[r[i]][r[j]]; ok {
					r[i], r[j] = r[j], r[i]
					shift++
					continue rLoop
				}
			}
			i++
		}
		f(shift, r[len(r)/2])
	}
}
