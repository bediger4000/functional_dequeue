package fdq

import (
	"fmt"
	"os"
)

// Partial implementation of the functional, real-time (i.e. amortized O(1))
// double-ended queue of section 5 of "Real-time Deques, Multi-head Turing
// machines..." Doesn't do transfers over m dequeue operations, does them all
// in a single batch. Exists so I can understand section 5 of the paper.

type Stack6 struct {
	opCount      int
	tempStackOps int
	head         *stack
	tail         *stack

	currentPopL  func(*Stack6) any
	currentPushL func(*Stack6, any)
	currentPopR  func(*Stack6) any
	currentPushR func(*Stack6, any)
}

var emptySixStack Dequeue = (*Stack6)(nil)

const sixstackType = "stack6a"

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
	if l.tail.Size()+l.head.Size() == 4 {
		l.rearrange4size()
	}
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
		l = newStack6()
	}
	l.opCount++
	l.currentPushR(l, datum)
	return l
}

func smallPushRight(l *Stack6, datum any) {
	l.tail = l.tail.Push(datum)
	if l.tail.Size()+l.head.Size() == 4 {
		l.rearrange4size()
	}
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
		l = newStack6()
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
	stackOps := l.tail.Operations() + l.head.Operations() + l.tempStackOps
	fmt.Fprintf(fout, "Dequeue operations %d, stack operations: %d => %.3f\n", l.opCount, stackOps, float64(stackOps)/float64(l.opCount))
}

func (l *Stack6) Type() string {
	return sixstackType
}

func (l *Stack6) Operations() (int, int) {
	return l.opCount, l.tail.Operations() + l.head.Operations()
}

// rearrange4size called when a push (L or R) gets the Dequeque
// to 4 items. Rearrange to 2 on head stack, 2 on tail stack,
// set currenct functions to "large size"
func (l *Stack6) rearrange4size() {
	defer setLargeFunctions(l)
	if l.head.Size() == 2 {
		return
	}
	var listnodes [4]*stackNode
	for idx := 0; l.head.Size() > 0; idx++ {
		listnodes[idx], l.head = l.head.PopNode()
	}
	for idx := 3; l.tail.Size() > 0; idx-- {
		listnodes[idx], l.tail = l.tail.PopNode()
	}
	l.head = l.head.PushNode(listnodes[1])
	l.head = l.head.PushNode(listnodes[0])
	l.tail = l.tail.PushNode(listnodes[2])
	l.tail = l.tail.PushNode(listnodes[3])
}

func setLargeFunctions(l *Stack6) {
	fmt.Println("set large functions")
	l.currentPopL = largePopLeft
	l.currentPushL = largePushLeft
	l.currentPopR = largePopRight
	l.currentPushR = largePushRight
}

func largePushLeft(l *Stack6, datum any) {
	l.head = l.head.Push(datum)
	if l.initiateTransferCriteria() {
		l.transfer()
	}
}

func largePushRight(l *Stack6, datum any) {
	l.tail = l.tail.Push(datum)
	if l.initiateTransferCriteria() {
		l.transfer()
	}
}

func largePopLeft(l *Stack6) any {
	var datum any
	datum, l.head = l.head.Pop()

	if l.initiateTransferCriteria() {
		l.transfer()
	}

	if l.tail.Size()+l.head.Size() == 4 {
		l.setSmallFunctions()
	}

	return datum
}

func largePopRight(l *Stack6) any {
	var datum any
	datum, l.tail = l.tail.Pop()

	if l.initiateTransferCriteria() {
		l.transfer()
	}

	if l.tail.Size()+l.head.Size() == 4 {
		l.setSmallFunctions()
	}

	return datum
}

func (l *Stack6) transfer() {
	m := l.head.Size()
	n := l.tail.Size()
	B := &(l.head)
	S := &(l.tail)
	if m < n {
		B, S = S, B
	}
	m = (*S).Size()
	k := (*B).Size() - 3*m
	fmt.Printf("rearranging initiated, m = %d, k = %d\n", m, k)
	l.Print(os.Stdout)

	var auxB, auxS, newB, newS *stack
	var datum any

	b, s := (*B).stk, (*S).stk

	steps := 0
	// (a) Reverse 2*m+k-1 items from B to auxB
	// (b) Reverse all items on S to auxS. S has size m
	// First reverse m items of B to auxB, all items of S to auxS
	// Leave *B, *S unchanged, use *stackNode pointers to get at data
	var i int
	for i = 0; i < m; i++ {
		auxB = auxB.Push(b.data)
		b = b.next

		auxS = auxS.Push(s.data)
		s = s.next

		steps++
	}

	// Second, reverse the rest of 2m+k-1 items of B to auxB
	c := 0
	for ; i < 2*m+k-1; i++ {
		auxB = auxB.Push(b.data)
		b = b.next
		c++
	}
	steps += c / 2

	// size(auxB) is 2m+k-1 here.
	// size(B) here is (3m+k)-(2m+k-1) = m+1
	// size(auxS) is m
	// size(B) + size(auxS) = m+1 + m = 2m+1
	// The problem is that k is either 1, 2, or 3,
	// so 2m+k-1 can be 2m, 2m+1, 2m+2, so it's not as easy to do
	// for-loops over section 5's parts (c), (d) and (e)
	// Also note that (e) has to happen after (d) - you can't really
	// do them "concurrently", otherwise the stack order gets messed up.

	for i = 0; i < m+1; i++ {

		// (c) Reverse auxB on to newB
		datum, auxB = auxB.Pop()
		newB = newB.Push(datum)

		// (d) Reverse what's left of B on to newS
		newS = newS.Push(b.data)
		b = b.next
		steps++
	}

	// (c) finish reversing auxB on to newB
	for ; i < 2*m+k-1; i++ {
		datum, auxB = auxB.Pop()
		newB = newB.Push(datum)
	}

	// (e) reverse auxS on to newS
	for i = 0; i < m; i++ {
		datum, auxS = auxS.Pop()
		newS = newS.Push(datum)
	}

	// not correct: *B, *S start out with some operation count
	l.tempStackOps += auxB.Operations() + auxS.Operations() + (*B).Operations() + (*S).Operations()

	*B, *S = newB, newS

	l.Print(os.Stdout)
}

func (l *Stack6) explicittransfer() {
	m := l.head.Size()
	n := l.tail.Size()
	B := &(l.head)
	S := &(l.tail)
	if m < n {
		B, S = S, B
	}
	m = (*S).Size()
	k := (*B).Size() - 3*m
	fmt.Printf("rearranging initiated, m = %d, k = %d\n", m, k)

	var auxB, auxS, newB, newS *stack
	var datum any

	localOps := 0
	// (a) Reverse 2*m+k-1 items from B to auxB
	for i := 0; i < 2*m+k-1; i++ {
		datum, *B = (*B).Pop()
		auxB = auxB.Push(datum)
		localOps += 2
	}
	// (b) Reverse all items on S to auxS
	for datum, *S = (*S).Pop(); datum != nil; datum, *S = (*S).Pop() {
		auxS = auxS.Push(datum)
		localOps += 2
	}
	// (c) Reverse auxB on to newB
	for datum, auxB = auxB.Pop(); datum != nil; datum, auxB = auxB.Pop() {
		newB = newB.Push(datum)
		localOps += 2
	}
	// (d) Reverse B on to newS
	for datum, *B = (*B).Pop(); datum != nil; datum, *B = (*B).Pop() {
		newS = newS.Push(datum)
		localOps += 2
	}
	// (e) Reverse auxS on to newS
	for datum, auxS = auxS.Pop(); datum != nil; datum, auxS = auxS.Pop() {
		newS = newS.Push(datum)
		localOps += 2
	}

	fmt.Printf("%d local stack operations\n", localOps)

	l.tempStackOps += auxB.Operations() + auxS.Operations() + (*B).Operations() + (*S).Operations()

	*B, *S = newB, newS
}

func (l *Stack6) initiateTransferCriteria() bool {
	m := l.head.Size()
	n := l.tail.Size()
	if m > n {
		m, n = n, m
	}
	// m <= n at this point
	fmt.Printf("m %d, n %d, 3*m >= n: %v\n", m, n, 3*m >= n)
	return !(m > 0 && n > 0 && 3*m >= n)
}

func (l *Stack6) setSmallFunctions() {
	fmt.Println("set small functions")
	l.currentPopL = smallPopLeft
	l.currentPushL = smallPushLeft
	l.currentPopR = smallPopRight
	l.currentPushR = smallPushRight
}
