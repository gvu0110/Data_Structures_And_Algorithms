//Given an array = [2,5,1,2,3,5,1,2,4]:
//It should return 2

//Given an array = [2,1,1,2,3,5,1,2,4]:
//It should return 1

//Given an array = [2,3,4,5]:
//It should return undefined

package main

import (
	"fmt"
	"strconv"
)

func firstRecurringCharacter(input []int) string {
	hashTable := make(map[int]string, len(input))
	for _, num := range input {
		if _, ok := hashTable[num]; !ok {
			hashTable[num] = ""
		} else {
			return strconv.Itoa(num)
		}
	}
	return "undefined"
}

func main() {
	fmt.Println(firstRecurringCharacter([]int{2, 5, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(firstRecurringCharacter([]int{2, 1, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(firstRecurringCharacter([]int{2, 3, 4, 5}))
}
