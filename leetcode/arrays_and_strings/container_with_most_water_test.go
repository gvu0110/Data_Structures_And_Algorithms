package arrays_and_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxArea(height []int) int {
	max := 0
	area := 0
	leftIndex := 0
	rightIndex := len(height) - 1
	for leftIndex < rightIndex {
		if height[leftIndex] >= height[rightIndex] {
			area = (rightIndex - leftIndex) * height[rightIndex]
			rightIndex--
		} else {
			area = (rightIndex - leftIndex) * height[leftIndex]
			leftIndex++
		}

		if max < area {
			max = area
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.EqualValues(t, 49, result)
	result = maxArea([]int{1, 1})
	assert.EqualValues(t, 1, result)
	result = maxArea([]int{4, 3, 2, 1, 4})
	assert.EqualValues(t, 16, result)
	result = maxArea([]int{1, 2, 1})
	assert.EqualValues(t, 2, result)
}
