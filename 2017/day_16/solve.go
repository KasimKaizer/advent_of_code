package day16

import (
	"cmp"
	"strconv"
	"strings"
	"sync"
)

// keep track of a variable called "offset", I don't really plan to move all
// the elements from the back of the list to the front (its wasted compute)
// but just redefine where the "start" of the list would be.
// For the other two operations, its simple enough, first one just entails moving elements.
// also keep track of the actual position of the elements when they are moved. so i can save
// on the need for constant lookup for a "program's" position (as they are called in the problem)

func SolveOne(size int, input []string) (string, error) {
	cSet, cPos := createSets(size)
	offSet, err := dance(input, 0, cSet, cPos)
	return cSetToStr(cSet, offSet), err
}

func SolveTwo(size int, input []string) (string, error) {
	cSet, cPos := createSets(size)
	var dances []string
	var offset int
	var once sync.Once
	var first string
	for range 1_000_000_000 {
		var err error
		offset, err = dance(input, offset, cSet, cPos)
		if err != nil {
			return "", err
		}
		pattern := cSetToStr(cSet, offset)
		if pattern == first {
			return dances[999_999_999%len(dances)], nil
		}
		once.Do(func() { first = pattern })
		dances = append(dances, pattern)
	}
	return dances[999_999_999], nil
}

func dance(input []string, offset int, cSet []byte, cPos []int) (int, error) {
	for _, inst := range input {
		var err error
		switch inst[0] {
		case 's':
			offset, err = spin(inst[1:], offset, len(cSet))
		case 'x':
			err = swapPlaces(inst[1:], offset, cSet, cPos)
		case 'p':
			swapChars(inst[1:], cSet, cPos)
		}
		if err != nil {
			return 0, err
		}
	}
	return offset, nil
}

func spin(s string, offset, size int) (int, error) {
	num, err := strconv.Atoi(s)
	return (offset + (size - num)) % size, err
}

func swapPlaces(s string, offset int, cSet []byte, cPos []int) error {
	idxs := strings.Split(s, "/")
	idx1, err1 := strconv.Atoi(idxs[0])
	idx2, err2 := strconv.Atoi(idxs[1])
	if cmp.Or(err1, err2) != nil {
		return cmp.Or(err1, err2)
	}
	i, j := ((idx1 + offset) % len(cSet)), ((idx2 + offset) % len(cSet))
	cSet[i], cSet[j] = cSet[j], cSet[i]
	cPos[cSet[i]], cPos[cSet[j]] = i, j
	return nil
}

func swapChars(s string, cSet []byte, cPos []int) {
	chars := strings.Split(s, "/")
	char1, char2 := chars[0][0], chars[1][0]
	cSet[cPos[char1]], cSet[cPos[char2]] = cSet[cPos[char2]], cSet[cPos[char1]]
	cPos[char1], cPos[char2] = cPos[char2], cPos[char1]
}

func createSets(num int) ([]byte, []int) {
	charSet := make([]byte, num)
	charPos := make([]int, 128)
	for i := range num {
		charSet[i] = ('a' + byte(i))
		charPos[('a' + byte(i))] = i
	}
	return charSet, charPos
}

func cSetToStr(charSet []byte, offSet int) string {
	var output strings.Builder
	for i := range charSet {
		output.WriteByte(charSet[((i + offSet) % len(charSet))])
	}
	return output.String()
}
