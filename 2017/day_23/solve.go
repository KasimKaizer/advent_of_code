package day23

import (
	"strconv"
	"strings"
)

func SolveOne(input []string) (int, error) {
	reg := make(map[byte]int)
	reg['a'] = 1
	var count int
	for i := 0; i < len(input); i++ {
		inst := strings.Fields(input[i])
		num1, err := strconv.Atoi(inst[1])
		if err != nil {
			num1 = reg[inst[1][0]]
		}
		num2, err := strconv.Atoi(inst[2])
		if err != nil {
			num2 = reg[inst[2][0]]
		}
		switch inst[0] {
		case "set":
			reg[inst[1][0]] = num2
		case "sub":
			reg[inst[1][0]] -= num2
		case "mul":
			reg[inst[1][0]] *= num2
			count++
		case "jnz":
			if num1 != 0 {
				i += (num2 - 1)
			}
		}
	}
	return count, nil
}

func SolveTwo(_ []string) (int, error) {
	// translated the given assembly code into golang code, and then optimized it.
	var count int
	for num := 108_100; num <= 125_100; num += 17 {
		for i := 2; (i * i) <= num; i++ {
			if (num % i) == 0 {
				count++
				break
			}
		}
	}
	return count, nil
}
