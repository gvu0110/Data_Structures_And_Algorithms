// Implement a queue using only two stacks
package main

import "fmt"

type Stack struct {
	array []string
}

type Queue struct {
	s1 Stack
	s2 Stack
}

func newQueue() *Queue {
	return &Queue{}
}

func (s *Stack) Push(value string) {
	s.array = append(s.array, value)
}

func (s *Stack) Pop() string {
	removeItem := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return removeItem
}

func (s *Stack) Peek() string {
	return s.array[len(s.array)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.array) == 0
}

// Enqueue element to the back of queue
func (q *Queue) Enqueue(value string) {
	q.s1.Push(value)
}

// Removes the element from in front of queue and returns that element
// Solutions:
// Pop all elements from s1 and to push them on to an additional stack s2, which stores the elements of s1 in reversed order
// This way the bottom element of s1 will be positioned on top of s2 and we can simply pop it from stack s2
// Once s2 is empty, the algorithm transfer data from s1 to s2 again
func (q *Queue) Dequeue() string {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Pop()
}

func (q *Queue) Peek() string {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Peek()
}

func (q *Queue) IsEmpty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}

func main() {
	q := newQueue()
	q.Enqueue("Adam")
	q.Enqueue("Katie")
	q.Enqueue("Mikhaila")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Peek())
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
}
