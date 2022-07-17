# Real-time Double-Ended Queues

I found a "daily coding problem":

---

## Daily Coding Problem: Problem #1116 [Hard]

This problem was asked by Google.

A quack is a data structure combining properties of both stacks and queues.
It can be viewed as a list of elements written left to right such that three
operations are possible:

* `push(x)`: add a new item x to the left end of the list
* `pop()`: remove and return the item on the left end of the list
* `pull()`: remove the item on the right end of the list.

Implement a quack using three stacks and `O(1)` additional memory,
so that the amortized time for any push, pop, or pull operation is `O(1)`.

---

It is said that this problem comes from `Algorithms`, 4th Ed,
by Robert Sedgewick and Kevin Wayne:

```
Queue with three stacks. Implement a queue with three stacks so that each queue operation takes a constant (worst-case) number of stack operations. Warning : high degree of difficulty.
```

It's also said that later editions rephrase the question as "a finite number of stacks".

I could not solve this, so I searched for an answer.
The answer I found was the no such algorithm is currently known,
but that a 6-stack algorithm exists.

As near as I can tell, the 6-stack algorithm is described in

#### Real-Time Deques, Multihead Turing Machines and Purely Functional Programming

by Tyng-Ruey Chuang and Benjamin Goldberg

which appeared in the proceedings of an ACM functional programming conference
that occurred in Denmark, 1993.

## What I did

1. Define a `Dequeue` Go interface for a double-ended queue ("dequeue").
A dequeue will meet either meet the definition of a "quack"
or be deformed to meet that definition.
2. Make a `Dequeue` implementation that works: a doubly-linked list,
to ensure that my Go interface is correct. [Code](fdq/dllist.go)
3. Write a traditional 2-stack `Dequeue` implementation.
Chuang and Goldberg discuss this in section 2 of their paper.
I mistakenly believed that a 2-stack-queue was more or less a curiosity,
an example of how 2 linked lists could work as a more capable data structure.
Not the case -
Lisp lists are apparently often used this way
because it's so slow to find the tail of a list.
[Code](fdq/twostack.go)
4. Write a `Dequeue` implementation that implements the algorithm of Section 3
of Chuang and Goldberg.
[Code](fdq/halfstack.go)
5. Write a `Dequeue` implementation of Section 5 of Chuang and Goldberg that
doesn't do all of the algorithm. Section 5 "distributes" the internal stack operations
over a number of Dequeue operations so that it's O(1).
[Code](fdq/sixstack.go)
6. Write a `Dequeue` implementation of Section 5 of Chuang and Goldberg that
does all of the algorithm. Have it count actual stack operations to see if
it's O(1) or not, because Section 5 is really hand wavy about how the algorithm
gets to O(1).

### Handwavy "steps" of Chuang and Goldberg

---

Note that procedures a and b above
can be carried out concurrently,
which takes no more than 2m + 3 steps.
A step costs two units of time and space by
moving one element from B to auxB and moving one element from S to auxS.
Similarly, procedure c can be carried out concurrently with d and e,
taking a total of at most 2m + 3 steps.
In total, 4m + 6 steps are sufficient to complete the transfer process.
Since the transfer process will be distributed evenly over the next m deque operations,
we allocate 6 steps to the deque operation that violates the invariant,
and 4 steps to each of the m deque operations that follow.

---

I think this is vague, and it's incorrect.

"A step costs two units of time and space by
moving one element from B to auxB and moving one element from S to auxS."

`2m+k-1` items get moved from B to auxB during procedure (a).
`m` items get moved from S to auxS during procedure (b).
You could pop S and push that on auxS in the same loop
as popping B and pushing on auxB, but why bother?
2 pops and 2 pushes in each loop iteration are still 2 pops and 2 pushes.
Since only m items exist on S,
you'd have to do more overlapping to get that many pops and pushes done
in a certain number of "steps".

Also, k can be 1, 2, 3, so saying "at most 2m+3 steps" doesn't make sense.
If k = 3,  2m + k - 1 = 2m + 2

Counting up the number of pops and pushes makes more sense:

* Procedure (a) - "Pop and reverse the topmost 2m + k - 1 elements of
stack B into an auxiliary stack auxB". 2m+k-1 pops, 2m+k-1 pushes
* Procedure (b) - "Reverse stack S into an auxiliary auxS", m pops and pushes
* Procedure (c) - "Reverse stack auxB into a new stack newB",  2m+k-1 pops, 2m+k-1 pushes
* Procedure (d) - "Reverse stack B into a new stack newS",
this is what's left of B after (a). 3m+k - (2m + k - 1) = m + 1 pops, m+1 pushes
* Procedure (e) - "Reverse stack auxS onto stack newS", m pops, m pushes

Total pops (same number of pushes):  2m + k - 1 + m + 2m + k - 1 + m + 1  + m = 7m + 2k - 1

Total pops and pushes: 14m + 4k - 2

```
number of stack operations: (14m+4k-2)/m

Stack ops amortised over m dequeue ops: 14 + 4k/m - 2/m

If m = 1 and k = 3, 14+12-2 = 24 stack ops per dequeue op

If m is large and k = 1,

Number of stack ops: 14 + 4/m - 2/m = 14 + 6/m

If m is large, 6/m is vanishingly small, the algorithm will take just more than 14
stack ops per m dequeue ops.
```

|left size|right size|m|k|stack ops per dequeue op|
|:--------|:---------|:---:|:---:|----------------|
|6|  1|2|3|24.00|
|2|  7|2|1|15.00|
|3| 10|3|1|14.67|
|4| 13|4|1|14.50|
|5| 16|5|1|14.40|
|6| 19|6|1|14.33|
|7| 22|7|1|14.29|
|8| 25|8|1|14.25|
|10|31|10|1|14.20|
|20|61|20|1|14.10|
|100|301|100|1|14.02|
|300|901|300|1|14.01|

## Dequeue testing environment

I wrote a semi-interactive program to create,
push and pop (left and right),
print contents of,
and delete dequeues.

### Build

```sh
$ go build .
```

### Run

```sh
$ ./functional_dequeue halfstack
> type
halfstack
> pushl 1
> pushl 2
> pushl 3
> pushl 4
> pushl 5
> print
head (5:5): 5 -> 4 -> 3 -> 2 -> 1 -> 
tail (0:0): 
Dequeue operations 5, stack operations: 5 => 1.000
> popr
pop right: 1
> print
head (2:10): 5 -> 4 -> 
tail (2:4): 2 -> 3 -> 
Dequeue operations 6, stack operations: 14 => 2.333
> new
Available implementations: [dllist halfstack stack6a twostack]
Choose dequeue implementation:
```
