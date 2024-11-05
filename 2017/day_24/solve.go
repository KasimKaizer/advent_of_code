package day24

import (
	"strconv"
	"strings"
)

type connection struct {
	leftSide, rightSide int
}

func (c *connection) sum() int {
	return c.leftSide + c.rightSide
}

func (c *connection) hasPin(pin int) (int, bool) {
	switch pin {
	case c.leftSide:
		return c.rightSide, true
	case c.rightSide:
		return c.leftSide, true
	}
	return 0, false
}

func SolveOne(input []string) (int, error) {
	cons, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	mem := make(map[int]int)
	buildBridge(0, 0, 0, cons, mem)
	var maxVal int
	for _, val := range mem {
		maxVal = max(maxVal, val)
	}
	return maxVal, nil
}

func SolveTwo(input []string) (int, error) {
	cons, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	mem := make(map[int]int)
	buildBridge(0, 0, 0, cons, mem)
	var maxVal, maxLen int
	for len, val := range mem {
		switch {
		case len > maxLen:
			maxVal, maxLen = val, len
		case maxLen == len:
			maxVal = max(maxVal, val)
		}
	}
	return maxVal, nil
}

func buildBridge(nextCon, leng, total int, cons []*connection, mem map[int]int) {
	for idx, curCon := range cons {
		next, ok := curCon.hasPin(nextCon)
		if !ok {
			continue
		}
		lastIdx := len(cons) - 1
		cons[lastIdx], cons[idx] = cons[idx], cons[lastIdx]
		buildBridge(next, leng+1, total+curCon.sum(), cons[:lastIdx], mem)
		cons[idx], cons[lastIdx] = cons[lastIdx], cons[idx]
	}
	mem[leng] = max(mem[leng], total)
}

func parseInput(input []string) ([]*connection, error) {
	var output []*connection
	for _, con := range input {
		sCon := strings.Split(strings.TrimSpace(con), "/")
		ls, err := strconv.Atoi(sCon[0])
		if err != nil {
			return nil, err
		}
		rs, err := strconv.Atoi(sCon[1])
		if err != nil {
			return nil, err
		}
		output = append(output, &connection{leftSide: ls, rightSide: rs})
	}
	return output, nil
}
