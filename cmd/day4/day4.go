package day4

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

var (
	day4Cmd = &cobra.Command{
		Use:   "day4",
		Short: "Command used for the day4 challenge.",
		Long:  "Command used for the day4 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Command returns the command.
func Command() *cobra.Command {
	return day4Cmd
}

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFileDouble("cmd/day4/input.txt")
	if err != nil {
		return
	}

	passports := readPassports(reader)
	fmt.Printf("Part One: %d \n", partOne(passports))
	fmt.Printf("Part Two: %d \n", partTwo(passports))

	return
}

func partOne(passports []Passport) (count int) {
	for _, pass := range passports {
		if pass.IsValid1() {
			count++
		}
	}
	return
}

func partTwo(passports []Passport) (count int) {
	for _, pass := range passports {
		if pass.IsValid2() {
			count++
		}
	}
	return
}
