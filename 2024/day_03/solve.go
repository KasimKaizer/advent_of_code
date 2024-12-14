package day03

import (
	"regexp"
	"strconv"
)

func SolveOne(input string) (int, error) {
	return calSumOfProd(input, `(mul)\((\d+),(\d+)\)`, func(_ string) bool { return true })
}

func SolveTwo(input string) (int, error) {
	rtnVal := true
	return calSumOfProd(input, `(mul|do|don't)\((\d+)?,?(\d+)?\)`, func(s string) bool {
		if s == "don't" {
			rtnVal = false
		}
		if s == "do" {
			rtnVal = true
		}
		return rtnVal
	})
}

func calSumOfProd(in string, regx string, predicate func(string) bool) (int, error) {
	re, err := regexp.Compile(regx)
	if err != nil {
		return 0, err
	}
	mths := re.FindAllStringSubmatch(in, -1)
	var total int
	for _, m := range mths {
		if !predicate(m[1]) || m[2] == "" || m[3] == "" {
			continue
		}
		num1, err := strconv.Atoi(m[2])
		if err != nil {
			return 0, err
		}
		num2, err := strconv.Atoi(m[3])
		if err != nil {
			return 0, err
		}
		total += (num1 * num2)
	}
	return total, nil
}
