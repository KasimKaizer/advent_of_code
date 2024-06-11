package day09

import (
	"strconv"
	"strings"
)

// formula:
// g(i, S) = min(C(i,k) + g(k, S - {k}))
//      k is a subset of S
//
// we first calculate the last leg then ie (k, 1), then move on upwards.
// use bitwise operations to store the value, as to avoid duplication
// mem 2d array would have an size of [n][1<<n]
// where first [] would be pos and second [] would be mask
// mask is basically  a binary number indicating nodes we have visited
// if we haven't visited any nodes then mask would be 00000, and if
// we have visited all the cities then it would be (1<<n)-1 or 11111
//
// solution inspired from here: https://github.com/alexchao26/advent-of-code-go/blob/main/2015/day09/main.go
// and here: https://github.com/KasimKaizer/coding_quest/blob/main/problem_25/shopping.go

// SolveOne solves the first part of the problem.
func SolveOne(input []string) (int, error) {
	return solve(input, func(x, y int) int { return min(x, y) })
}

// SolveTwo solves the second part of the problem.
func SolveTwo(input []string) (int, error) {
	return solve(input, func(x, y int) int { return max(x, y) })
}

type graph map[string]map[string]int

func newGraph(input []string) (graph, error) {
	g := make(graph)
	for _, inst := range input {
		data := strings.Fields(inst)

		if g[data[0]] == nil {
			g[data[0]] = make(map[string]int)
		}

		if g[data[2]] == nil {
			g[data[2]] = make(map[string]int)
		}
		num, err := strconv.Atoi(data[4])
		if err != nil {
			return nil, err
		}
		g[data[0]][data[2]] = num
		g[data[2]][data[0]] = num
	}
	return g, nil
}

type config struct {
	graph graph
	lkup  map[string]int
	mem   [][]int
	ops   func(int, int) int
}

func newConfig(input []string, ops func(int, int) int) (*config, error) {
	g, err := newGraph(input)
	if err != nil {
		return nil, err
	}
	lkup := make(map[string]int)
	idx := 0
	for key := range g {
		lkup[key] = idx
		idx++
	}
	mem := make([][]int, len(g))
	for i := range mem {
		mem[i] = make([]int, 1<<len(g))
	}
	return &config{graph: g, lkup: lkup, mem: mem, ops: ops}, nil
}

func (c *config) travel(pos string, mask int) int {
	if mask == (1<<len(c.graph))-1 {
		return 0
	}
	if c.mem[c.lkup[pos]][mask] != 0 {
		return c.mem[c.lkup[pos]][mask]
	}

	var res int
	for key := range c.graph {
		if mask&(1<<c.lkup[key]) != 0 {
			continue
		}
		dist := c.travel(key, mask|(1<<c.lkup[key])) + c.graph[pos][key]
		if res == 0 {
			res = dist
			continue
		}
		res = c.ops(res, dist)
	}
	c.mem[c.lkup[pos]][mask] = res
	return res
}

func solve(input []string, ops func(int, int) int) (int, error) {
	c, err := newConfig(input, ops)
	if err != nil {
		return 0, err
	}
	var res int
	for key := range c.graph {
		dist := c.travel(key, 0|(1<<c.lkup[key]))
		if res == 0 {
			res = dist
			continue
		}
		res = c.ops(res, dist)
	}
	return res, nil
}
