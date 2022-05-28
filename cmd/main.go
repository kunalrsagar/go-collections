package main

import (
	"fmt"
	"github.com/kunalrsagar/collections"
	"log"
)

func main() {
	// queue()
}

func queue() {
	q := collections.NewQueue()
	q.Enqueue(1)
	q.Peek()
	q.Enqueue(2)
	q.Peek()
	q.Enqueue(3)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

func linkedList() {
	// LINKED LIST
	ll := collections.NewLinkedList(10)
	ll.Append(2)
	ll.Append(19)
	ll.Append(20)
	ll.Prepend(18)
	err := ll.Insert(2, "🚀")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = ll.Remove(1)
	if err != nil {
		log.Fatal(err.Error())
	}

	ll.Print()
	ll.Reverse()
	ll.Print()
}

func stack_l() {
	// STACK LinkedList implementation
	s := collections.NewStack()
	s.Push(10)
	s.Push(19)
	s.Push(20)

	if val, ok := s.Peek(); ok {
		log.Println(val)
	}
	s.Print()

	if val, ok := s.Pop(); ok {
		log.Println(val)
	}
	s.Print()

	s.Pop()
	s.Print()

	s.Pop()
	s.Print()
}

func stack_a() {
	// STACK array based implementation
	s := collections.NewStackA()
	s.Push(10)
	s.Push(19)
	s.Push(20)

	if val, ok := s.Peek(); ok {
		log.Println(val)
	}
	s.Print()

	if val, ok := s.Pop(); ok {
		log.Println(val)
	}
	s.Print()

	s.Pop()
	s.Print()

	s.Pop()
	s.Print()
}
