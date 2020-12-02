package day1

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

var (
	day1Cmd = &cobra.Command{
		Use:   "day1",
		Short: "Command used for the day1 challenge.",
		Long:  "Command used for the day1 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Command returns the command.
func Command() *cobra.Command {
	return day1Cmd
}

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFile("cmd/day1/input.txt")
	if err != nil {
		return
	}

	// Get results of each challenge.
	numbers := reader.ToNumbers()
	one := partOne(numbers)
	two := partTwo(numbers)

	// SHOW RESULTS
	fmt.Println(fmt.Sprintf("Part one, the answer is: %d", one))
	fmt.Println(fmt.Sprintf("Part two, the answer is: %d", two))

	return nil
}

func partOne(numbers []int) int {
	for _, num := range numbers {
		for _, otherNum := range numbers {
			if num+otherNum == 2020 {
				return num * otherNum
			}
		}
	}

	return 0
}

func partTwo(numbers []int) int {
	for _, num := range numbers {
		for _, secondNum := range numbers {
			for _, thirdNum := range numbers {
				if num+secondNum+thirdNum == 2020 {
					return num * secondNum * thirdNum
				}
			}
		}
	}

	return 0
}
