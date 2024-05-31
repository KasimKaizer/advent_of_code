package day08

func SolveOne(input []string) int {
	return solve(input, opsOne)
}

func SolveTwo(input []string) int {
	return solve(input, opsTwo)
}

func solve(input []string, ops func(string) (int, int)) int {
	var tCount, sCount int
	for _, str := range input {
		t, s := ops(str)
		tCount += t
		sCount += s
	}
	return (tCount - sCount)
}

func opsOne(str string) (int, int) {
	var total, accum int
	for i := range len(str) {
		if accum > 0 {
			accum--
			continue
		}
		total++
		if (str[i] != '\\') || (i == len(str)-1) {
			continue
		}
		switch str[i+1] {
		case '\\', '"':
			accum = 1
		case 'x':
			accum = 3
		}
	}
	return len(str), (total - 2)
}

func opsTwo(str string) (int, int) {
	var total int
	for i := range len(str) {
		total++
		switch str[i] {
		case '\\', '"':
			total++
		}
	}
	return (total + 2), len(str)
}
