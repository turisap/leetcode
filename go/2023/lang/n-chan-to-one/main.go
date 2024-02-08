package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func mergeWorker(ctx context.Context, in, out chan string, counter *atomic.Int64) {
	for msg := range in {
		out <- msg
		time.Sleep(200 * time.Millisecond)

		select {
		case <-ctx.Done():
			fmt.Printf("[mergeWorker ]: terminating due to context cancel\n")
		default:

		}
	}

	v := counter.Add(1)

	if v == 3 {
		close(out)
	}
}

func pushWorker(ctx context.Context, in chan string, id int) {
worker:
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("[pushWorker #%d]: terminating due to context cancel\n", id)
			break worker
		default:
			in <- fmt.Sprintf("[%d]:%d", id, i)
		}
	}
	close(in)
}

func ncoIterative(ctx context.Context, counter *atomic.Int64, args ...chan string) <-chan string {
	resChan := make(chan string)

	for _, arg := range args {
		go mergeWorker(ctx, arg, resChan, counter)
	}

	return resChan
}

func main() {
	counter := &atomic.Int64{}
	channels := []chan string{make(chan string), make(chan string), make(chan string)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for i, ch := range channels {
		go pushWorker(ctx, ch, i)
	}

	resCh := ncoIterative(ctx, counter, channels...)

main:
	for msg := range resCh {
		fmt.Println(msg)

		select {
		case <-ctx.Done():
			fmt.Println("terminating main due to ctx")
			break main
		default:

		}
	}

	time.Sleep(time.Second)
}
