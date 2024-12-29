package day04

type direction struct {
	row, col int
}

func SolveOne(input []string) (int, error) {
	return solve(input, 'X', findCount)
}

func SolveTwo(input []string) (int, error) {
	return solve(input, 'A', pairUp)
}

func solve(input []string, start byte, ops func([]string, direction) int) (int, error) {
	var total int
	for row := range len(input) {
		for col := range len(input[row]) {
			if input[row][col] != start {
				continue
			}
			total += ops(input, direction{row, col})
		}
	}
	return total, nil
}

var pair = map[byte]byte{
	'M': 'S',
	'S': 'M',
}

func pairUp(in []string, pos direction) int {
	// eliminate impossible
	if pos.row < 1 || pos.row >= len(in)-1 || pos.col < 1 || pos.col >= len(in[pos.row])-1 {
		return 0
	}

	if in[pos.row-1][pos.col-1] != pair[in[pos.row+1][pos.col+1]] || // check top left pairs with bottom right
		in[pos.row-1][pos.col+1] != pair[in[pos.row+1][pos.col-1]] { // check top right pairs with bottom left
		return 0
	}
	return 1
}

var dirs = []direction{
	{-1, -1}, // top left
	{0, -1},  // top
	{1, -1},  // top right
	{1, 0},   // right
	{1, 1},   // bottom right
	{0, 1},   // bottom
	{-1, 1},  // botttom left
	{-1, 0},  // left
}

var toFind = []byte{'X', 'M', 'A', 'S'}

func findCount(in []string, pos direction) int {
	var count int
	for _, d := range dirs {
		x, y, p := pos.row, pos.col, 0
		for p < len(toFind) && x >= 0 && y >= 0 &&
			x < len(in) && y < len(in[x]) &&
			in[x][y] == toFind[p] {
			x, y, p = (x + d.row), (y + d.col), (p + 1)
		}
		if p == len(toFind) {
			count++
		}
	}
	return count
}
