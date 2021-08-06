package arrays_and_strings

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func myAtoi(s string) int {
	runesMap := map[string]int{}
	for i := 0; i <= 9; i++ {
		runesMap[fmt.Sprint(i)] = i
	}

	result := 0
	numNegativeSigns := 0
	for _, r := range s {
		if r == ' ' {
			continue
		}

		if r == '-' || r == '+' {
			if r == '-' {
				numNegativeSigns++
			}
			continue
		}

		if r >= '0' && r <= '9' {
			digit := runesMap[string(r)]
			result = result*10 + digit
		} else {
			break
		}
	}

	if numNegativeSigns%2 == 1 {
		result *= -1
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func TestMyAToI(t *testing.T) {
	result := myAtoi("42")
	assert.EqualValues(t, 42, result)
	result = myAtoi("   -42")
	assert.EqualValues(t, -42, result)
	result = myAtoi("-12")
	assert.EqualValues(t, -12, result)
	result = myAtoi("+-12")
	assert.EqualValues(t, -12, result)
	result = myAtoi("--12")
	assert.EqualValues(t, 12, result)
	result = myAtoi("-91283472332")
	assert.EqualValues(t, -2147483648, result)
	result = myAtoi("4193 with words")
	assert.EqualValues(t, 4193, result)
	result = myAtoi("words and 987")
	assert.EqualValues(t, 0, result)
}
