package day15

import (
	"regexp"
	"strconv"
)

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func SolveOne(input []string) (int, error) {
	ings, err := parse(input)
	if err != nil {
		return 0, err
	}
	return cook(100, ingredient{}, ings, func(_ ingredient) bool { return true }), nil
}

func SolveTwo(input []string) (int, error) {
	ings, err := parse(input)
	if err != nil {
		return 0, err
	}
	return cook(100, ingredient{}, ings, func(i ingredient) bool { return i.calories == 500 }), nil
}

func cook(tSpoon int, i ingredient, ings []ingredient, condition func(ingredient) bool) int {
	if tSpoon == 0 {
		if !condition(i) || i.capacity <= 0 || i.durability <= 0 || i.flavor <= 0 || i.texture <= 0 {
			return 0
		}
		return (i.capacity * i.durability * i.flavor * i.texture)
	}
	if len(ings) == 0 {
		return 0
	}
	curI := ings[len(ings)-1]
	ings = ings[:len(ings)-1]
	var maxVal int
	for idx := range tSpoon + 1 {
		maxVal = max(maxVal,
			cook(
				tSpoon-idx,
				ingredient{
					i.capacity + (curI.capacity * idx),
					i.durability + (curI.durability * idx),
					i.flavor + (curI.flavor * idx),
					i.texture + (curI.texture * idx),
					i.calories + (curI.calories * idx),
				},
				ings,
				condition),
		)
	}
	return maxVal
}

var re = regexp.MustCompile(`(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func parse(input []string) ([]ingredient, error) {
	var ings []ingredient
	for _, inst := range input {
		var ing ingredient
		var err error
		data := re.FindAllStringSubmatch(inst, -1)
		ing.capacity, err = strconv.Atoi(data[0][2])
		if err != nil {
			return nil, err
		}
		ing.durability, err = strconv.Atoi(data[0][3])
		if err != nil {
			return nil, err
		}
		ing.flavor, err = strconv.Atoi(data[0][4])
		if err != nil {
			return nil, err
		}
		ing.texture, err = strconv.Atoi(data[0][5])
		if err != nil {
			return nil, err
		}
		ing.calories, err = strconv.Atoi(data[0][6])
		if err != nil {
			return nil, err
		}
		ings = append(ings, ing)
	}
	return ings, nil
}
