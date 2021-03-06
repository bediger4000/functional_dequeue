package fdq

import (
	"fmt"
	"os"
)

// Implementation of the Dequeue described in Section 3 of "Real-time Deques,
// Multi-head Turing machines...".  The section is titled "Functional Deques with
// good amortized performance".
//
// Basically a two-stack Dequeue, but when an empty stack gets popped, it reverses
// half of the other stack's contents on to the empty stack, and pops an element
// from the now not-empty stack. Halves the number of operations that a simplistic
// 2-stack dequeue would make when an empty stack gets popped.

type HalfStack struct {
	opCount int
	head    *stack
	tail    *stack
}

var emptyHalfStack Dequeue = (*HalfStack)(nil)

const halfstackType = "halfstack"

func init() {
	Implementations = append(Implementations, halfstackType)
	if NewFunctions == nil {
		NewFunctions = make(map[string]Dequeue)
	}
	NewFunctions[halfstackType] = emptyHalfStack

}

func (l *HalfStack) PushLeft(datum any) Dequeue {
	if l == nil {
		l = &HalfStack{}
	}
	l.opCount++
	l.head = l.head.Push(datum)
	return l
}

func (l *HalfStack) PopLeft() (any, Dequeue) {
	if l == nil {
		l = &HalfStack{}
	}
	l.opCount++
	var data any
	if l.head.Size() == 0 && l.tail.Size() > 0 {
		n := l.tail.Size() / 2
		// Linus Torvald's "Good Taste in Programming" technique.
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
			l.head = l.head.PushNode(newhead)
		}
	}
	if l.head.Size() == 0 {
		return nil, l
	}
	data, l.head = l.head.Pop()
	return data, l
}

func (l *HalfStack) PushRight(datum any) Dequeue {
	if l == nil {
		l = &HalfStack{}
	}
	l.opCount++
	l.tail = l.tail.Push(datum)
	return l
}

func (l *HalfStack) PopRight() (any, Dequeue) {
	if l == nil {
		l = &HalfStack{}
	}
	l.opCount++
	var data any
	if l.tail.Size() == 0 && l.head.Size() > 0 {
		n := l.head.Size() / 2

		indirect := &(l.head.stk)
		for i := 0; i < n; i++ {
			indirect = &(*indirect).next
			l.head.opCount++
		}
		newhead := *indirect
		*indirect = nil

		var tmp *stackNode
		for ; newhead != nil; newhead = tmp {
			l.head.size--
			tmp = newhead.next
			l.head.opCount++
			l.tail = l.tail.PushNode(newhead)
		}
	}
	if l.tail.Size() == 0 {
		return nil, l
	}
	data, l.tail = l.tail.Pop()
	return data, l
}

func (l *HalfStack) Print(fout *os.File) {
	if l == nil {
		l = &HalfStack{}
	}
	fmt.Fprintf(fout, "head (%d:%d): ", l.head.Size(), l.head.Operations())
	for p := l.head.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "tail (%d:%d): ", l.tail.Size(), l.tail.Operations())
	for p := l.tail.Node(); p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	stackOps := l.tail.Operations() + l.head.Operations()
	fmt.Fprintf(fout, "Dequeue operations %d, stack operations: %d => %.3f\n", l.opCount, stackOps, float64(stackOps)/float64(l.opCount))
}

func (l *HalfStack) Type() string {
	return halfstackType
}

func (l *HalfStack) Operations() (int, int) {
	return l.opCount, l.tail.Operations() + l.head.Operations()
}
