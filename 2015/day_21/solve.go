package day21

import (
	"cmp"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/KasimKaizer/advent_of_code/pkg/parse"
)

type player struct {
	hitPoints int
	damage    int
	armor     int
}

type item struct {
	cost   int
	damage int
	armor  int
}

func SolveOne(input []string) (int, error) {
	boss, err := createBoss(input)
	if err != nil {
		return 0, err
	}
	minCost := math.MaxInt
	err = generatePlayerBuilds(func(p player, cost int) {
		if fight(p, boss) {
			minCost = min(minCost, cost)
		}
	})
	return minCost, err
}

func SolveTwo(input []string) (int, error) {
	boss, err := createBoss(input)
	if err != nil {
		return 0, err
	}
	var maxCost int
	err = generatePlayerBuilds(func(p player, cost int) {
		if !fight(p, boss) {
			maxCost = max(maxCost, cost)
		}
	})
	return maxCost, err
}

func generatePlayerBuilds(f func(player, int)) error {
	weapons, err := parseItems("base_1.txt")
	armors, err1 := parseItems("base_2.txt")
	rings, err2 := parseItems("base_3.txt")
	if *cmp.Or(&err, &err1, &err2) != nil {
		return *cmp.Or(&err, &err1, &err2)
	}
	var empty item
	armors = append(armors, empty)
	rings = append(rings, empty, empty)
	for _, w := range weapons {
		for _, a := range armors {
			for k, r1 := range rings {
				for l, r2 := range rings {
					if l == k {
						continue
					}
					f(createPlayer(w, a, r1, r2))
				}
			}
		}
	}
	return nil
}

func fight(p1, p2 player) bool {
	plyrs := []*player{&p1, &p2}
	for {
		for i := range 2 {
			atkr, defr := plyrs[i], plyrs[(i+1)%2]
			defr.hitPoints -= max(atkr.damage-defr.armor, 1)
			if defr.hitPoints <= 0 {
				return p1.hitPoints > 0
			}
		}
	}
}

var re = regexp.MustCompile(` \d+`)

func parseItems(filePath string) ([]item, error) {
	var items []item
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	data, err := parse.ToStringSlice(f)
	if err != nil {
		return nil, err
	}
	for _, d := range data[1:] {
		m := re.FindAllString(d, -1)
		cost, err := strconv.Atoi(strings.TrimSpace(m[0]))
		damg, err1 := strconv.Atoi(strings.TrimSpace(m[1]))
		armr, err2 := strconv.Atoi(strings.TrimSpace(m[2]))
		if *cmp.Or(&err, &err1, &err2) != nil {
			return nil, *cmp.Or(&err, &err1, &err2)
		}
		items = append(items, item{cost: cost, damage: damg, armor: armr})
	}
	return items, nil
}

func createPlayer(weapon, armor, ring1, ring2 item) (player, int) {
	return player{
		hitPoints: 100,
		damage:    weapon.damage + armor.damage + ring1.damage + ring2.damage,
		armor:     weapon.armor + armor.armor + ring1.armor + ring2.armor,
	}, weapon.cost + armor.cost + ring1.cost + ring2.cost
}

func createBoss(in []string) (player, error) {
	var roll []int
	for _, inst := range in {
		num, err := strconv.Atoi(strings.TrimSpace(strings.Split(inst, ":")[1]))
		if err != nil {
			return player{}, err
		}
		roll = append(roll, num)
	}
	return player{hitPoints: roll[0], damage: roll[1], armor: roll[2]}, nil
}
