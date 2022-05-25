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
	return l
}
func (l *ListQueue) PopLeft() (any, Dequeue) {
	return nil, l
}

func (l *ListQueue) PushRight(datum any) Dequeue {
	return l
}

func (l *ListQueue) PopRight() (any, Dequeue) {
	return nil, l
}
