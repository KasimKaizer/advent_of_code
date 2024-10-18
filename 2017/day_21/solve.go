package day21

import (
	"errors"
	"strings"
)

func SolveOne(input []string, count int) (int, error) {
	return solve(input, count)
}

func SolveTwo(input []string, count int) (int, error) {
	return solve(input, count)
}

func solve(input []string, itrNum int) (int, error) {
	ruleBook := parseInput(input)
	mainSqr := []string{".#.", "..#", "###"}
	for range itrNum {
		var newSqr [][]byte
		size := 2 + len(mainSqr)%2
		height, length := size, size
		for {
			temp, ok := ruleBook[getCurSqr(mainSqr, length, height, size)]
			if !ok {
				return 0, errors.New("pattern match not found")
			}
			sTemp := strings.Split(temp, "/")
			newSqrPos := (height / size) * len(sTemp)
			for i, j := (newSqrPos - len(sTemp)), 0; i < newSqrPos; i, j = i+1, j+1 {
				if len(newSqr) < newSqrPos {
					newSqr = append(newSqr, []byte(sTemp[j]))
					continue
				}
				newSqr[i] = append(newSqr[i], sTemp[j]...)
			}
			if length < len(mainSqr) {
				length += size
				continue
			}
			if height < len(mainSqr) {
				length = size
				height += size
				continue
			}
			break
		}
		mainSqr = toStrSlc(newSqr)
	}
	return countLights(mainSqr), nil
}

func countLights(srq []string) int {
	var count int
	for i := range srq {
		for _, c := range srq[i] {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func toStrSlc(in [][]byte) []string {
	out := make([]string, 0, len(in))
	for i := range in {
		out = append(out, string(in[i]))
	}
	return out
}

func getCurSqr(grid []string, length, height, size int) string {
	var cur strings.Builder
	for i := height - size; i < height; i++ {
		cur.WriteString(grid[i][length-size : length])
		if i != height-1 {
			cur.WriteByte('/')
		}
	}
	return cur.String()
}

func parseInput(in []string) map[string]string {
	out := make(map[string]string)
	for _, inst := range in {
		sInst := strings.Split(inst, " => ")
		base, val := sInst[0], sInst[1]
		for range 4 {
			out[base] = val
			out[flip(base)] = val
			base = rotate(base)
		}
	}
	return out
}

func flip(in string) string {
	sBase := strings.Split(in, "/")
	var fliped strings.Builder
	for i := range len(sBase) {
		for j := range len(sBase) {
			fliped.WriteByte(sBase[i][len(sBase)-j-1])
		}
		if i != len(sBase)-1 {
			fliped.WriteByte('/')
		}
	}
	return fliped.String()
}

func rotate(in string) string {
	sBase := strings.Split(in, "/")
	var fliped strings.Builder
	for i := range len(sBase) {
		for j := range len(sBase) {
			fliped.WriteByte(sBase[j][len(sBase)-i-1])
		}
		if i != len(sBase)-1 {
			fliped.WriteByte('/')
		}
	}
	return fliped.String()
}
