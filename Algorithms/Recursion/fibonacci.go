package main

import "fmt"

func fibonacciRecursive(i int) int {
	if i < 0 {
		return 0
	} else if i < 2 {
		return i
	} else {
		return fibonacciRecursive(i-1) + fibonacciRecursive(i-2)
	}
}

func fibonacciIterative(i int) int {
	if i < 0 {
		return 0
	}
	if 0 < i && i < 2 {
		return i
	}
	array := []int{0, 1}
	for n := 2; n <= i; n++ {
		array = append(array, array[n-1]+array[n-2])
	}
	return array[i]
}

func main() {
	fmt.Println(fibonacciRecursive(8))
	fmt.Println(fibonacciIterative(8))
}
