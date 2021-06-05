package main

import "fmt"

type Stack struct {
	array []string
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value string) {
	s.array = append(s.array, value)
}

func (s *Stack) Pop() {
	s.array = s.array[:len(s.array)-1]
}

func (s *Stack) Peek() string {
	if len(s.array) < 1 {
		return ""
	}
	return s.array[len(s.array)-1]
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
