package main

import "fmt"

func findFactorialIterative(n int) int {
	if n < 1 {
		return 0
	}
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func findFactorialRecursive(n int) int {
	if n > 1 {
		return n * findFactorialRecursive(n-1)
	} else if n == 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(findFactorialIterative(5))
	fmt.Println(findFactorialRecursive(6))
}
