package day19

import (
	"errors"
	"strings"
)

type coord struct {
	x, y int
}

var dirs = []coord{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // left
	{-1, 0}, // down
}

func SolveOne(input []string) (any, error) {
	txt, _ := solve(input)
	return txt, nil
}

func SolveTwo(input []string) (any, error) {
	_, count := solve(input)
	return count, nil
}

func solve(input []string) (string, int) {
	loc, err := findStart(input)
	if err != nil {
		return "", 0
	}
	dir := coord{x: 0, y: 1}
	var out strings.Builder
	count := 1
	for {
		x, y := loc.x+dir.x, loc.y+dir.y
		if x < 0 || y < 0 || y >= len(input) ||
			x >= len(input[y]) || input[y][x] == ' ' {
			var flag bool
			dir, flag = findNext(input, loc, dir)
			if !flag {
				break
			}
			continue
		}
		if input[y][x] >= 'A' && input[y][x] <= 'Z' {
			out.WriteByte(input[y][x])
		}
		count++
		loc.x, loc.y = x, y
	}
	return out.String(), count
}

func findStart(in []string) (coord, error) {
	for i, char := range in[0] {
		if char != ' ' {
			return coord{x: i, y: 0}, nil
		}
	}
	return coord{0, 0}, errors.New("findStart: no starting point found")
}

func findNext(in []string, c coord, last coord) (coord, bool) {
	for _, dir := range dirs {
		x, y := dir.x+c.x, dir.y+c.y
		if y < 0 || x < 0 || y >= len(in) || x >= len(in[y]) ||
			dir == opp(last) || in[y][x] == ' ' {
			continue
		}
		return dir, true
	}
	return coord{}, false
}

func opp(c coord) coord {
	return coord{x: c.x * -1, y: c.y * -1}
}
