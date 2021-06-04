package main

import "fmt"

type Array struct {
	length int
	data   []string
}

func (a *Array) get(index int) string {
	return a.data[index]
}

func (a *Array) push(item string) {
	a.data = append(a.data, item)
	a.length++
}

func (a *Array) pop() {
	a.data = a.data[:a.length-1]
	a.length--
}

func (a *Array) pushFront(item string) {
	a.data = append([]string{item}, a.data...)
	a.length++
}

func (a *Array) popFront() {
	a.data = a.data[1:]
	a.length--
}

func (a *Array) delete(index int) {
	a.data = append(a.data[:index], a.data[index+1:]...)
	a.length--
}

func (a *Array) expand(index int, size int) {
	a.data = append(a.data[:index], append(make([]string, size), a.data[index:]...)...)
	a.length += size
}

func (a *Array) extend(size int) {
	a.data = append(a.data, make([]string, size)...)
	a.length += size
}

// TODO: cut, copy

func main() {
	a := Array{}
	a.push("my")
	a.push("name")
	a.push("is")
	a.push("Adam")
	a.pushFront("hi")
	fmt.Printf("%#v\n", a)
	a.pop()
	fmt.Printf("%#v\n", a)
	a.popFront()
	fmt.Printf("%#v\n", a)
	a.expand(2, 5)
	fmt.Printf("%#v\n", a)
	a.delete(1)
	fmt.Printf("%#v\n", a)
	a.extend(2)
	fmt.Printf("%#v\n", a)
}
