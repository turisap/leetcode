package main

import (
	"context"
	"sync"
)

func Merge(channels ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, ch := range channels {
			go func() {
				defer wg.Done()

				if ch == nil {
					return
				}

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
func MergeCancelable(ctx context.Context, channels ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, ch := range channels {
			go func() {
				defer wg.Done()

				if ch == nil {
					return
				}

				for {
					select {
					case <-ctx.Done():
						return
					case n, ok := <-ch:
						if !ok {
							return
						}

						select {
						case merged <- n:
							continue
						case <-ctx.Done():
							return
						}

					}
				}
			}()
		}

		wg.Wait()
		close(merged)
	}()

	return merged
}
