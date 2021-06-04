package main

import (
	"fmt"
	"sort"
)

func mergeSortedArrays(a1 []int, a2 []int) []int {
	a := append(a1, a2...)
	sort.Ints(a)
	return a
}

func main() {
	fmt.Printf("%#v\n", mergeSortedArrays([]int{0, 3, 4, 31}, []int{3, 4, 6, 30}))
}
