package day18

type coord struct {
	x, y int
}

func SolveOne(input []string, times int) (int, error) {
	return solve(parseInput(input), times, func(_ coord) bool { return true })
}

func SolveTwo(input []string, times int) (int, error) {
	grid := parseInput(input)
	lastX, lastY := len(grid[0])-1, len(grid)-1
	grid[0][0] = '#'         // top left corner
	grid[0][lastX] = '#'     // top right corner
	grid[lastY][0] = '#'     // bottom left corner
	grid[lastY][lastX] = '#' // bottom right corner
	return solve(grid, times, func(c coord) bool {
		if (c.y == 0 || c.y == lastY) &&
			(c.x == 0 || c.x == lastX) {
			return false
		}
		return true
	})
}

func solve(grid [][]byte, times int, ops func(coord) bool) (int, error) {
	var count int
	for i := range times {
		newGrid := make([][]byte, len(grid))
		for y := range len(grid) {
			for x := range len(grid[y]) {
				var s byte = '#'
				c := coord{x, y}
				if ops(c) {
					s = lookAround(c, grid)
				}
				newGrid[y] = append(newGrid[y], s)
				if i == times-1 && s == '#' {
					count++
				}
			}
		}
		grid = newGrid
	}
	return count, nil
}

var dirs = []coord{
	{0, -1},  // top
	{0, 1},   // bottom
	{1, 0},   // right
	{-1, 0},  // left
	{1, -1},  // top right
	{-1, -1}, // top left
	{1, 1},   // bottom right
	{-1, 1},  // bottom left
}

func lookAround(c coord, grid [][]byte) byte {
	var on int
	for _, dir := range dirs {
		x, y := (c.x + dir.x), (c.y + dir.y)
		if y >= len(grid) || y < 0 ||
			x >= len(grid[y]) || x < 0 ||
			grid[y][x] == '.' {
			continue
		}
		on++
	}
	if (grid[c.y][c.x] == '#' && (on == 2 || on == 3)) ||
		(grid[c.y][c.x] == '.' && on == 3) {
		return '#'
	}
	return '.'
}

func parseInput(input []string) [][]byte {
	var output [][]byte
	for _, row := range input {
		output = append(output, []byte(row))
	}
	return output
}
