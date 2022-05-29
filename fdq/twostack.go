package fdq

import (
	"fmt"
	"os"
)

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
	var data any
	if l.head.Size() == 0 {
		for l.tail.Size() > 0 {
			data, l.tail = l.tail.Pop()
			l.head = l.head.Push(data)
		}
	}
	if l.head.Size() == 0 {
		return nil, l
	}
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
	if l.tail.Size() == 0 {
		var data any
		for l.head.Size() > 0 {
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
	fmt.Fprintf(fout, "head (%d): ", l.head.Size())
	for p := l.head.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "tail (%d): ", l.tail.Size())
	for p := l.tail.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	stackOps := l.tail.Operations() + l.head.Operations()
	fmt.Fprintf(fout, "Dequeue operations %d, stack operations: %d => %.3f\n", l.opCount, stackOps, float64(stackOps)/float64(l.opCount))
}

func (l *TwoStack) Type() string {
	return "twostack"
}

func (l *TwoStack) Operations() (int, int) {
	return l.opCount, l.tail.Operations() + l.head.Operations()
}
