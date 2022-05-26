package fdq

import "os"

type Dequeue interface {
	PushLeft(any) Dequeue
	PopLeft() (any, Dequeue)
	PushRight(any) Dequeue
	PopRight() (any, Dequeue)
	Print(*os.File)
}
