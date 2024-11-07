package day25

import (
	"strconv"
	"strings"
)

type instruction struct {
	toWrite  int
	netDir   int
	nxtState string
}

type state struct {
	zero, one *instruction
}

func SolveOne(input []string) (int, error) {
	state, itr, sTbl, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	trnMcn := make(map[int]int)
	var idx int
	for range itr {
		inst := sTbl[state].zero
		if trnMcn[idx] == 1 {
			inst = sTbl[state].one
		}
		trnMcn[idx] = inst.toWrite
		idx += inst.netDir
		state = inst.nxtState
	}
	var total int
	for _, val := range trnMcn {
		if val == 1 {
			total++
		}
	}
	return total, nil
}

func SolveTwo(input []string) (int, error) {
	panic("This challange has no Part 2")
}

func parseInput(in []string) (string, int, map[string]state, error) {
	// get the start state.
	start := parseVal(in[0])
	// get the total number of runs.
	temp := strings.Fields(in[1])
	count, err := strconv.Atoi(temp[len(temp)-2])
	if err != nil {
		return "", 0, nil, err
	}
	// get the state map
	sTbl := parseState(in[3:])
	return start, count, sTbl, nil
}

func parseState(in []string) map[string]state {
	sTbl := make(map[string]state)
	for i := 0; i < len(in); i += 10 {
		name := parseVal(in[i])
		zero := parseInst(in[i+2 : i+2+3])
		one := parseInst(in[i+6 : i+6+3])
		sTbl[name] = state{zero: zero, one: one}
	}
	return sTbl
}

func parseInst(in []string) *instruction {
	writeVal := 0
	if parseVal(in[0]) == "1" {
		writeVal = 1
	}
	nxtDir := 1
	if parseVal(in[1]) == "left" {
		nxtDir = -1
	}
	nxtState := parseVal(in[2])
	return &instruction{toWrite: writeVal, netDir: nxtDir, nxtState: nxtState}
}

func parseVal(in string) string {
	temp := strings.Fields(in)
	return strings.TrimRight(temp[len(temp)-1], ".:")
}
