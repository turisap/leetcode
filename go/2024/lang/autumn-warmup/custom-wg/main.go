package main

import (
	"fmt"
	"time"
)

type WG struct {
	tr chan struct{}
}

func NewWG(n int) *WG {
	return &WG{
		tr: make(chan struct{}, n),
	}
}

func (wg *WG) Add(N int) {
	for _ = range N {
		wg.tr <- struct{}{}
	}
}

func (wg *WG) Done() {
	<-wg.tr
}

func (wg *WG) GetCap() int {
	return cap(wg.tr) - len(wg.tr)
}

func (wg *WG) Wait() {
	for {
		time.Sleep(time.Millisecond * 100)
		if len(wg.tr) == 0 {
			return
		}
	}
}

func main() {
	wg := NewWG(2)

	go func() {
		time.Sleep(time.Second)
		wg.Done()
		wg.Done()
		wg.Done()
		wg.Done()
	}()

	wg.Add(1)
	fmt.Println("Cap", wg.GetCap())
	wg.Add(1)
	fmt.Println("Cap", wg.GetCap())
	wg.Add(1)
	fmt.Println("Cap", wg.GetCap())
	wg.Add(1)
	fmt.Println("Cap", wg.GetCap())

	wg.Wait()
	fmt.Println("exit")
}
