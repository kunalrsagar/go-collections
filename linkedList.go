package collections

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func NewLinkedList(value interface{}) *LinkedList {
	head := NewNode(value)
	return &LinkedList{Head: head, Tail: head, Length: 1}
}

func (l *LinkedList) Append(value interface{}) {
	newNode := NewNode(value)
	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode
	l.Length++
}

func (l *LinkedList) Prepend(value interface{}) {
	newNode := NewNode(value)
	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
	l.Length++
}

func (l *LinkedList) Insert(index int, value interface{}) error {

	// Check index, it should be greater than 0
	// it should be less than length of LinkedList so that overflow is avoided.
	if index > l.Length || index < 0 {
		return errors.New("insert value: index should >= 0 OR < length of LinkedList")
	}

	// If value needs to be head then prepend method can be used
	if index == 0 {
		l.Prepend(value)
		return nil
	}

	// If value needs to be at tail then append method can be used
	if index == l.Length {
		l.Append(value)
		return nil
	}

	// If value needs to be inserted in between head and tail,
	// traverse to reach the index
	// create new node
	newNode := NewNode(value)

	leader := l.traverseToIndex(index - 1)
	follower := l.traverseToIndex(index)

	follower.Prev = newNode
	newNode.Next = follower
	leader.Next = newNode
	newNode.Prev = leader
	l.Length++

	return nil
}

func (l *LinkedList) Remove(index int) error {
	if index >= l.Length || index < 0 {
		return errors.New("remove value: index should >= 0 OR < length of LinkedList")
	}

	if index == 0 {
		if l.Head.Next != nil {
			// if next item exist in head, make it head
			l.Head = l.Head.Next
			l.Length--
			return nil
		} else {
			// if Head is the only item that needs to be removed,
			// then make head nil
			l.Head = nil
			return nil
		}
	}

	leader := l.traverseToIndex(index - 1)
	unwantedNode := leader.Next
	follower := unwantedNode.Next
	leader.Next = follower
	follower.Prev = leader
	l.Length--

	return nil
}

func (l *LinkedList) traverseToIndex(index int) *Node {
	n := l.Head
	i := 0
	for n != nil && i < index {
		n = n.Next
		i++
	}
	return n
}

// Reverse function will reverse the linkedList without using doubly Node
func (l *LinkedList) Reverse() {
	if l.Head.Next == nil {
		return
	}

	first := l.Head
	l.Tail = l.Head
	second := l.Head.Next
	for second != nil {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
	}
	l.Head.Next = nil
	l.Head = first
}

func (l *LinkedList) Print() {
	n := l.Head
	i := 0
	fmt.Print("[")
	for n != nil {
		fmt.Print(n.Value)
		i++
		n = n.Next
		if n != nil {
			fmt.Print(",")
		}
	}
	fmt.Printf("]; Length %d\n", l.Length)
}
