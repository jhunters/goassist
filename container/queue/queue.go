package queue

import (
	"sync"

	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/container/listx"
)

// Queue provides a container as the rule of FIFO(first in first out) manner
type Queue[E any] struct {
	l    *listx.List[*QueueEle[E]]
	lock sync.Mutex

	cond *sync.Cond
}

type QueueEle[E any] struct {
	v E
}

// NewQueue create a new queue
func NewQueue[E any]() *Queue[E] {
	return &Queue[E]{l: listx.NewList[*QueueEle[E]](), cond: sync.NewCond(&sync.Mutex{})}
}

// Enqueue add element to queue
func (q *Queue[E]) Enqueue(e E) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.l.PushBack(&QueueEle[E]{e})
	q.cond.Signal()
}

// Dequeue dequeue element from top
func (q *Queue[E]) Dequeue() E {
	var ret E
	e := q.dequeueEle()
	if e != nil {
		ret = e.v
	}

	return ret
}

// Dequeue dequeue element from top
func (q *Queue[E]) dequeueEle() *QueueEle[E] {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.l.RemoveFront()
}

// Clear all elements
func (q *Queue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.l.Clear()
}

// SubScribe
func (q *Queue[E]) SubScribe(c base.Consumer[E]) {
	go func(q *Queue[E]) {
		for {
			q.cond.L.Lock()
			v := q.dequeueEle()
			if v == nil {
				// no value and wait
				q.cond.Wait()
			} else {
				c(v.v)
			}

			q.cond.L.Unlock()
		}

	}(q)
}
