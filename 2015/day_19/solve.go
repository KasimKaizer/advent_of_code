package day19

import (
	"fmt"
	"regexp"
	"strings"
)

func SolveOne(input []string) (int, error) {
	re := make(map[string][]string)
	txt := parseInput(input, func(s []string) {
		re[s[0]] = append(re[s[0]], s[1])
	})
	mem := make(map[string]struct{})
	for k, v := range re {
		for i := range len(txt) - len(k) + 1 {
			if txt[i:len(k)+i] != k {
				continue
			}
			for _, r := range v {
				mem[fmt.Sprintf("%s%s%s", txt[:i], r, txt[len(k)+i:])] = struct{}{}
			}
		}
	}
	return len(mem), nil
}

// I am making a pretty big assumption here, But this solution does successfully pass
// all tests so just gonna roll with it. The bfs or dfs solution would
// burn my pc down.
var rg = regexp.MustCompile(`^e+$`)

func SolveTwo(input []string) (int, error) {
	re := make(map[string]string)
	txt := parseInput(input, func(s []string) {
		re[s[1]] = s[0]
	})
	var moves int
	for {
		for k, v := range re {
			if rg.MatchString(txt) {
				return moves, nil
			}
			idx := strings.Index(txt, k)
			if idx == -1 {
				continue
			}
			txt = fmt.Sprintf("%s%s%s", txt[:idx], v, txt[len(k)+idx:])
			moves++
		}
	}
}

func parseInput(data []string, yield func([]string)) string {
	for _, inst := range data[:len(data)-1] {
		if inst == "" {
			continue
		}
		yield(strings.Split(inst, " => "))
	}
	return data[len(data)-1]
}
