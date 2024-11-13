package day25

import (
	"regexp"
	"strconv"
)

func SolveOne(input string) (int, error) {
	rowNum, colNum, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	num := 20151125 // starting number
	for i := 2; ; i++ {
		for row, col := i, 1; col <= i; row, col = row-1, col+1 {
			num *= 252533
			num %= 33554393
			if row == rowNum && col == colNum {
				return num, nil
			}
		}
	}
}

func SolveTwo(_ string) (int, error) {
	panic("This challenge doesn't have a part 2")
}

var re = regexp.MustCompile(`To continue, please consult the code grid in the manual.  Enter the code at row (\d+), column (\d+)\.`)

func parseInput(in string) (int, int, error) {
	mtchs := re.FindAllStringSubmatch(in, -1)
	rowNum, err := strconv.Atoi(mtchs[0][1])
	if err != nil {
		return 0, 0, err
	}
	colNum, err := strconv.Atoi(mtchs[0][2])
	return rowNum, colNum, err
}
