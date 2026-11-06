package main

import (
	"sync"
)

type Client interface {
	Get(address string) (string, error)
}

type Cache struct {
	client Client
	res    map[string]*Result
	mu     *sync.Mutex
}

// Don't update signature of NewCache
func NewCache(client Client) *Cache {
	// TODO: Implement
	return &Cache{client: client, mu: new(sync.Mutex), res: make(map[string]*Result)}
}

type Result struct {
	res  string
	err  error
	done chan struct{}
}

// Cache Client.Get result
func (c *Cache) Get(address string) (string, error) {
	c.mu.Lock()
	cached, ok := c.res[address]
	if ok {
		c.mu.Unlock()
		<-cached.done

		return cached.res, cached.err
	}
	result := &Result{
		done: make(chan struct{}),
	}
	c.res[address] = result
	c.mu.Unlock()

	fetched, err := c.client.Get(address)

	c.mu.Lock()
	defer c.mu.Unlock()
	if err != nil {
		c.res[address].err = err
		c.res[address].res = ""
		close(c.res[address].done)

		return "", err
	}

	c.res[address].err = nil
	c.res[address].res = fetched
	close(c.res[address].done)

	return fetched, nil
}
