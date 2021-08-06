package main

import "fmt"

func newMemoizedFibonancci() func(int) int {
	cache := make(map[int]int)
	var fibonancci func(int) int
	fibonancci = func(n int) int {
		if n == 1 || n == 2 {
			cache[n] = 1
		}
		if _, ok := cache[n]; !ok {
			cache[n] = fibonancci(n-1) + fibonancci(n-2)
		}
		return cache[n]
	}
	return fibonancci
}

// TODO: implement with struct

func main() {
	fib := newMemoizedFibonancci()
	fmt.Println(fib(8))
}
