package cmd

import (
	"fmt"
	"os"

	"github.com/cguertin14/advent-of-code-2020/cmd/day1"
	"github.com/cguertin14/advent-of-code-2020/cmd/day2"
	"github.com/cguertin14/advent-of-code-2020/cmd/day3"
	"github.com/cguertin14/advent-of-code-2020/cmd/day4"
	"github.com/cguertin14/advent-of-code-2020/cmd/day5"
	"github.com/cguertin14/advent-of-code-2020/cmd/day6"
	"github.com/cguertin14/advent-of-code-2020/cmd/day7"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc2020",
	Short: "AOC2020 is a CLI to control aoc challenges.",
	Long:  "AOC2020 is a CLI to control advent of code (https://adventofcode.com/2020) challenges.",
}

func init() {
	// Add daily commands here.
	rootCmd.AddCommand(day1.Command())
	rootCmd.AddCommand(day2.Command())
	rootCmd.AddCommand(day3.Command())
	rootCmd.AddCommand(day4.Command())
	rootCmd.AddCommand(day5.Command())
	rootCmd.AddCommand(day6.Command())
	rootCmd.AddCommand(day7.Command())
}

// Execute is the function that runs the cobra command.
func Execute() (err error) {
	if err = rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return
	}

	return
}
