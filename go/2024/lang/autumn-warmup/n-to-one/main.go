package main

import (
	"fmt"
	"sync"
)

func main() {
	merged := mergeChannels(produce())

	for v := range merged {
		fmt.Println(v)
	}
}

const N = 5

func mergeChannels(in []chan int) chan int {
	var wg = &sync.WaitGroup{}
	out := make(chan int)

	for _, v := range in {
		wg.Add(1)

		go consumeChan(v, out, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func consumeChan(ch, res chan int, wg *sync.WaitGroup) {
	for v := range ch {
		res <- v
	}

	wg.Done()
}

func produce() []chan int {
	var res []chan int

	for i := 0; i < N; i++ {
		res = append(res, make(chan int))
	}

	for _, ch := range res {
		go func(ch chan int) {
			for v := range 10 {
				ch <- v
			}
			close(ch)
		}(ch)
	}

	return res
}
