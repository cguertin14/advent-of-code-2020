package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Reader is a struct to read stuff.
type Reader struct {
	Lines []string
}

// ToNumbers returns a list of numbers from the lines.
func (reader *Reader) ToNumbers() (numbers []int) {
	for _, i := range reader.Lines {
		if num, err := strconv.Atoi(i); err == nil {
			numbers = append(numbers, num)
		}
	}
	return
}

// ReadFromFile reads a text file and
// returns its individual lines.
func (reader *Reader) ReadFromFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reader.Lines = strings.Split(string(data), "\n")

	return reader.Lines, nil
}

// ReadFromLiteral reads a string and returns its numbers.
func (reader *Reader) ReadFromLiteral(literal string) []string {
	reader.Lines = strings.Split(literal, "\n")
	return reader.Lines
}
