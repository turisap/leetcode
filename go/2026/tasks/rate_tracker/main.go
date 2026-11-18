package main

import (
	"sync"
	"time"
)

type Monitor interface {
	SendRate(int)
}

type Handler struct {
	counter float64
	mu      sync.Mutex
}

func (h *Handler) Handle() {
	// Some work
	// Calculate rate of this function
	h.mu.Lock()
	defer h.mu.Unlock()
	h.counter++
}

func (h *Handler) LogRate(monitor Monitor, d time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()

	rate := h.counter / d.Seconds()

	monitor.SendRate(int(rate))
}
