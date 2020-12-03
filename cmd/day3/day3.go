package day3

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

// Down represent the move length, downwards
const Down = 1

// Right represent the move length, to the right
const Right = 3

var (
	day3Cmd = &cobra.Command{
		Use:   "day3",
		Short: "Command used for the day3 challenge.",
		Long:  "Command used for the day3 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Command returns the command.
func Command() *cobra.Command {
	return day3Cmd
}

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFile("cmd/day3/input.txt")
	if err != nil {
		return
	}

	countTrees(reader)

	return
}

func countTrees(reader utils.Reader) {
	lines := reader.Lines
	trees := make([][]bool, len(lines))

	for i, s := range lines {
		trees[i] = make([]bool, len(s))
		for j, c := range s {
			trees[i][j] = (c == '#')
		}
	}

	// Part one
	fmt.Println(fmt.Sprintf("There are %d trees in part one.", checkSlope(1, 3, trees)))

	// Part two
	a := checkSlope(1, 1, trees)
	b := checkSlope(1, 3, trees)
	c := checkSlope(1, 5, trees)
	d := checkSlope(1, 7, trees)
	e := checkSlope(2, 1, trees)
	fmt.Printf("Part two: %d\n", a*b*c*d*e)
}

func checkSlope(down, right int, trees [][]bool) (count int) {
	count = 0
	for index := 0; index*down < len(trees); index++ {
		current := index * down
		col := (index * right) % len(trees[current])
		if trees[current][col] {
			count++
		}
	}
	return
}
