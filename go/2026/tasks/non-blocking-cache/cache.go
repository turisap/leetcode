package main

import (
	"sync"
	"time"
)

type Client interface {
	Get(address string) (string, error)
}

type Cache struct {
	client Client
	res    map[string]Result
	mu     *sync.Mutex
}

// Don't update signature of NewCache
func NewCache(client Client) *Cache {
	// TODO: Implement
	return &Cache{client: client, mu: new(sync.Mutex), res: make(map[string]Result)}
}

type Result struct {
	progressing bool
	res         string
	err         error
}

// Cache Client.Get result
func (c *Cache) Get(address string) (string, error) {
	// cache lookup
	// if cached and not in progress return cached
	// if cached in progress poll
	// fetch if not cached
	c.mu.Lock()
	defer c.mu.Unlock()
	r, ok := c.res[address]
	// cached
	if ok {
		if r.progressing {
			t := time.NewTicker(time.Millisecond * 100)
			defer t.Stop()

			for {
				select {
				case <-t.C:
					polled := c.res[address]
					if !polled.progressing {
						return polled.res, nil
					}
				}
			}
		}

		if !r.progressing && r.err != nil {
			return "", r.err
		}

		return r.res, nil
	}

	c.res[address] = Result{
		progressing: true,
		res:         "",
	}

	fetched, err := c.client.Get(address)
	if err != nil {
		c.res[address] = Result{
			progressing: false,
			res:         "",
			err:         err,
		}

		return "", err
	}

	c.res[address] = Result{
		progressing: false,
		res:         fetched,
	}

	return fetched, nil
}
