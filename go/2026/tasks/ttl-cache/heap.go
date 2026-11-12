package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type TTLHeap struct {
	heap []Value
	mu   *sync.RWMutex
}

type valueHeap []Value

func (h valueHeap) Len() int           { return len(h) }
func (h valueHeap) Less(i, j int) bool { return h[i].ttl.Before(h[j].ttl) }
func (h valueHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *valueHeap) Push(x any) {
	*h = append(*h, x.(Value))
}

func (h *valueHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// NewTTLHeap initializes a ready-to-use TTLHeap.
func NewTTLHeap() *TTLHeap {
	return &TTLHeap{
		heap: make([]Value, 0),
		mu:   &sync.RWMutex{},
	}
}

// Push safely adds a Value to the heap.
func (t *TTLHeap) Push(v Value) {
	t.mu.Lock()
	defer t.mu.Unlock()

	h := valueHeap(t.heap)
	heap.Push(&h, v)
	t.heap = h
}

// Pop safely removes and returns the Value with the earliest TTL.
func (t *TTLHeap) Pop() (Value, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.heap) == 0 {
		return Value{}, false
	}

	h := valueHeap(t.heap)
	item := heap.Pop(&h).(Value)
	t.heap = h
	return item, true
}

// Peek safely returns the Value with the earliest TTL without removing it.
func (t *TTLHeap) Peek() (Value, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if len(t.heap) == 0 {
		return Value{}, false
	}

	return t.heap[0], true
}

// Len safely returns the current number of elements in the heap.
func (t *TTLHeap) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return len(t.heap)
}

func main() {
	now := time.Now()
	th := NewTTLHeap()

	th.Push(Value{value: "expires in 10m", ttl: now.Add(10 * time.Minute)})
	th.Push(Value{value: "expires in 10s", ttl: now.Add(10 * time.Second)})
	th.Push(Value{value: "expires in 1m", ttl: now.Add(1 * time.Minute)})

	if next, ok := th.Peek(); ok {
		fmt.Printf("Next to expire: %s\n\n", next.value)
	}

	for th.Len() > 0 {
		if item, ok := th.Pop(); ok {
			fmt.Printf("Popped: %-15s | TTL: %s\n", item.value, item.ttl.Format(time.TimeOnly))
		}
	}
}
