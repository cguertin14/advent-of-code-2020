package day2

import (
	"testing"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/stretchr/testify/assert"
)

func fetchPasswords(literal string) (reader utils.Reader) {
	reader = utils.Reader{}
	reader.ReadFromLiteral(literal)
	return
}

func TestBothPartsWork(t *testing.T) {
	t.Run("test part one works", func(t *testing.T) {
		reader := fetchPasswords(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`)

		totalValidPwds := validatePasswordsA(buildPasswords(reader))
		assert.Equal(t, 2, totalValidPwds)
	})

	t.Run("test part two works", func(t *testing.T) {
		reader := fetchPasswords(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`)
		totalValidPwds := validatePasswordsB(buildPasswords(reader))
		assert.Equal(t, 1, totalValidPwds)
	})
}
