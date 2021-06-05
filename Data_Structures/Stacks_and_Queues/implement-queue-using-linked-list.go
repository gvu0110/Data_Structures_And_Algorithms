package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func newQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value interface{}) {
	if q.first == nil {
		q.last = &Node{
			value: value,
		}
		q.first = q.last
	} else {
		q.last.next = &Node{
			value: value,
		}
		q.last = q.last.next
	}
	q.length++
}

func (q *Queue) Dequeue() {
	if q.first == nil {
		panic("Cannot dequeue an empty queue!")
	}
	q.first = q.first.next
	q.length--
}

func (q *Queue) Peek() *Node {
	return q.first
}

func main() {
	q := newQueue()
	q.Enqueue("Joy")
	q.Enqueue("Matt")
	q.Enqueue("Pavel")
	q.Enqueue("Samir")
	fmt.Printf("%#v\n", q.Peek())
	q.Dequeue()
	q.Dequeue()
	fmt.Printf("%#v\n", q.Peek())
	q.Dequeue()
	q.Dequeue()
	fmt.Printf("%#v\n", q.Peek())
}
