package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

func selectionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		minIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}
	return numbers
}

func insertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

func mergeSort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}
	middle := int(len(numbers) / 2)
	left := make([]int, middle)
	right := make([]int, len(numbers)-middle)
	for i := 0; i < len(numbers); i++ {
		if i < middle {
			left[i] = numbers[i]
		} else {
			right[i-middle] = numbers[i]
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func quickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	splitIndex, rightIndex := 0, len(numbers)-1
	pivotIndex := rand.Int() % len(numbers)
	numbers[pivotIndex], numbers[rightIndex] = numbers[rightIndex], numbers[pivotIndex]
	for i, _ := range numbers {
		if numbers[i] < numbers[rightIndex] {
			numbers[splitIndex], numbers[i] = numbers[i], numbers[splitIndex]
			splitIndex++
		}
	}
	numbers[splitIndex], numbers[rightIndex] = numbers[rightIndex], numbers[splitIndex]
	quickSort(numbers[:splitIndex])
	quickSort(numbers[splitIndex+1:])
	return numbers
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Printf("%#v\n", bubbleSort(numbers))
	fmt.Printf("%#v\n", selectionSort(numbers))
	fmt.Printf("%#v\n", insertionSort(numbers))
	fmt.Printf("%#v\n", mergeSort(numbers))
	fmt.Printf("%#v\n", quickSort(numbers))
}
