package main

import "fmt"

type HashTable struct {
	data map[string]int
}

func newHashTable(size int) *HashTable {
	return &HashTable{
		data: make(map[string]int, size),
	}
}

func (h *HashTable) set(key string, value int) {
	h.data[key] = value
}

func (h *HashTable) get(key string) int {
	return h.data[key]
}

func main() {
	h := newHashTable(10)
	h.set("grapes", 10)
	h.set("apple", 20)
	h.set("orange", 30)
	fmt.Printf("%#v\n", h)
	fmt.Printf("%d\n", h.get("apple"))
}
