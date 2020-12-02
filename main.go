package main

import (
	"fmt"
	"os"

	"github.com/cguertin14/advent-of-code-2020/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
