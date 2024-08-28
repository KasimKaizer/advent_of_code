package day20

func SolveOne(input int) (int, error) {
	return calPresents(10, 3_000_000, input), nil
}

func SolveTwo(input int) (int, error) {
	return calPresents(11, 50, input), nil
}

func calPresents(multi, limit, want int) int {
	for door := 1; ; door++ {
		var total int
		for i := 1; i*i <= door; i++ {
			if door%i != 0 {
				continue
			}
			for _, num := range []int{i, door / i} {
				if door/num > limit {
					continue
				}
				total += num
				if num*num == door {
					break
				}
			}
		}
		if total*multi >= want {
			return door
		}
	}
}
