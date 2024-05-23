package day04

import (
	"crypto/md5" //nolint:gosec // this is advent of code.
	"encoding/hex"
	"fmt"
	"strings"
)

func SolveOne(input string) int {
	return solve(input, "00000")
}

func SolveTwo(input string) int {
	return solve(input, "000000")
}

func solve(input, prefix string) int {
	num := 0
	for {
		hash := getMD5(fmt.Sprintf("%s%d", input, num))
		if strings.HasPrefix(hash, prefix) {
			break
		}
		num++
	}
	return num
}

func getMD5(input string) string {
	mHash := md5.Sum([]byte(input)) //nolint:gosec // this is advent of code.
	return hex.EncodeToString(mHash[:])
}
