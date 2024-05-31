package day07

import (
	"math"
	"strconv"
	"strings"
)

// new solution, adapted from https://github.com/alexchao26/advent-of-code-go/blob/main/2015/day07/main.go

var opsFunc = map[string]func(int, int) int{ // nolint:gochecknoglobals // its fine
	"AND":    func(i1, i2 int) int { return i1 & i2 },
	"OR":     func(i1, i2 int) int { return i1 | i2 },
	"LSHIFT": func(i1, i2 int) int { return i1 << i2 },
	"RSHIFT": func(i1, i2 int) int { return i1 >> i2 },
}

func SolveOne(input []string, wire string) int {
	wireMap := make(map[string]string)
	for _, conn := range input {
		splitConn := strings.Split(conn, " -> ")
		wireMap[splitConn[1]] = splitConn[0]
	}
	mem := make(map[string]int)
	return search(mem, wireMap, wire)
}

func SolveTwo(input []string, wire string) int {
	wireMap := make(map[string]string)
	for _, conn := range input {
		splitConn := strings.Split(conn, " -> ")
		wireMap[splitConn[1]] = splitConn[0]
	}
	mem := make(map[string]int)
	wireMap["b"] = strconv.Itoa(search(mem, wireMap, wire))
	clear(mem)
	return search(mem, wireMap, wire)
}
func search(mem map[string]int, wires map[string]string, toSearch string) int {
	num, ok := mem[toSearch]
	if ok {
		return num
	}
	num, err := strconv.Atoi(toSearch)
	if err == nil {
		return num
	}
	equ := strings.Fields(wires[toSearch])
	var volt int
	switch len(equ) {
	case 1:
		volt = search(mem, wires, equ[0])
	case 2:
		volt = math.MaxUint16 ^ search(mem, wires, equ[1])
	case 3:
		volt = opsFunc[equ[1]](search(mem, wires, equ[0]), search(mem, wires, equ[2]))
	}
	mem[toSearch] = volt
	return volt
}

// Old solution:
// var opsFunc = map[string]func(uint16, uint16) uint16{ // nolint:gochecknoglobals // its fine
// 	"AND":    func(i1, i2 uint16) uint16 { return i1 & i2 },
// 	"OR":     func(i1, i2 uint16) uint16 { return i1 | i2 },
// 	"NOT":    func(i1, _ uint16) uint16 { return ^i1 }, // second value is disregarded
// 	"LSHIFT": func(i1, i2 uint16) uint16 { return i1 << i2 },
// 	"RSHIFT": func(i1, i2 uint16) uint16 { return i1 >> i2 },
// }

// func SolveOne(input []string, wire string) int {
// 	return solve(input, wire)
// }

// // this is a very very very finicky solution, I am surprised it works.
// func SolveTwo(input []string, wire string) int {
// 	newB := SolveOne(input, wire)
// 	for i, inp := range input {
// 		if strings.HasSuffix(inp, "-> b") {
// 			input[i] = fmt.Sprintf("%d -> b", newB)
// 			break
// 		}
// 	}
// 	return solve(input, wire)
// }

// func solve(input []string, wire string) int {
// 	wires := make(map[string]uint16)
// 	for i := 0; i < len(input); i++ {
// 		data := strings.Split(input[i], " -> ")
// 		equ := strings.Fields(data[0])
// 		var x, y, idx uint16 // all these values are 0.
// 		var ok bool
// 		if len(equ) == 2 {
// 			idx = 1
// 		}
// 		temp, err := strconv.ParseUint(equ[idx], 10, 16)
// 		x = uint16(temp)
// 		if err != nil {
// 			x, ok = wires[equ[idx]]
// 			if !ok {
// 				input = append(input, input[i])
// 				continue
// 			}
// 		}
// 		if len(equ) == 1 {
// 			wires[data[1]] = x
// 			continue
// 		}
// 		if len(equ) == 3 {
// 			temp, err := strconv.ParseUint(equ[2], 10, 16) //nolint:govet // its fine.
// 			y = uint16(temp)
// 			if err != nil {
// 				y, ok = wires[equ[2]]
// 				if !ok {
// 					input = append(input, input[i])
// 					continue
// 				}
// 			}
// 		}
// 		wires[data[1]] = opsFunc[equ[idx^1]](x, y)
// 	}
// 	return int(wires[wire])
// }
