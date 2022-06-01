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

I could not solve this, so I searched for an answer.
The answer I found was the no such algorithm is currently known,
but that a 6-stack algorithm exists.
Need a reference - that stackexchange.com comment

As near as I can tell, the 6-stack algorithm is described in

#### Real-Time Deques, Multihead Turing Machines and Purely Functional Programming

by Tyng-Ruey Chuang and Benjamin Goldberg

which appeared in the proceedings of an ACM function programming conference
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
4. Write a `Dequeue` implementation that
[Code](fdq/halfstack.go)

## Dequeue testing environment
