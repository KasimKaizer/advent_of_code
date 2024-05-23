package parse

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ToIntMatrix(reader io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(reader)
	var output [][]int
	for scanner.Scan() {
		var row []int
		splitRow := strings.Fields(scanner.Text())
		for _, numChar := range splitRow {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		output = append(output, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func ToStringMatrix(reader io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(reader)
	var output [][]string
	for scanner.Scan() {
		output = append(output, strings.Fields(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func ToIntSlice(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	var output []int
	for scanner.Scan() {
		splitData := strings.Fields(scanner.Text())
		for _, numChar := range splitData {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			output = append(output, num)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func ToStringSlice(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	out := make([]string, 0)
	for scanner.Scan() {
		out = append(out, strings.TrimSuffix(scanner.Text(), "\n"))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func ToInt(reader io.Reader) (int, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(strings.TrimSpace(string(data)))
	return num, err
}

func ToString(reader io.Reader) (string, error) {
	data, err := io.ReadAll(reader)
	return strings.TrimSpace(string(data)), err
}

type Opener interface {
	Open() (io.ReadCloser, error)
}

type FileOpener struct {
	path string
}

func NewFileOpener(filePath string) *FileOpener {
	return &FileOpener{path: filePath}
}
func (f *FileOpener) Open() (io.ReadCloser, error) {
	return os.Open(f.path)
}

type TextOpener struct {
	text string
}

func NewTextOpener(text string) *TextOpener {
	return &TextOpener{text: text}
}
func (t *TextOpener) Open() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(t.text)), nil // error is always nil.
}
