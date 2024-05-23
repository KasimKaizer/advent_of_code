package day03

type coord struct {
	x, y int
}

var dirs = [...]coord{ //nolint:gochecknoglobals // its fine to have this be global.
	'^': {0, 1},
	'v': {0, -1},
	'>': {1, 0},
	'<': {-1, 0},
}

func SolveOne(input string) int {
	return solve(input, coord{})
}

func SolveTwo(input string) int {
	return solve(input, coord{}, coord{})
}

func solve(input string, santa ...coord) int {
	grid := map[coord]struct{}{
		{0, 0}: {},
	}
	i := -1
	visited := 1
	for _, char := range []byte(input) {
		i = (i + 1) % len(santa)
		santa[i].x, santa[i].y = (santa[i].x + dirs[char].x), (santa[i].y + dirs[char].y)
		_, ok := grid[santa[i]]
		if ok {
			continue
		}
		grid[santa[i]] = struct{}{}
		visited++
	}
	return visited
}
