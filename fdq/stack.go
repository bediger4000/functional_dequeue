package fdq

import "fmt"

type stackNode struct {
	data any
	next *stackNode
}

type stack struct {
	stk     *stackNode
	size    int
	opCount int
}

func (s *stack) Operations() int {
	if s == nil {
		return 0
	}
	return s.opCount
}

func (s *stack) Node() *stackNode {
	if s == nil {
		return nil
	}
	return s.stk
}

func (s *stack) Push(data any) *stack {
	node := &stackNode{data: data}
	return s.PushNode(node)
}

func (s *stack) PushNode(node *stackNode) *stack {
	if s == nil {
		s = &stack{}
	}
	node.next = s.stk
	s.stk = node
	s.size++
	s.opCount++
	return s
}

func (s *stack) Pop() (any, *stack) {
	if s == nil {
		s = &stack{}
	}

	var node *stackNode

	node, s = s.PopNode()

	var data any

	if node != nil {
		data = node.data
	}

	return data, s
}

func (s *stack) PopNode() (*stackNode, *stack) {
	if s == nil {
		s = &stack{}
	}
	s.opCount++

	if s.stk == nil {
		return nil, s
	}

	node := s.stk
	s.stk = s.stk.next
	s.size--
	return node, s
}

func (s *stack) Size() int {
	if s == nil {
		return 0
	}
	return s.size
}

func (s *stack) Print() {
	for node := s.stk; node != nil; node = node.next {
		fmt.Printf("%v -> ", node.data)
	}
}
