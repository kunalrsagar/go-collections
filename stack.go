package collections

import (
	"fmt"
)

// Stack using LinkedList implementation
type Stack struct {
	Top    *Node
	Bottom *Node
	Length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.Top != nil {
		return s.Top.Value, true
	} else {
		return nil, false
	}
}

func (s *Stack) Push(value interface{}) {
	node := NewNode(value)
	if s.Length == 0 {
		s.Top = node
		s.Bottom = node
	} else {
		temp := s.Top
		s.Top = node
		s.Top.Next = temp
		temp.Prev = s.Top
	}
	s.Length++
}

func (s *Stack) Print() {
	n := s.Top
	fmt.Println("")
	fmt.Println("--")
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
	fmt.Println("--")
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.Length == 0 {
		return nil, false
	}

	popNode := s.Top
	if s.Top.Next != nil {
		s.Top = s.Top.Next
		s.Top.Prev = nil
	} else {
		s.Top = nil
		s.Bottom = nil
	}
	s.Length--

	return popNode.Value, true
}
