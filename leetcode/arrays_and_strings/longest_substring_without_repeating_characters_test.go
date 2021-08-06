package arrays_and_strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {
	result := 0
	bytesMap := map[byte]int{}
	startIndex := 0

	for endIndex := 0; endIndex < len(s); endIndex++ {
		value, ok := bytesMap[s[endIndex]]
		if ok {
			if value >= startIndex {
				startIndex = value + 1
			}
		}

		bytesMap[s[endIndex]] = endIndex
		fmt.Println(startIndex)
		if result < (endIndex - startIndex + 1) {
			result = endIndex - startIndex + 1
		}
	}
	return result
}

func TestLengthOfLongestSubstring(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")
	assert.EqualValues(t, 3, result)
	result = lengthOfLongestSubstring("bbbbb")
	assert.EqualValues(t, 1, result)
	result = lengthOfLongestSubstring("pwwkew")
	assert.EqualValues(t, 3, result)
	result = lengthOfLongestSubstring("")
	assert.EqualValues(t, 0, result)
	result = lengthOfLongestSubstring(" ")
	assert.EqualValues(t, 1, result)
	result = lengthOfLongestSubstring("abba")
	assert.EqualValues(t, 2, result)
}
