package fdq

type listNode struct {
	data any

	prev *listNode
	next *listNode
}

type ListQueue struct {
	head *listNode
	tail *listNode
}

func NewDllist() Dequeue {
	return &ListQueue{}
}

//var _ Dequeque = (*ListQueue)(nil)

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
	l.head.prev = nil

	return node.data, l
}

func (l *ListQueue) PushRight(datum any) Dequeue {
	return l
}

func (l *ListQueue) PopRight() (any, Dequeue) {
	return nil, l
}
