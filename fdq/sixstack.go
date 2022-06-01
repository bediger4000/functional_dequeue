package fdq

import (
	"fmt"
	"os"
)

type SixStack struct {
	opCount int
	left    *stack
	tail    *stack
}

var emptySixStack Dequeue = (*SixStack)(nil)

const sixstackType = "sixstack"

func init() {
	Implementations = append(Implementations, sixstackType)
	if NewFunctions == nil {
		NewFunctions = make(map[string]Dequeue)
	}
	NewFunctions[sixstackType] = emptySixStack
}

func (l *SixStack) PushLeft(datum any) Dequeue {
	if l == nil {
		l = &SixStack{}
	}
	l.opCount++
	l.left = l.left.Push(datum)
	return l
}

func (l *SixStack) PopLeft() (any, Dequeue) {
	if l == nil {
		l = &SixStack{}
	}
	l.opCount++
	var data any
	if l.left.Size() == 0 && l.tail.Size() > 0 {
		n := l.tail.Size() / 2
		indirect := &(l.tail.stk)
		for i := 0; i < n; i++ {
			indirect = &(*indirect).next
			l.tail.opCount++
		}
		newhead := *indirect
		*indirect = nil
		var tmp *stackNode
		for ; newhead != nil; newhead = tmp {
			l.tail.size--
			tmp = newhead.next
			l.tail.opCount++
			l.left = l.left.PushNode(newhead)
		}
	}
	if l.left.Size() == 0 {
		return nil, l
	}
	data, l.left = l.left.Pop()
	return data, l
}

func (l *SixStack) PushRight(datum any) Dequeue {
	if l == nil {
		l = &SixStack{}
	}
	l.opCount++
	l.tail = l.tail.Push(datum)
	return l
}

func (l *SixStack) PopRight() (any, Dequeue) {
	if l == nil {
		l = &SixStack{}
	}
	l.opCount++
	var data any
	if l.tail.Size() == 0 && l.left.Size() > 0 {
		n := l.left.Size() / 2

		indirect := &(l.left.stk)
		for i := 0; i < n; i++ {
			indirect = &(*indirect).next
			l.left.opCount++
		}
		newhead := *indirect
		*indirect = nil

		var tmp *stackNode
		for ; newhead != nil; newhead = tmp {
			l.left.size--
			tmp = newhead.next
			l.left.opCount++
			l.tail = l.tail.PushNode(newhead)
		}
	}
	if l.tail.Size() == 0 {
		return nil, l
	}
	data, l.tail = l.tail.Pop()
	return data, l
}

func (l *SixStack) Print(fout *os.File) {
	if l == nil {
		l = &SixStack{}
	}
	fmt.Fprintf(fout, "left (%d:%d): ", l.left.Size(), l.left.Operations())
	for p := l.left.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "tail (%d:%d): ", l.tail.Size(), l.tail.Operations())
	for p := l.tail.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	stackOps := l.tail.Operations() + l.left.Operations()
	fmt.Fprintf(fout, "Dequeue operations %d, stack operations: %d => %.3f\n", l.opCount, stackOps, float64(stackOps)/float64(l.opCount))
}

func (l *SixStack) Type() string {
	return halfstackType
}

func (l *SixStack) Operations() (int, int) {
	return l.opCount, l.tail.Operations() + l.left.Operations()
}
