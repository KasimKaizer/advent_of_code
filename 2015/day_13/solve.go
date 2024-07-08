package day13

import (
	"regexp"
	"strconv"
)

func SolveOne(input []string) (int, error) {
	c, err := newConfig(input)
	if err != nil {
		return 0, nil
	}

	return c.find(c.start, 1), nil
}

func SolveTwo(input []string) (int, error) {
	c, err := newConfig(input)
	if err != nil {
		return 0, nil
	}
	extraName := "agon"
	c.link[extraName] = make(map[string]int)
	c.mem[extraName] = make(map[int]int)
	for name := range c.lkp {
		c.link[extraName][name] = 0
		c.link[name][extraName] = 0
	}
	c.lkp[extraName] = len(c.lkp)
	return c.find(c.start, 1), nil
}

type config struct {
	lkp   map[string]int
	link  map[string]map[string]int
	mem   map[string]map[int]int
	start string
}

var re = regexp.MustCompile(`(\w+) would (lose|gain) (\d+) happiness units by sitting next to (\w+)\.`)

func newConfig(input []string) (*config, error) {
	lkp := make(map[string]int)
	link := make(map[string]map[string]int)
	mem := make(map[string]map[int]int)
	var start string
	var idx int
	for _, inst := range input {
		matches := re.FindAllStringSubmatch(inst, -1)
		// all matches would be at the 0 index, also the first match would be the whole string.
		num, err := strconv.Atoi(matches[0][3])
		if err != nil {
			return nil, err
		}
		if matches[0][2] == "lose" {
			num = -num
		}
		_, ok := link[matches[0][1]]
		if !ok {
			link[matches[0][1]] = make(map[string]int)
			mem[matches[0][1]] = make(map[int]int)
			lkp[matches[0][1]] = idx
			if start == "" {
				start = matches[0][1]
			}
			idx++
		}
		link[matches[0][1]][matches[0][4]] = num
	}
	return &config{lkp: lkp, link: link, mem: mem, start: start}, nil
}

func (c *config) find(name string, mask int) int {
	if mask == (1<<len(c.lkp))-1 {
		return c.link[c.start][name] + c.link[name][c.start]
	}
	if c.mem[name][mask] != 0 {
		return c.mem[name][mask]
	}
	var total int
	for subName, score := range c.link[name] {
		if mask&(1<<c.lkp[subName]) != 0 {
			continue
		}
		tScr := c.find(subName, mask|(1<<c.lkp[subName])) + score + c.link[subName][name]
		total = max(total, tScr)
	}
	c.mem[name][mask] = total
	return total
}
