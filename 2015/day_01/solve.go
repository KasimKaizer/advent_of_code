package day01

var bracketMap = [...]int{ //nolint:gochecknoglobals
	'(': 1,
	')': -1,
}

func SolveOne(input string) int {
	count := 0
	for idx := range len(input) {
		count += bracketMap[input[idx]]
	}
	return count
}

func SolveTwo(input string) int {
	count := 0
	for idx := range len(input) {
		count += bracketMap[input[idx]]
		if count == -1 {
			return (idx + 1)
		}
	}
	return -1
}
