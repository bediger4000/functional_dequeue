package fdq

import "os"

var Implementations []string

type Dequeue interface {
	PushLeft(any) Dequeue
	PopLeft() (any, Dequeue)
	PushRight(any) Dequeue
	PopRight() (any, Dequeue)
	Print(*os.File)
	Type() string
	Operations() (int, int) // Dequeue op count, internal op count
}
