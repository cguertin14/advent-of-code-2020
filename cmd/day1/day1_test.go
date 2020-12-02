package day1

import (
	"testing"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/stretchr/testify/assert"
)

func fetchNumbers(t *testing.T) []int {
	reader := utils.Reader{}
	literal := `
1721
979
366
299
675
1456
	`

	reader.ReadFromLiteral(literal)
	return reader.ToNumbers()
}

func TestBothPartsWork(t *testing.T) {
	t.Run("test part one works", func(t *testing.T) {
		one := partOne(fetchNumbers(t))
		assert.Equal(t, 514579, one)
	})

	t.Run("test part two works", func(t *testing.T) {
		two := partTwo(fetchNumbers(t))
		assert.Equal(t, 241861950, two)
	})
}
