package main

import "sync"

func Merge(channels ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, ch := range channels {
			go func() {
				defer wg.Done()

				for n := range ch {
					merged <- n
				}
			}()
		}

		wg.Wait()
		close(merged)
	}()

	return merged
}
