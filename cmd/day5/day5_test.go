package day5

import (
	"testing"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/stretchr/testify/assert"
)

func fetchOne() (reader utils.Reader) {
	reader = utils.Reader{}
	reader.ReadFromLiteral("BFFFBBFRRR")
	return
}

func fetchTwo() (reader utils.Reader) {
	reader = utils.Reader{}
	reader.ReadFromLiteral(`FFBFBFBRRL
FBFFBFBLRR
FFFBBFBRLR
FBBFFBFRRR
FBBBFFFLRR
FBBBFFBRRL
FBBBBFBRRL
BFBBFFBLLL
BBFFFFFRRR`)
	return
}

func TestBothPartsWork(t *testing.T) {
	t.Run("test part one works", func(t *testing.T) {
		seats := parseSeats(fetchOne())
		assert.Equal(t, 567, partOne(seats))
	})

	t.Run("test part two works", func(t *testing.T) {
		seats := parseSeats(fetchTwo())
		assert.Equal(t, 110, partTwo(seats))
	})
}
