package day10

import "strings"

// look and say seq is basically run length encoding, but with just integers.
// gonna still do it like run length because I don't know how big the number would get.
// could have used conway's constant but I didn't know about the look and say problem seq hand.

func SolveOne(input string, rep int) int {
	return solve(input, rep)
}

func SolveTwo(input string, rep int) int {
	return solve(input, rep)
}

func solve(input string, rep int) int {
	for range rep {
		input = lookAndSaySeq(input)
	}
	return len(input)
}

func lookAndSaySeq(input string) string {
	var output strings.Builder
	var count int
	for i := range len(input) { // input is always a string containing integers.
		if i != len(input)-1 && input[i] == input[i+1] {
			count++
			continue
		}
		count++
		// first byte represents the repetition and second byte represents the number
		// being repeated.
		output.Write([]byte{byte('0' + count), input[i]})
		count = 0
	}
	return output.String()
}
