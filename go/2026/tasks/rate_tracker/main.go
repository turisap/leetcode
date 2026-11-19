package main

import (
	"context"
	"sync"
	"time"
)

type Monitor interface {
	SendRate(int)
}

type Handler struct {
	counter int
	mu      sync.Mutex
}

func (h *Handler) Handle() {
	// Some work
	// Calculate rate of this function
	h.mu.Lock()
	defer h.mu.Unlock()
	h.counter++
}

func (h *Handler) LogRate(ctx context.Context, monitor Monitor, d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.mu.Lock()
			val := h.counter
			h.counter = 0
			h.mu.Unlock()

			monitor.SendRate(val)
		case <-ctx.Done():
			return
		}
	}
}
