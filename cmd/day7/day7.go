package day7

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

var (
	day7Cmd = &cobra.Command{
		Use:   "day7",
		Short: "Command used for the day7 challenge.",
		Long:  "Command used for the day7 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Command returns the command.
func Command() *cobra.Command {
	return day7Cmd
}

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFile("cmd/day7/input.txt")
	if err != nil {
		return
	}

	containers, containedBy := parseBags(reader)
	fmt.Printf("Part one: %d\n", partOne(containers, containedBy))
	fmt.Printf("Part two: %d\n", partTwo(containers))

	return
}

func partOne(containers map[string]map[string]int, containedBy map[string][]string) int {
	return countShinyGoldContainers(containers, containedBy)
}

func partTwo(containers map[string]map[string]int) int {
	return countShinyGoldContained(containers)
}
