package fdq

import (
	"fmt"
	"os"
)

type stackNode struct {
	data any

	next *stackNode
}

type TwoStack struct {
	head *stackNode
	tail *stackNode
}

var _ Dequeue = (*TwoStack)(nil)

func (l *TwoStack) PushLeft(datum any) Dequeue {
	if l == nil {
		l = &TwoStack{}
	}
	node := &stackNode{data: datum}
	node.next = l.head
	l.head = node
	return l
}
func (l *TwoStack) PopLeft() (any, Dequeue) {
	if l == nil {
		l = &TwoStack{}
	}
	if l.head == nil {
		for node := l.tail; node != nil; node = l.tail {
			l.tail = l.tail.next
			node.next = l.head
			l.head = node
		}
	}
	if l.head == nil {
		return nil, l
	}
	node := l.head
	l.head = l.head.next
	return node.data, l
}

func (l *TwoStack) PushRight(datum any) Dequeue {
	if l == nil {
		l = &TwoStack{}
	}
	node := &stackNode{data: datum}
	node.next = l.tail
	l.tail = node
	return l
}

func (l *TwoStack) PopRight() (any, Dequeue) {
	if l == nil {
		l = &TwoStack{}
	}
	if l.tail == nil {
		for node := l.head; node != nil; node = l.head {
			l.head = l.head.next
			node.next = l.tail
			l.tail = node
		}
	}
	if l.tail == nil {
		return nil, l
	}
	node := l.tail
	l.tail = l.tail.next
	return node.data, l
}

func (l *TwoStack) Print(fout *os.File) {
	if l == nil {
		l = &TwoStack{}
	}
	fmt.Fprintf(fout, "head: ")
	for p := l.head; p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
	fmt.Fprintf(fout, "tail: ")
	for p := l.tail; p != nil; p = p.next {
		fmt.Fprintf(fout, "%s -> ", p.data)
	}
	fmt.Fprintf(fout, "\n")
}
