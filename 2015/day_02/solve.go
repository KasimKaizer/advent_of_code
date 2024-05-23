package day02

import (
	"math"
	"strconv"
	"strings"
)

func SolveOne(input []string) (int, error) {
	return solve(input, calculateArea)
}

func SolveTwo(input []string) (int, error) {
	return solve(input, calculateRibbon)
}

func solve(input []string, ops func(string) (int, error)) (int, error) {
	total := 0
	for _, spec := range input {
		area, err := ops(spec)
		if err != nil {
			return 0, err
		}
		total += area
	}
	return total, nil
}

func calculateArea(spec string) (int, error) {
	total, smallest := 0, math.MaxInt
	metrics, err := getMetrics(spec)
	if err != nil {
		return 0, err
	}
	for idx := range metrics {
		next := (idx + 1) % len(metrics)
		area := metrics[idx] * metrics[next]
		smallest = min(smallest, area)
		total += area
	}
	return (2 * total) + smallest, nil
}

func calculateRibbon(spec string) (int, error) {
	metrics, err := getMetrics(spec)
	if err != nil {
		return 0, err
	}
	multi, sum, biggest := 1, 0, 0
	for _, num := range metrics {
		biggest = max(biggest, num)
		sum += num
		multi *= num
	}
	return ((sum - biggest) * 2) + multi, nil
}

func getMetrics(spec string) ([]int, error) {
	metrics := make([]int, 3)
	for idx, numChar := range strings.Split(spec, "x") {
		num, err := strconv.Atoi(numChar)
		if err != nil {
			return nil, err
		}
		metrics[idx] = num
	}
	return metrics, nil
}
