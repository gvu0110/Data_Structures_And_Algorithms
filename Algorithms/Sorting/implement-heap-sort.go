package main

import (
	"fmt"
	"log"
)

// Heap : contains a slice which holds the underlying heap data
type Heap struct { // the total number of elements, which doesn't go down on extract
	Items []int
}

// GetLeftIndex : get left index of a Heap node
func (h *Heap) GetLeftIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

// GetRightIndex : get right index of a Heap node
func (h *Heap) GetRightIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

// GetParentIndex : get parent index of a Heap node
func (h *Heap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// HasLeft : check if node at index has left node
func (h *Heap) HasLeft(index int) bool {
	return h.GetLeftIndex(index) < len(h.Items)
}

// HasRight : check if node at index has right node
func (h *Heap) HasRight(index int) bool {
	return h.GetRightIndex(index) < len(h.Items)
}

// HasParent : check if node at index has parent node
func (h *Heap) HasParent(index int) bool {
	return h.GetParentIndex(index) >= 0
}

// Left : get left node value, given an index
func (h *Heap) Left(index int) int {
	return h.Items[h.GetLeftIndex(index)]
}

// Right : get right node value, given an index
func (h *Heap) Right(index int) int {
	return h.Items[h.GetRightIndex(index)]
}

// Parent : get parent node value, given an index
func (h *Heap) Parent(index int) int {
	return h.Items[h.GetParentIndex(index)]
}

// Swap : swap values of two nodes at specified indeces
func (h *Heap) Swap(indexOne, indexTwo int) {
	h.Items[indexOne], h.Items[indexTwo] = h.Items[indexTwo], h.Items[indexOne]
}

// MaxHeap : represents a Max heap data structure
type MaxHeap struct {
	*Heap
}

// New : returns a new instance of a Heap
func New(input []int) *MaxHeap {
	h := &MaxHeap{
		&Heap{
			Items: input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMaxHeap()
	}

	return h
}

func (h *MaxHeap) buildMaxHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

// Insert : adds an element to the heap
func (h *MaxHeap) Insert(item int) *MaxHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.maxHeapifyUp(lastElementIndex)

	return h
}

// ExtractMax : returns the maximum element and removes it from the Heap
func (h *MaxHeap) ExtractMax() int {
	if len(h.Items) == 0 {
		log.Fatal("No items in the heap")
	}
	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	// shrinking slice
	h.Items = h.Items[:len(h.Items)-1]

	h.maxHeapifyDown(0)

	return minItem
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.HasParent(index) && h.Parent(index) < h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	// iterate until we have a child node smaller than the index value
	for (h.HasLeft(index) && h.Items[index] < h.Left(index)) ||
		(h.HasRight(index) && h.Items[index] < h.Right(index)) {
		// if both children are smaller
		if (h.HasLeft(index) && h.Items[index] < h.Left(index)) &&
			(h.HasRight(index) && h.Items[index] < h.Right(index)) {
			// compare the children and swap with the largest
			if h.Left(index) > h.Right(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
				index = h.GetRightIndex(index)
			}
			// else if only left child is smaller than swap with it
		} else if h.HasLeft(index) && h.Items[index] < h.Left(index) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
			// else it must be the right child that is smaller, so we swap with it
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

func heapSort(arr []int) []int {
	maxHeap := New(arr)
	size := len(maxHeap.Items)
	for size > 0 {
		size--
		arr[size] = maxHeap.ExtractMax()
	}
	return arr
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Printf("%#v\n", heapSort(numbers))
}
