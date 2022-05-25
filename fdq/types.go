package fdq

type Dequeue interface {
	PushLeft(any) Dequeue
	PopLeft() (any, Dequeue)
	PushRight(any) Dequeue
	PopRight() (any, Dequeue)
}
