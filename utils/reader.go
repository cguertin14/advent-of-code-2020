package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func strToArr(input string) (numbers []int) {
	for _, i := range strings.Split(input, "\n") {
		if num, err := strconv.Atoi(i); err == nil {
			numbers = append(numbers, num)
		}
	}

	return
}

// ReadFromFile reads a text file and returns its numbers.
func ReadFromFile(path string) (numbers []int, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return strToArr(string(data)), nil
}

// ReadFromLiteral reads a string and returns its numbers.
func ReadFromLiteral(literal string) (numbers []int) {
	return strToArr(literal)
}
