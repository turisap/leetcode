package main

import (
	"context"
	"sync"
	"time"
)

type TtlCache struct {
	values *sync.Map
	ctx    context.Context
	cancel context.CancelFunc
}

type Value struct {
	value string
	ttl   time.Time
}

func NewTtlCache() *TtlCache {
	ctx, cancel := context.WithCancel(context.Background())
	values := &sync.Map{}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				values.Range(func(key, value interface{}) bool {

					cached, casted := value.(Value)
					if !casted {
						return true
					}

					if time.Now().After(cached.ttl) && cached.ttl.IsZero() {
						values.Delete(key)
					}

					return true
				})
			case <-ctx.Done():
				return
			}
		}
	}()

	return &TtlCache{
		ctx:    ctx,
		cancel: cancel,
		values: values,
	}
}

func (c *TtlCache) Set(key string, value string, ttl time.Duration) {
	c.values.Store(key, Value{value, time.Now().Add(ttl)})
}

func (c *TtlCache) Get(key string) (string, bool) {
	val, ok := c.values.Load(key)
	if !ok {
		return "", false
	}

	cached, casted := val.(Value)
	if !casted {
		return "", false
	}

	return cached.value, true
}

func (c *TtlCache) Delete(key string) {
	c.values.Delete(key)
}

func (c *TtlCache) Stop() {
	c.cancel()
}
