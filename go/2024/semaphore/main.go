package main

import (
	"fmt"
	"time"
)

type sema chan struct{}

func NewSema(k int) sema {
	return make(sema, k)
}

func (s sema) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}
func (s sema) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	s := NewSema(4)

	s.Acquire(2)
	s.Acquire(2)

	go func() {
		time.Sleep(time.Second)
		s.Release(2)
	}()

	fmt.Println("all acquired")
	s.Acquire(2)
	fmt.Println("acquire 2 more")
}
