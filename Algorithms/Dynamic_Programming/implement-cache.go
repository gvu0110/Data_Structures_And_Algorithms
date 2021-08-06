package main

import "fmt"

func memoizedAddTo80() func(int) int {
	cache := make(map[int]int)
	// Use closures for caching
	return func(n int) int {
		if _, ok := cache[n]; !ok {
			fmt.Println("Take long time")
			cache[n] = n + 80
		}
		return cache[n]
	}
}

func main() {
	memoized := memoizedAddTo80()
	fmt.Printf("1st call: %d\n", memoized(5))
	fmt.Printf("2nd call: %d\n", memoized(5))
	fmt.Printf("3rd call: %d\n", memoized(5))
}
