package main

import (
	"context"
	"sync"
	"time"
)

type TtlCache struct {
	values   *sync.Map
	ctx      context.Context
	cancel   context.CancelFunc
	heap     *TTLHeap
	valueSet chan struct{}
}

type Value struct {
	value string
	key   string
	ttl   time.Time
}

func NewTtlCache() *TtlCache {
	ctx, cancel := context.WithCancel(context.Background())
	values := &sync.Map{}
	valueSet := make(chan struct{}, 1)
	heap := NewTTLHeap()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				v, ok := heap.Peek()
				if !ok {
					select {
					case <-ctx.Done():
						return
					case <-valueSet:
						continue
					}
				}

				if v.ttl.After(time.Now()) {
					t := time.NewTimer(v.ttl.Sub(time.Now()))

					select {
					case <-ctx.Done():
						t.Stop()
						return
					case <-valueSet:
						t.Stop()
						continue
					case <-t.C:
						t.Stop()
						continue
					}
				}

				if v.ttl.Before(time.Now()) {
					v, _ = heap.Pop()
					values.CompareAndDelete(v.key, v)
				}

			}
		}
		// ticker := time.NewTicker(time.Second)
		// defer ticker.Stop()
		//
		// for {
		// 	select {
		// 	case <-ticker.C:
		// 		values.Range(func(key, value interface{}) bool {
		//
		// 			cached, casted := value.(Value)
		// 			if !casted {
		// 				return true
		// 			}
		//
		// 			if time.Now().After(cached.ttl) && !cached.ttl.IsZero() {
		// 				values.CompareAndDelete(key, cached)
		// 			}
		//
		// 			return true
		// 		})
		// 	case <-ctx.Done():
		// 		return
		// 	}
		// }
	}()

	return &TtlCache{
		ctx:      ctx,
		cancel:   cancel,
		values:   values,
		heap:     heap,
		valueSet: valueSet,
	}
}

func (c *TtlCache) Set(key string, value string, ttl time.Duration) {
	var expires time.Time
	if ttl == 0 {
		expires = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	} else {
		expires = time.Now().Add(ttl)
	}
	v := Value{value: value, ttl: expires, key: key}

	c.values.Store(key, v)

	if ttl != 0 {
		c.heap.Push(v)
		select {
		case c.valueSet <- struct{}{}:
			return
		default:
			return
		}
	}
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
