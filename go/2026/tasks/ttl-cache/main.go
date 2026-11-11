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

					if time.Now().After(cached.ttl) && !cached.ttl.IsZero() {
						values.CompareAndDelete(key, cached)
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
	var expires time.Time
	if ttl == 0 {
		expires = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	} else {
		expires = time.Now().Add(ttl)
	}
	c.values.Store(key, Value{value, expires})
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

	if !cached.ttl.After(time.Now()) {
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
