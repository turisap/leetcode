package main

import (
	"context"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	processCtx context.Context
	cancelFunc context.CancelFunc
	ticker     *time.Ticker
}

func NewRateLimiter(n int) *RateLimiter {
	tokens := make(chan struct{}, n)
	// can start full if u want that
	// for i := 0; i < n; i++ {
	// 	tokens <- struct{}{}
	// }

	ctx, cancel := context.WithCancel(context.Background())
	t := time.NewTicker(time.Duration(1000_000_000 / n))

	go func() {
		defer t.Stop()

		for _ = range t.C {
			select {
			case tokens <- struct{}{}:
			case <-ctx.Done():
				return
			default:
			}
		}
	}()

	return &RateLimiter{
		tokens:     tokens,
		cancelFunc: cancel,
		processCtx: ctx,
		ticker:     t,
	}
}

func (r *RateLimiter) CanTake() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Take() {
	<-r.tokens
}

func (r *RateLimiter) Stop() {
	r.cancelFunc()
}
