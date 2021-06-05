package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Prepend(value interface{}) {
	if l.head == nil {
		l.tail = &Node{
			value: value,
		}
		l.head = l.tail
	} else {
		l.head = &Node{
			next:  l.head,
			value: value,
		}
	}
	l.size++
}

func (l *LinkedList) Append(value interface{}) {
	if l.tail == nil {
		l.tail = &Node{
			value: value,
		}
		l.head = l.tail
	} else {
		newNode := &Node{
			value: value,
		}
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList) GetAt(index int) *Node {
	node := l.head
	if index < 0 {
		return nil
	}
	if index > (l.size - 1) {
		return nil
	}
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *LinkedList) Insert(index int, value interface{}) {
	if index == 0 {
		l.Prepend(value)
	} else if index < 0 {
		panic("Index can not be negative")
	} else if index == l.size {
		l.Append(value)
	} else if index > l.size {
		panic("Index is larger than linked list size")
	} else {
		node := l.GetAt(index - 1)
		node.next = &Node{
			next:  node.next,
			value: value,
		}
		l.size++
	}
}

func (l *LinkedList) Delete(index int) {
	if index == 0 {
		l.head = l.head.next
		l.size--
	} else if index < 0 {
		panic("Index can not be negative")
	} else if index > (l.size - 1) {
		panic("Index is larger than linked list size")
	} else {
		node := l.GetAt(index)
		prevNode := l.GetAt(index - 1)
		prevNode.next = node.next
		if prevNode.next == nil {
			l.tail = prevNode
		}
		l.size--
	}
}

func (l *LinkedList) Display() {
	node := l.head
	for node != nil {
		fmt.Printf("%+v -> ", node.value)
		node = node.next
	}
	fmt.Println()
	fmt.Printf("Size: %d\n", l.size)
	fmt.Printf("Head: %d\n", l.head.value)
	fmt.Printf("Tail: %d\n", l.tail.value)
}

func main() {
	l1 := newLinkedList()
	l1.Prepend(5)
	l1.Prepend(9)
	l1.Append(3)
	l1.Append(6)
	fmt.Printf("%#v\n", l1.GetAt(0))
	l1.Insert(2, 99)
	l1.Delete(4)
	l1.Display()

	l2 := newLinkedList()
	l2.Append(3)
	l2.Append(6)
	l2.Prepend(5)
	l2.Prepend(9)
	l2.Display()
}
