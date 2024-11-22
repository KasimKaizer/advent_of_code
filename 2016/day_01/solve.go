package day01

import "strconv"

type dir struct {
	x, y int
}

type dirName int

const (
	up dirName = iota + 1
	down
	right
	left
)

var nxtDir = map[byte]map[dirName]dirName{
	'L': {
		up:    left,
		down:  right,
		right: up,
		left:  down,
	},
	'R': {
		up:    right,
		down:  left,
		right: down,
		left:  up,
	},
}

var dirMap = [...]dir{
	up:    {0, 1},
	down:  {0, -1},
	right: {1, 0},
	left:  {-1, 0},
}

func solve(input []string, f func(dir) bool) (int, error) {
	cur, curDir := dir{0, 0}, up
mainloop:
	for _, inst := range input {
		num, err := strconv.Atoi(inst[1:])
		if err != nil {
			return 0, err
		}
		curDir = nxtDir[inst[0]][curDir]
		for range num {
			cur.x, cur.y = (cur.x + dirMap[curDir].x), (cur.y + dirMap[curDir].y)
			if f(cur) {
				break mainloop
			}
		}
	}
	return max(cur.x, -cur.x) + max(cur.y, -cur.y), nil
}

func SolveOne(input []string) (int, error) {
	return solve(input, func(_ dir) bool { return false })
}

func SolveTwo(input []string) (int, error) {
	grid := make(map[dir]struct{})
	return solve(input, func(d dir) bool {
		if _, ok := grid[d]; ok {
			return true
		}
		grid[d] = struct{}{}
		return false
	})
}
