package day22

type coord struct {
	x, y int
}

type dir int

const (
	up dir = iota
	down
	left
	right
)

var dirMap = []coord{
	up:    {0, -1},
	down:  {0, 1},
	left:  {-1, 0},
	right: {1, 0},
}

func SolveOne(input []string) (int, error) {
	var infCount int
	solve(input, 10_000, func(grid map[coord]byte, c coord) dir {
		if status := grid[c]; status == '.' || status == 0 {
			grid[c] = '#'
			infCount++
			return left
		}
		grid[c] = '.'
		return right
	})
	return infCount, nil
}

func SolveTwo(input []string) (int, error) {
	var infCount int
	solve(input, 10_000_000, func(grid map[coord]byte, c coord) dir {
		switch grid[c] {
		case '.', 0:
			grid[c] = 'W'
			return left
		case 'W':
			grid[c] = '#'
			infCount++
			return up
		case '#':
			grid[c] = 'F'
			return right
		case 'F':
			grid[c] = '.'
			return down
		}
		return -1
	})
	return infCount, nil
}

func solve(input []string, itr int, f func(map[coord]byte, coord) dir) {
	var pos dir
	grid, curPos := parseInput(input)
	for range itr {
		pos = determinDir(pos, f(grid, curPos))
		curPos.x += dirMap[pos].x
		curPos.y += dirMap[pos].y
	}
}

func determinDir(prev, turn dir) dir {
	if turn == up {
		return prev
	}
	switch prev {
	case up:
		if turn == down {
			return down
		}
		return turn
	case down:
		if turn == down {
			return up
		}
		if turn == left {
			return right
		}
		return left
	case left:
		if turn == down {
			return right
		}
		if turn == left {
			return down
		}
		return up
	case right:
		if turn == down {
			return left
		}
		if turn == left {
			return up
		}
		return down
	}
	return -1
}

func parseInput(in []string) (map[coord]byte, coord) {
	grid := make(map[coord]byte)
	for y := range in {
		for x := range in[y] {
			grid[coord{x, y}] = in[y][x]
		}
	}
	return grid, coord{(len(in[0]) / 2), (len(in) / 2)}
}
