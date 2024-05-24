package day05

import (
	"regexp"
)

func SolveOne(input []string) int {
	return solve(input, isValidOne)
}

func SolveTwo(input []string) int {
	return solve(input, isValidTwo)
}

func solve(input []string, valFunc func(string) bool) int {
	var count int
	for _, text := range input {
		if valFunc(text) {
			count++
		}
	}
	return count
}

var reBad = regexp.MustCompile(`ab|cd|pq|xy`)

func isValidOne(text string) bool {
	if reBad.MatchString(text) {
		return false
	}
	var vCount int
	var hasDouble bool
	for i := range len(text) {
		if i != len(text)-1 && text[i] == text[i+1] {
			hasDouble = true
		}
		switch text[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vCount++
		}
		if vCount >= 3 && hasDouble {
			return true
		}
	}
	return false
}

func isValidTwo(text string) bool {
	var hasDouble, hasPair bool
	mem := make(map[string]struct{})
	var last string
	for i := range len(text) - 1 {
		if i < len(text)-2 && text[i] == text[i+2] {
			hasDouble = true
		}
		_, ok := mem[text[i:i+2]]
		if ok && text[i:i+2] != last {
			hasPair = true
		}
		if hasDouble && hasPair {
			return true
		}
		last = text[i : i+2]
		mem[text[i:i+2]] = struct{}{}
	}
	return false
}
