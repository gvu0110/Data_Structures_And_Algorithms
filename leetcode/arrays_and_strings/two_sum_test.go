package arrays_and_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	diff := map[int]int{}
	for i, value := range nums {
		j, ok := diff[value]
		if ok {
			result[0] = j
			result[1] = i
		} else {
			diff[target-value] = i
		}
	}
	return result
}

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	assert.EqualValues(t, []int{0, 1}, result)
	result = twoSum([]int{3, 2, 4}, 6)
	assert.EqualValues(t, []int{1, 2}, result)
	result = twoSum([]int{3, 3}, 6)
	assert.EqualValues(t, []int{0, 1}, result)
}
