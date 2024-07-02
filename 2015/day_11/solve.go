package day11

func SolveOne(input string) string {
	return solve(input)
}

func SolveTwo(input string) string {
	return solve(input)
}

func solve(input string) string {
	for {
		input = genNextPass(input)

		if validatePass(input) {
			return input
		}
	}
}

func genNextPass(old string) string {
	oldSlc := []byte(old)
	for i := len(oldSlc) - 1; i > 0; i-- {
		if oldSlc[i] != 'z' {
			oldSlc[i]++
			return string(oldSlc)
		}
		oldSlc[i] = 'a'
	}
	return string(oldSlc)
}

func validatePass(pass string) bool {
	var incr bool
	var overlap int
	for i := 0; i < len(pass); i += 2 { // divide the string in chunk of 2 and process that.
		// check if the first element of the chunk is an illegal character.
		if pass[i] == 'i' || pass[i] == 'o' || pass[i] == 'l' {
			return false
		}

		// if there is only one element in the chunk, skip rest of the checks.
		// av cd xy (x) <- like this.
		if i == len(pass)-1 {
			continue
		}

		// check if the second element of the chunk is an illegal character.
		if pass[i+1] == 'i' || pass[i+1] == 'o' || pass[i+1] == 'l' {
			return false
		}

		if pass[i]+1 == pass[i+1] { // check is the two elements of the chunk are in series.
			// if yes then check if the element before the chunk or the element after the
			// chunk is in series.
			if (i < len(pass)-2 && pass[i]+2 == pass[i+2]) ||
				(i != 0 && pass[i-1] == pass[i]-1) {
				incr = true
				continue
			}
		}

		// check if the two elements of the chuck are the same, also make sure the element
		// just before the chuck is not same.
		// or
		// check if the second element of the chuck is equal to the element just after the chuck.
		if (pass[i] == pass[i+1] &&
			(i == 0 || pass[i-1] != pass[i])) ||
			(i < len(pass)-2 && pass[i+1] == pass[i+2]) {
			overlap++
		}
	}
	// return if there is an increasing sequence like this 'abc' or 'xyz'.
	// and
	// if there are at least 2 overlaps, i.e. two pairs of same characters, like 'aa', 'bb'.
	return incr && (overlap >= 2)
}
