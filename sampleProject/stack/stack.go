package stack

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return -1
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.value
}

func (s *Stack) Push(value int) {
	node := &Node{value: value}
	node.next = s.top
	s.top = node
}

func (s *Stack) Clear() {
	s.top = nil
}

func (s *Stack) Contains(value int) bool {
	current := s.top
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *Stack) Increment() {
	current := s.top
	for current != nil {
		current.value++
		current = current.next
	}
}

func (s *Stack) Print() {
	current := s.top
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func (s *Node) printReverse() {
	if s.next == nil {
		fmt.Print(s.value, " ")
		return
	}
	s.next.printReverse()
	fmt.Print(s.value, " ")
}

func (s *Stack) PrintReverse() {
	if s.top != nil {
		s.top.printReverse()
	}
	fmt.Println()
}
