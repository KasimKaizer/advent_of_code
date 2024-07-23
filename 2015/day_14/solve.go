package day14

import (
	"regexp"
	"strconv"
)

type reindeer struct {
	name   string
	speed  int
	time   int
	rest   int
	status int
}

func SolveOne(input []string, sec int) (int, error) {
	var maxDist int
	err := solve(input, sec, func(m map[*reindeer]int) {
		for _, v := range m {
			maxDist = max(maxDist, v)
		}
	})
	return maxDist, err
}

func SolveTwo(input []string, sec int) (int, error) {
	score := make(map[*reindeer]int)
	var maxScore int
	err := solve(input, sec, func(m map[*reindeer]int) {
		var highest []*reindeer
		var maxDist int
		for k, v := range m {
			switch {
			case v > maxDist:
				maxDist = v
				highest = []*reindeer{k}
			case v == maxDist:
				highest = append(highest, k)
			}
		}
		for _, r := range highest {
			score[r]++
			maxScore = max(maxScore, score[r])
		}
	})
	return maxScore, err
}

func solve(input []string, sec int, yeild func(map[*reindeer]int)) error {
	reindeers := make(map[*reindeer]int)
	for _, inst := range input {
		rein, err := newReindeer(inst)
		if err != nil {
			return err
		}
		reindeers[rein] = 0
	}
	for range sec {
		for rein := range reindeers {
			reindeers[rein] += rein.move()
		}
		yeild(reindeers)
	}
	return nil
}

var re = regexp.MustCompile(`(\w+) can fly (\d+) km\/s for (\d+) seconds, but then must rest for (\d+) seconds\.`)

func newReindeer(inst string) (*reindeer, error) {
	data := re.FindAllStringSubmatch(inst, -1)
	name := data[0][1]
	speed, err := strconv.Atoi(data[0][2])
	if err != nil {
		return nil, err
	}
	time, err := strconv.Atoi(data[0][3])
	if err != nil {
		return nil, err
	}
	rest, err := strconv.Atoi(data[0][4])
	if err != nil {
		return nil, err
	}
	return &reindeer{
		name:   name,
		speed:  speed,
		time:   time,
		rest:   rest,
		status: 0,
	}, nil
}

func (r *reindeer) move() int {
	if r.status < 0 {
		r.status++
		return 0
	}
	r.status++
	if r.status == r.time {
		r.status = -r.rest
	}
	return r.speed
}
