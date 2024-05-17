package main

import (
	"fmt"
	"sync"
)

func mergeChans[T interface{}](input []chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup

	go func() {
		for _, v := range input {
			wg.Add(1)
			go func(src chan T) {
				defer wg.Done()
				for msg := range src {
					out <- msg
				}
			}(v)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	input := []chan string{}

	for i := 0; i < 3; i++ {
		ch := make(chan string)
		input = append(input, ch)

		go func(wn int) {
			defer close(ch)
			for c := 0; c < 5; c++ {
				ch <- fmt.Sprintf("w:%d, msg:%d", wn, c)
			}
		}(i)
	}

	out := mergeChans(input)

	for msg := range out {
		fmt.Println(msg)
	}
}
