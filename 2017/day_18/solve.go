package day18

import (
	"strconv"
	"sync"
	"time"
)

var ops = map[string]func(int, int) int{
	"add": func(i1, i2 int) int { return i1 + i2 },
	"mul": func(i1, i2 int) int { return i1 * i2 },
	"mod": func(i1, i2 int) int { return ((i1 % i2) + i2) % i2 },
	"set": func(_, i2 int) int { return i2 },
}

func SolveOne(input [][]string) (int, error) {
	var prev []int
	solve(0, input, func(m map[string]int, s []string, i int) bool {
		if s[0] == "snd" {
			prev = append(prev, i)
			m[s[1]] = 0
		}
		if s[0] == "rcv" && i != 0 {
			return true
		}
		return false
	})
	return prev[len(prev)-1], nil
}

func SolveTwo(input [][]string) (int, error) {
	ch1, ch2 := make(chan int, 100), make(chan int, 100)
	var count int
	var wg sync.WaitGroup
	wg.Add(2)
	go parallelRun(0, input, ch1, ch2, &wg, func() {})
	go parallelRun(1, input, ch2, ch1, &wg, func() { count++ })
	wg.Wait()
	return count, nil
}

func parallelRun(id int, input [][]string, snd chan<- int, rcv <-chan int, wg *sync.WaitGroup, f func()) {
	defer wg.Done()
	solve(id, input, func(m map[string]int, s []string, num int) bool {
		if s[0] == "snd" {
			snd <- num
			f()
		}
		if s[0] == "rcv" {
			select {
			case rcvNum := <-rcv:
				m[s[1]] = rcvNum
			case <-time.After(100 * time.Millisecond):
				return true
			}
		}
		return false
	})
}

func solve(id int, input [][]string, f func(map[string]int, []string, int) bool) {
	regs := map[string]int{"p": id}
	for i := 0; i < len(input); i++ {
		name := input[i][1]
		num1, err := strconv.Atoi(name)
		if err != nil {
			num1 = regs[name]
		}
		var num2 int
		if len(input[i]) > 2 {
			num2, err = strconv.Atoi(input[i][2])
			if err != nil {
				num2 = regs[input[i][2]]
			}
		}
		switch input[i][0] {
		case "jgz":
			if num1 > 0 {
				i += (num2 - 1)
			}
		case "snd", "rcv":
			if f(regs, input[i], num1) {
				return
			}
		default:
			regs[name] = ops[input[i][0]](num1, num2)
		}
	}
}

// Alternative solution for part 2.
//
// func parallelRun(id int, input [][]string, snd chan<- int, rcv <-chan int) int {
// 	var count int
// 	solve(id, input, func(m map[string]int, s []string, num int) bool {
// 		if s[0] == "snd" {
// 			snd <- num
// 			count++
// 		}
// 		if s[0] == "rcv" {
// 			select {
// 			case rcvNum := <-rcv:
// 				m[s[1]] = rcvNum
// 			case <-time.After(100 * time.Millisecond):
// 				return true
// 			}
// 		}
// 		return false
// 	})
// 	return count
// }
//
// func SolveTwo(input [][]string) (int, error) {
// 	ch1, ch2 := make(chan int, 100), make(chan int, 100)
// 	go parallelRun(0, input, ch1, ch2)
// 	return parallelRun(1, input, ch2, ch1), nil
// }
