package day02

func SolveOne(input [][]int) (int, error) {
	return solve(input, 0)
}

func SolveTwo(input [][]int) (int, error) {
	return solve(input, 1)
}

func solve(input [][]int, saves int) (int, error) {
	var res int
	for _, data := range input {
		if checkReports(data, saves) {
			res++
		}
	}
	return res, nil
}

func checkReports(in []int, saves int) bool {
	intSaves := saves
	low, high := 1, 3
	if in[0] < in[len(in)-1] {
		low, high = -3, -1
	}
mainLoop:
	for i := 0; i < len(in)-1; {
		for j := 1; saves >= 0; j++ {
			if i+j >= len(in) {
				break mainLoop
			}
			diff := in[i] - in[i+j]
			if diff >= low && diff <= high {
				i += j
				continue mainLoop
			}
			saves--
		}
		if i == 0 && intSaves > 0 {
			return checkReports(in[1:], intSaves-1)
		}
		return false
	}
	return true
}
