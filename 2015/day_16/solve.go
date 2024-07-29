package day16

import (
	"regexp"
	"strconv"
	"strings"
)

func SolveOne(toSeach []string, input []string) (int, error) {
	return solve(toSeach, input, func(_ string, i int) func(int) bool {
		return func(j int) bool { return j == i }
	})
}

func SolveTwo(toSearch []string, input []string) (int, error) {
	return solve(toSearch, input, func(n string, i int) func(int) bool {
		switch n {
		case "cat", "trees":
			return func(j int) bool { return j > i }
		case "pomeranians", "goldfish":
			return func(j int) bool { return j < i }
		default:
			return func(j int) bool { return j == i }
		}
	})
}

var re = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

func solve(toSeach []string, input []string, dictator func(string, int) func(int) bool) (int, error) {
	search, err := praseToSearch(toSeach, dictator)
	if err != nil {
		return 0, err
	}
main:
	for i, inst := range input {
		rowData := re.FindAllStringSubmatch(inst, -1)
		for j := 2; j < len(rowData[0])-1; j += 2 {
			pred, ok := search[rowData[0][j]]
			if !ok {
				continue main
			}
			num, err := strconv.Atoi(rowData[0][j+1])
			if err != nil {
				return 0, err
			}
			if !pred(num) {
				continue main
			}
		}
		return i + 1, nil
	}
	return -1, nil
}

func praseToSearch(data []string, dictator func(string, int) func(int) bool) (map[string]func(int) bool, error) {
	output := make(map[string]func(int) bool)
	for _, inst := range data {
		pair := strings.Split(inst, ": ")
		num, err := strconv.Atoi(pair[1])
		if err != nil {
			return nil, err
		}
		output[pair[0]] = dictator(pair[0], num)
	}
	return output, nil
}
