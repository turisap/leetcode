package main

import (
	"errors"
	"sync"
)

var ErrQueueFull = errors.New("queue is full")

type Queue struct {
	vals  []int
	head  int
	count int
	mu    sync.RWMutex
}

func NewQueue(size int) *Queue {
	return &Queue{
		vals:  make([]int, size),
		head:  0,
		count: 0,
		mu:    sync.RWMutex{},
	}
}

func (q *Queue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.count == 0 {
		return -1
	}

	popped := q.vals[q.head]

	q.count--
	q.head++
	q.head %= len(q.vals)

	return popped
}

func (q *Queue) Peek() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.count == 0 {
		return -1
	}

	return q.vals[q.head]
}

func (q *Queue) Push(val int) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.vals) == q.count {
		return ErrQueueFull
	}

	idx := (q.count + q.head) % len(q.vals)
	q.vals[idx] = val
	q.count++

	return nil
}

// [_,_]
// [5,_]
// [5,10]
// [_,10]
