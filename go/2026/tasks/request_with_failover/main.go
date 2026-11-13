package main

import (
	"context"
	"errors"
	"time"
)

type Client interface {
	Get(ctx context.Context, address string) (string, error)
}

var ErrRequestsFailed = errors.New("task execution failed")

// RequestWithFailover attempts to request a data from available addresses:
// 1. If error, immediately try the address without waiting
// 2. If an address doesn't respond within 500ms, try the next but keep the original request running
// 3. Return the first successful response, or all ErrRequestsFailed if all nodes fail
// 4. Properly handle context cancellation throughout the process
func RequestWithFailover(ctx context.Context, client Client, addresses []string) (string, error) {
	results := make(chan string, len(addresses))
	errCh := make(chan error, len(addresses))
	processCtx, cancel := context.WithCancel(ctx)

	defer cancel()

	for _, address := range addresses {
		t := time.NewTimer(time.Millisecond * 500)
		failed := make(chan struct{})

		go func(pCtx context.Context, a string) {
			resp, err := client.Get(pCtx, a)
			if err != nil {
				errCh <- err
				close(failed)
				return
			}

			results <- resp
		}(processCtx, address)

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-t.C:
			continue
		case result := <-results:
			return result, nil
		case <-failed:
			continue
		}
	}

	for i := 0; i < len(addresses); i++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case result := <-results:
			return result, nil
		case <-errCh:
			continue
		}
	}

	return "", ErrRequestsFailed
}
