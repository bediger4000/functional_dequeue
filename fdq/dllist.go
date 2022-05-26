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
	head *listNode
	tail *listNode
}

var _ Dequeue = (*ListQueue)(nil)

func NewDllist() Dequeue {
	return &ListQueue{}
}

func (l *ListQueue) PushLeft(datum any) Dequeue {
	node := &listNode{data: datum}
	if l.head == nil {
		l.head = node
		l.tail = node
		return l
	}
	node.next = l.head
	l.head.prev = node
	l.head = node
	return l
}
func (l *ListQueue) PopLeft() (any, Dequeue) {
	if l.head == nil {
		return nil, l
	}

	node := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}

	return node.data, l
}

func (l *ListQueue) PushRight(datum any) Dequeue {
	node := &listNode{data: datum}
	if l.tail == nil {
		l.head = node
		l.tail = node
		return l
	}
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
	return l
}

func (l *ListQueue) PopRight() (any, Dequeue) {
	if l.tail == nil {
		return nil, l
	}

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
}
