package day2

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/spf13/cobra"
)

var (
	day2Cmd = &cobra.Command{
		Use:   "day2",
		Short: "Command used for the day2 challenge.",
		Long:  "Command used for the day2 advent-of-code challenge.",
		RunE:  Execute,
	}
)

// Execute runs the command.
func Execute(cmd *cobra.Command, args []string) (err error) {
	reader := utils.Reader{}
	_, err = reader.ReadFromFile("cmd/day2/input.txt")
	if err != nil {
		return
	}

	// 1- Build all passwords.
	passwords := buildPasswords(reader)

	// 2- Validate all passwords (part 1)
	validCount := validatePasswordsA(passwords)

	// 3- Print the total of valid passwords.
	fmt.Println(fmt.Sprintf("There are %d valid passwords in the list. (part 1)", validCount))

	// 4- Validate all passwords (part 2)
	validCount = validatePasswordsB(passwords)

	// 3- Print the total of valid passwords.
	fmt.Println(fmt.Sprintf("There are %d valid passwords in the list. (part 2)", validCount))

	return
}

// Command returns the command.
func Command() *cobra.Command {
	return day2Cmd
}
