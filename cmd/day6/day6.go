package day6

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

var (
	day6Cmd = &cobra.Command{
		Use:   "day6",
		Short: "Command used for the day6 challenge.",
		Long:  "Command used for the day6 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Command returns the command.
func Command() *cobra.Command {
	return day6Cmd
}

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFile("cmd/day6/input.txt")
	if err != nil {
		return
	}

	groups, groupSize := parseGroups(reader)
	fmt.Printf("Part one: %d\n", partOne(groups))
	fmt.Printf("Part two: %d\n", partTwo(groups, groupSize))

	return
}

func partOne(groups []Group) (sum int) {
	for _, group := range groups {
		sum += len(group)
	}
	return
}

func partTwo(groups []Group, groupSize []int) (sum int) {
	for n, group := range groups {
		sum += group.CountEveryone(n, groupSize)
	}
	return
}
