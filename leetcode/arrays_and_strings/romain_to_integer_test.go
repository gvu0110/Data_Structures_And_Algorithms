package arrays_and_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func romanToInt(s string) int {
	digitMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	i := 0
	for i < len(s) {
		if i == len(s)-1 {
			result += digitMap[s[i]]
			break
		}

		if digitMap[s[i]] < digitMap[s[i+1]] {
			result += digitMap[s[i+1]] - digitMap[s[i]]
			i += 2
		} else {
			result += digitMap[s[i]]
			i++
		}
	}
	return result
}

func TestRomanToInt(t *testing.T) {
	result := romanToInt("III")
	assert.EqualValues(t, 3, result)
	result = romanToInt("IV")
	assert.EqualValues(t, 4, result)
	result = romanToInt("IX")
	assert.EqualValues(t, 9, result)
	result = romanToInt("LVIII")
	assert.EqualValues(t, 58, result)
	result = romanToInt("MCMXCIV")
	assert.EqualValues(t, 1994, result)
}
