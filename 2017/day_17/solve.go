package day17

// Gonna take the ever growing array approach, loop should run for 2017 times.
// Observation: inserting a new value after the last value of the buffer would grow the buffer,
// that means there is literary no way to replace the 0th value of the buffer, i.e '0'
// will always be the 0th value.

func SolveOne(input int) (int, error) {
	buf := []int{0}
	var idx int
	for i := range 2017 {
		idx = (idx + input) % len(buf)
		temp := append([]int{i + 1}, buf[idx+1:]...)
		buf = append(buf[:idx+1], temp...)
		idx++
	}
	return buf[(idx+1)%len(buf)], nil
}

func SolveTwo(input int) (int, error) {
	var idx, aftZero int
	leng := 1
	for i := range 50_000_000 {
		idx = ((idx + input) % leng) + 1
		if idx == 1 {
			aftZero = (i + 1)
		}
		leng++
	}
	return aftZero, nil
}
