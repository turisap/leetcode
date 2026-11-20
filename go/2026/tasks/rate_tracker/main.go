package main

import (
	"context"
	"sync/atomic"
	"time"
)

type Monitor interface {
	SendRate(int)
}

type Handler struct {
	counter atomic.Int64
}

func (h *Handler) Handle() {
	// Some work
	// Calculate rate of this function
	h.counter.Add(1)
}

func (h *Handler) LogRate(ctx context.Context, monitor Monitor, d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	ratesCh := make(chan int, 10)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				rate, ok := <-ratesCh
				if !ok {
					return
				}

				monitor.SendRate(rate)
			}
		}
	}()

	defer close(ratesCh)
	for {
		select {
		case <-ticker.C:
			rate := h.counter.Swap(0)
			select {
			case <-ctx.Done():
				return
			case ratesCh <- int(rate):
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}
