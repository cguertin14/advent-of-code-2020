package day5

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

var (
	day5Cmd = &cobra.Command{
		Use:   "day5",
		Short: "Command used for the day5 challenge.",
		Long:  "Command used for the day5 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Command returns the command.
func Command() *cobra.Command {
	return day5Cmd
}

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFile("cmd/day5/input.txt")
	if err != nil {
		return
	}

	seats := parseSeats(reader)
	fmt.Printf("Part one: %d\n", partOne(seats))
	fmt.Printf("Part two: %d\n", partTwo(seats))

	return
}

func partOne(seats []Seat) (highest int) {
	for _, seat := range seats {
		if seat.ID() > highest {
			highest = seat.ID()
		}
	}
	return
}

func partTwo(seats []Seat) int {
	traversed := make([]bool, 1024) // 2**10
	for _, s := range seats {
		traversed[s.ID()] = true
	}

	init := true
	for i, found := range traversed {
		if init && found {
			init = false
			continue
		}
		if !init && !found {
			return i
		}
	}

	return 0
}
