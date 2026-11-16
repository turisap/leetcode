package main

import "errors"

var ErrQueueFull = errors.New("queue is full")

type Queue struct {
	valChan chan int
}

func NewQueue(size int) *Queue {
	return &Queue{
		valChan: make(chan int, size),
	}
}

func (q *Queue) Push(val int) error {
	select {
	case q.valChan <- val:
		return nil
	default:
		return ErrQueueFull
	}
}

func (q *Queue) Pop() int {
	select {
	case val := <-q.valChan:
		return val
	default:
		return -1
	}
}
