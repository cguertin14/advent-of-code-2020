package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
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

func (reader *Reader) read(path, separator string) ([]string, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return
		}
		reader.Lines = strings.Split(string(data), separator)
	}()

	wg.Wait()

	return reader.Lines, nil
}

// ReadFromFile reads a text file and
// returns its individual lines.
func (reader *Reader) ReadFromFile(path string) ([]string, error) {
	return reader.read(path, "\n")
}

// ReadFromLiteral reads a string and returns its content.
func (reader *Reader) ReadFromLiteral(literal string) []string {
	reader.Lines = strings.Split(literal, "\n")
	return reader.Lines
}

// ReadFromLiteralDouble reads a string separated by double lines
// and returns its content.
func (reader *Reader) ReadFromLiteralDouble(literal string) []string {
	reader.Lines = strings.Split(literal, "\n\n")
	return reader.Lines
}

// ReadFromFileDouble reads a file's content separating by "\n\n".
func (reader *Reader) ReadFromFileDouble(path string) ([]string, error) {
	return reader.read(path, "\n\n")
}
