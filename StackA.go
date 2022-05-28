package collections

import (
	"fmt"
)

// StackA is array based implementation
type StackA struct {
	Values []interface{}
}

func NewStackA() *StackA {
	return &StackA{}
}

func (s *StackA) Print() {
	fmt.Println("--")
	for i := len(s.Values) - 1; i >= 0; i-- {
		fmt.Println(s.Values[i])
	}
	fmt.Println("--")
}

func (s *StackA) Push(val interface{}) {
	s.Values = append(s.Values, val)
}

func (s *StackA) Pop() (interface{}, bool) {
	if len(s.Values) > 0 {
		lastVal := s.Values[len(s.Values)-1]
		s.Values = append([]interface{}{}, s.Values[:len(s.Values)-1]...)
		return lastVal, true
	} else {
		return nil, false
	}
}

func (s *StackA) Peek() (interface{}, bool) {
	if len(s.Values) > 0 {
		lastVal := s.Values[len(s.Values)-1]
		return lastVal, true
	} else {
		return nil, false
	}
}
