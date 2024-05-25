package day06

import (
	"strconv"
	"strings"
)

// Simple line sweep problem, Was waiting for a problem like this to show up!
// unlike normal line sweep problems, this is a grid so we essentially do line
// sweep twice.
// This might be a good time to create a data structure pkg, with indexed queue etc etc.
// This is how I will approach this problem:
// we get a (x,y) start and a (x,y) end. so,
// I will create a sorted queue (according to x), where start event = 1 and end event = 0.2
// Can't seem to figure out how to get line sweep/plane sweep to work in this case so,
// gonna do a brute force with maps,

type action int

const (
	onOff action = iota
	on
	off
)

type calculator interface {
	processPoints(action, *position, *position)
	calculateArea() int
}

type position struct {
	X, Y int
}

func SolveOne(input []string) (int, error) {
	return solve(input, newGridOne())
}

func SolveTwo(input []string) (int, error) {
	return solve(input, newGridTwo())
}

func solve(input []string, grid calculator) (int, error) {
	for _, instruct := range input {
		ops, start, end, err := parseInput(instruct)
		if err != nil {
			return 0, err
		}
		grid.processPoints(ops, start, end)
	}
	return grid.calculateArea(), nil
}

type gridOne struct {
	data map[position]struct{}
}

func newGridOne() *gridOne {
	return &gridOne{data: make(map[position]struct{})}
}

func (g *gridOne) processPoints(ops action, start *position, end *position) {
	for i := start.X; i <= end.X; i++ {
		for j := start.Y; j <= end.Y; j++ {
			switch ops {
			case on:
				g.data[position{X: i, Y: j}] = struct{}{}
			case off:
				delete(g.data, position{X: i, Y: j})
			case onOff:
				_, ok := g.data[position{X: i, Y: j}]
				if ok {
					delete(g.data, position{X: i, Y: j})
					continue
				}
				g.data[position{X: i, Y: j}] = struct{}{}
			}
		}
	}
}

func (g *gridOne) calculateArea() int {
	return len(g.data)
}

type gridTwo struct {
	data map[position]int
}

func newGridTwo() *gridTwo {
	return &gridTwo{data: make(map[position]int)}
}

func (g *gridTwo) processPoints(ops action, start *position, end *position) {
	for i := start.X; i <= end.X; i++ {
		for j := start.Y; j <= end.Y; j++ {
			switch ops {
			case on:
				g.data[position{X: i, Y: j}]++
			case off:
				num, ok := g.data[position{X: i, Y: j}]
				if !ok || num <= 1 {
					delete(g.data, position{X: i, Y: j})
					continue
				}
				g.data[position{X: i, Y: j}]--
			case onOff:
				g.data[position{X: i, Y: j}] += 2
			}
		}
	}
}

func (g *gridTwo) calculateArea() int {
	var total int
	for _, num := range g.data {
		total += num
	}
	return total
}

func parseInput(instruct string) (action, *position, *position, error) {
	var ac action
	switch {
	case strings.HasPrefix(instruct, "turn on"):
		instruct = strings.TrimPrefix(instruct, "turn on")
		ac = on
	case strings.HasPrefix(instruct, "turn off"):
		instruct = strings.TrimPrefix(instruct, "turn off")
		ac = off
	default:
		instruct = strings.TrimPrefix(instruct, "toggle")
		ac = onOff
	}
	var pos [2][2]int
	for i, values := range strings.Split(instruct, "through") {
		for j, numChar := range strings.Split(strings.TrimSpace(values), ",") {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return 0, nil, nil, err
			}
			pos[i][j] = num
		}
	}
	return ac, &position{X: pos[0][0], Y: pos[0][1]}, &position{X: pos[1][0], Y: pos[1][1]}, nil
}
