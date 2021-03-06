package fdq

import (
	"fmt"
	"os"
)

type listNode struct {
	data any

	prev *listNode
	next *listNode
}

type ListQueue struct {
	opCount int
	head    *listNode
	tail    *listNode
}

const dllistType = "dllist"

var emptyDllist Dequeue = (*ListQueue)(nil)

func init() {
	Implementations = append(Implementations, dllistType)
	if NewFunctions == nil {
		NewFunctions = make(map[string]Dequeue)
	}
	NewFunctions[dllistType] = emptyDllist

}

func (l *ListQueue) PushLeft(datum any) Dequeue {
	if l == nil {
		l = &ListQueue{}
	}
	node := &listNode{data: datum}
	if l.head == nil {
		l.opCount++
		l.head = node
		l.tail = node
		return l
	}
	l.opCount++
	node.next = l.head
	l.head.prev = node
	l.head = node
	return l
}
func (l *ListQueue) PopLeft() (any, Dequeue) {
	if l == nil {
		l = &ListQueue{}
	}
	if l.head == nil {
		return nil, l
	}

	l.opCount++
	node := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}

	return node.data, l
}

func (l *ListQueue) PushRight(datum any) Dequeue {
	if l == nil {
		l = &ListQueue{}
	}
	node := &listNode{data: datum}
	if l.tail == nil {
		l.opCount++
		l.head = node
		l.tail = node
		return l
	}
	l.opCount++
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
	return l
}

func (l *ListQueue) PopRight() (any, Dequeue) {
	if l == nil {
		l = &ListQueue{}
	}
	if l.tail == nil {
		return nil, l
	}

	l.opCount++
	node := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return node.data, l
}

func (l *ListQueue) Print(fout *os.File) {
	if l == nil {
		l = &ListQueue{}
	}
	fmt.Fprintf(fout, "Left to right: ")
	for p := l.head; p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "Right to left: ")
	for p := l.tail; p != nil; p = p.prev {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "Operations: %d\n", l.opCount)
}

func (l *ListQueue) Type() string {
	return "dllist"
}

func (l *ListQueue) Operations() (int, int) {
	return l.opCount, l.opCount
}
