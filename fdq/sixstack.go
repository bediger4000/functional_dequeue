package fdq

import (
	"fmt"
	"os"
)

type Stack6 struct {
	opCount int
	head    *stack
	tail    *stack

	currentPopL  func(*Stack6) any
	currentPushL func(*Stack6, any)
	currentPopR  func(*Stack6) any
	currentPushR func(*Stack6, any)
}

var emptySixStack Dequeue = (*Stack6)(nil)

const sixstackType = "stack6"

func init() {
	Implementations = append(Implementations, sixstackType)
	if NewFunctions == nil {
		NewFunctions = make(map[string]Dequeue)
	}
	NewFunctions[sixstackType] = emptySixStack

}

func newStack6() *Stack6 {
	return &Stack6{
		currentPopL:  smallPopLeft,
		currentPushL: smallPushLeft,
		currentPopR:  smallPopRight,
		currentPushR: smallPushRight,
	}
}

func (l *Stack6) PushLeft(datum any) Dequeue {
	if l == nil {
		l = newStack6()
	}
	l.opCount++
	l.currentPushL(l, datum)
	return l
}

func smallPushLeft(l *Stack6, datum any) {
	l.head = l.head.Push(datum)
}

func (l *Stack6) PopLeft() (any, Dequeue) {
	if l == nil {
		l = newStack6()
	}
	l.opCount++
	data := l.currentPopL(l)
	return data, l
}

func smallPopLeft(l *Stack6) any {
	var data any
	if l.head.Size() == 0 {
		for l.tail.Size() > 0 {
			data, l.tail = l.tail.Pop()
			l.head = l.head.Push(data)
		}
	}
	if l.head == nil {
		return nil
	}
	data, l.head = l.head.Pop()
	return data
}

func (l *Stack6) PushRight(datum any) Dequeue {
	if l == nil {
		l = &Stack6{}
	}
	l.opCount++
	l.tail = l.tail.Push(datum)
	return l
}

func smallPushRight(l *Stack6, datum any) {
	l.tail = l.tail.Push(datum)
}

func (l *Stack6) PopRight() (any, Dequeue) {
	if l == nil {
		l = newStack6()
	}
	l.opCount++
	data := l.currentPopR(l)
	return data, l
}

func smallPopRight(l *Stack6) any {
	var data any
	if l.tail.Size() == 0 {
		for l.head.Size() > 0 {
			data, l.head = l.head.Pop()
			l.tail = l.tail.Push(data)
		}
	}
	if l.tail == nil {
		return nil
	}
	data, l.tail = l.tail.Pop()
	return data
}

func (l *Stack6) Print(fout *os.File) {
	if l == nil {
		l = &Stack6{}
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

func (l *Stack6) Type() string {
	return sixstackType
}

func (l *Stack6) Operations() (int, int) {
	return l.opCount, l.tail.Operations() + l.head.Operations()
}
