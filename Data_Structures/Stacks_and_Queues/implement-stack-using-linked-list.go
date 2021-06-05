package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	if s.top == nil {
		s.bottom = &Node{
			value: value,
		}
		s.top = s.bottom
	} else {
		s.top = &Node{
			next:  s.top,
			value: value,
		}
	}
	s.length++
}

func (s *Stack) Pop() {
	if s.top == nil {
		panic("Cannot pop an empty stack!")
	}
	s.top = s.top.next
	s.length--
}

func (s *Stack) Peek() *Node {
	return s.top
}

func main() {
	s := newStack()
	s.Push("google.com")
	s.Push("youtube.com")
	s.Push("discord.com")
	fmt.Printf("%#v\n", s.Peek())
	s.Pop()
	s.Pop()
	fmt.Printf("%#v\n", s.Peek())
	s.Pop()
	fmt.Printf("%#v\n", s.Peek())
}
