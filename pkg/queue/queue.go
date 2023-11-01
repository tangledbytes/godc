package queue

import (
	"sync/atomic"
)

type Node[T any] struct {
	Next atomic.Pointer[Node[T]]
	Data T
}

type Queue[T any] struct {
	head *atomic.Pointer[Node[T]]
	tail *atomic.Pointer[Node[T]]
	len  atomic.Int64
}

func New[T any]() *Queue[T] {
	sen := &Node[T]{}

	head := &atomic.Pointer[Node[T]]{}
	head.Store(sen)

	tail := &atomic.Pointer[Node[T]]{}
	tail.Store(sen)

	return &Queue[T]{
		head: head,
		tail: tail,
	}
}

func (q *Queue[T]) Len() int64 {
	return q.len.Load()
}

func (q *Queue[T]) Push(data T) {
	new := &Node[T]{Data: data}

	for {
		tail := q.tail.Load()
		next := tail.Next.Load()

		if tail == q.tail.Load() {
			// ideal case - tail.Next is nil
			if next == nil {
				if tail.Next.CompareAndSwap(nil, new) {
					q.tail.CompareAndSwap(tail, new)
					q.len.Add(1)
					return
				}
				// something changed the value of tail.Next
				// while we were reading it
			} else {
				q.tail.CompareAndSwap(tail, next)
			}
		}
		// Something changed the value of tail
		// while we were reading it
		//
		// -- retry
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	var def T

	for {
		head := q.head.Load()
		tail := q.tail.Load()
		next := head.Next.Load()

		// if the head is the same as when we started
		// then we can try to pop
		if head == q.head.Load() {
			// nothing to remove from the queue
			if head == tail {
				// this is the ideal case
				if next == nil {
					return def, false
				}

				// something changed the value of head.Next
				// let tail catch up
				q.tail.CompareAndSwap(tail, next)
			} else {
				if q.head.CompareAndSwap(head, next) {
					q.len.Add(-1)
					return next.Data, true
				}
			}
		}

		// head changed while we were reading it
		// -- retry
	}
}

// Peek returns the next item in the queue without removing it
func (q *Queue[T]) Peek() (T, bool) {
	var def T

	head := q.head.Load()
	next := head.Next.Load()

	if next == nil {
		return def, false
	}

	return next.Data, true
}
