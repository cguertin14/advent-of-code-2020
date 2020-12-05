package day4

import (
	"testing"

	"github.com/cguertin14/advent-of-code-2020/utils"
	"github.com/stretchr/testify/assert"
)

func fetchOne() (reader utils.Reader) {
	reader = utils.Reader{}
	reader.ReadFromLiteralDouble(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)
	return
}

func fetchTwo() (reader utils.Reader) {
	reader = utils.Reader{}
	reader.ReadFromLiteralDouble(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`)
	return
}

func TestBothPartsWork(t *testing.T) {
	t.Run("test part one works", func(t *testing.T) {
		passwords := readPassports(fetchOne())
		assert.Equal(t, 2, partOne(passwords))
	})

	t.Run("test part two works", func(t *testing.T) {
		passwords := readPassports(fetchTwo())
		assert.Equal(t, 3, partTwo(passwords))
	})
}
