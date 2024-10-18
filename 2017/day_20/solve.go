package day20

import (
	"errors"
	"regexp"
	"strconv"
)

type coord struct {
	x, y, z int
}

func (c *coord) add(c1 *coord) {
	c.x += c1.x
	c.y += c1.y
	c.z += c1.z
}

func (c *coord) mhDist() int {
	return max(c.x, -c.x) + max(c.y, -c.y) + max(c.z, -c.z)
}

type parti struct {
	pos, vel, asc *coord
	dead          bool
}

func (p *parti) travel() {
	p.vel.add(p.asc)
	p.pos.add(p.vel)
}

func SolveOne(input []string) (int, error) {
	partis, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	var curP int
	for i := range partis {
		if partis[i].asc.mhDist() > partis[curP].asc.mhDist() {
			continue
		}
		if partis[i].asc.mhDist() < partis[curP].asc.mhDist() {
			curP = i
			continue
		}
		if partis[i].vel.mhDist() > partis[curP].vel.mhDist() {
			continue
		}
		if partis[i].vel.mhDist() < partis[curP].vel.mhDist() ||
			partis[i].pos.mhDist() < partis[curP].pos.mhDist() {
			curP = i
		}
	}
	return curP, nil
}

func SolveTwo(input []string) (int, error) {
	partis, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	for range 200 {
		knownPos := make(map[coord][]int)
		for idx, p := range partis {
			p.travel()
			knownPos[*p.pos] = append(knownPos[*p.pos], idx)
		}
		for _, idxs := range knownPos {
			if len(idxs) < 2 {
				continue
			}
			for _, i := range idxs {
				partis[i].dead = true
			}
		}
	}
	var count int
	for _, p := range partis {
		if !p.dead {
			count++
		}
	}
	return count, nil
}

var re = regexp.MustCompile(`-?\d+`)

func parseInput(in []string) ([]*parti, error) {
	var p []*parti
	for _, inst := range in {
		valChars := re.FindAllString(inst, -1)
		var vals []int
		for _, c := range valChars {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			vals = append(vals, num)
		}
		if len(vals) < 9 {
			return nil, errors.New("parseInput: not enough values for a particle")
		}
		var co []coord
		for i := 0; i < len(vals); i += 3 {
			co = append(co, coord{x: vals[i], y: vals[i+1], z: vals[i+2]})
		}
		p = append(p, &parti{
			pos: &co[0],
			vel: &co[1],
			asc: &co[2],
		})
	}
	return p, nil
}
