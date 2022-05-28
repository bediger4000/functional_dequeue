package fdq

import (
	"fmt"
	"os"
)

type stackNode struct {
	data any

	next *stackNode
}

type stack struct {
	stk     *stackNode
	opCount int
}

type TwoStack struct {
	opCount int
	head    *stack
	tail    *stack
}

var _ Dequeue = (*TwoStack)(nil)

func init() {
	Implementations = append(Implementations, "twostack")
}

func (l *TwoStack) PushLeft(datum any) Dequeue {
	if l == nil {
		l = &TwoStack{}
	}
	l.opCount++
	l.head = l.head.Push(datum)
	return l
}

func (l *TwoStack) PopLeft() (any, Dequeue) {
	if l == nil {
		l = &TwoStack{}
	}
	l.opCount++
	if l.head.stk == nil {
		var data any
		for l.tail.stk != nil {
			data, l.tail = l.tail.Pop()
			l.head = l.head.Push(data)
		}
	}
	if l.head == nil {
		return nil, l
	}
	var data any
	data, l.head = l.head.Pop()
	return data, l
}

func (l *TwoStack) PushRight(datum any) Dequeue {
	if l == nil {
		l = &TwoStack{}
	}
	l.opCount++
	l.tail = l.tail.Push(datum)
	return l
}

func (l *TwoStack) PopRight() (any, Dequeue) {
	if l == nil {
		l = &TwoStack{}
	}
	l.opCount++
	if l.tail.stk == nil {
		var data any
		for l.head.stk != nil {
			data, l.head = l.head.Pop()
			l.tail = l.tail.Push(data)
		}
	}
	if l.tail == nil {
		return nil, l
	}
	var data any
	data, l.tail = l.tail.Pop()
	return data, l
}

func (l *TwoStack) Print(fout *os.File) {
	if l == nil {
		l = &TwoStack{}
	}
	fmt.Fprintf(fout, "head: ")
	for p := l.head.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "tail: ")
	for p := l.tail.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	stackOps := l.tail.Operations() + l.head.Operations()
	fmt.Fprintf(fout, "Dequeue operations %d, stack operations: %d\n", l.opCount, stackOps)
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
	return node.data, s
}
