package day23

import (
	"strconv"
	"strings"
)

var Ops = map[string]func(int) int{
	"hlf": func(i int) int { return i / 2 },
	"tpl": func(i int) int { return i * 3 },
	"inc": func(i int) int { return i + 1 },
	"jie": func(i int) int {
		if i%2 == 0 {
			return 1
		}
		return 0
	},
	"jio": func(i int) int {
		if i == 1 {
			return 1
		}
		return 0
	},
}

func solve(input []string, reg string, regs map[string]int) (int, error) {
	for i := 0; i < len(input); i++ {
		sInst := strings.Fields(input[i])
		char1 := strings.TrimRight(sInst[1], ",")
		var num1, num2 int
		var err error
		num1, err = strconv.Atoi(char1)
		if err != nil {
			num1 = regs[char1]
		}
		if len(sInst) > 2 {
			num2, err = strconv.Atoi(sInst[2])
			if err != nil {
				return 0, err
			}
		}
		switch sInst[0] {
		case "hlf", "tpl", "inc":
			regs[char1] = Ops[sInst[0]](num1)
		case "jmp":
			i += (num1 - 1)
		case "jie", "jio":
			if Ops[sInst[0]](num1) == 1 {
				i += (num2 - 1)
			}
		}
	}
	return regs[reg], nil
}

func SolveOne(input []string, reg string) (int, error) {
	return solve(input, reg, make(map[string]int))
}

func SolveTwo(input []string, reg string) (int, error) {
	return solve(input, reg, map[string]int{"a": 1})
}
