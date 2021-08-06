package arrays_and_strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func intToRoman(num int) string {
	digitMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	list := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumeral := ""
	for num > 0 {
		for _, value := range list {
			if num >= value {
				romanNumeral = romanNumeral + strings.Repeat(digitMap[value], num/value)
				num %= value
			}
		}
	}
	return romanNumeral
}

func TestIntToRoman(t *testing.T) {
	result := intToRoman(3)
	assert.EqualValues(t, "III", result)
	result = intToRoman(4)
	assert.EqualValues(t, "IV", result)
	result = intToRoman(9)
	assert.EqualValues(t, "IX", result)
	result = intToRoman(58)
	assert.EqualValues(t, "LVIII", result)
	result = intToRoman(1994)
	assert.EqualValues(t, "MCMXCIV", result)
}
