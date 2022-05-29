package fdq

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
	if s == nil {
		s = &stack{}
	}
	s.opCount++
	node := &stackNode{data: data}
	node.next = s.stk
	s.stk = node
	s.size++
	return s
}

func (s *stack) Pop() (any, *stack) {
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
	return node.data, s
}

func (s *stack) Size() int {
	if s == nil {
		return 0
	}
	return s.size
}
