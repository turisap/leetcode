package main

import (
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var done = sync.Map{}
var mu = &sync.Mutex{}
var result = make([]string, 0)

func Crawl(url string, depth int, fetcher Fetcher) ([]string, error) {
	if depth <= 0 {
		return nil, nil
	}

	_, loaded := done.LoadOrStore(url, true)
	if loaded {
		return nil, nil
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		return nil, err
	}

	done.Store(url, true)

	mu.Lock()
	result = append(result, body)
	mu.Unlock()

	wg := sync.WaitGroup{}

	for _, u := range urls {
		if _, ok := done.Load(u); ok {
			continue
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			_, nextErr := Crawl(u, depth-1, fetcher)
			done.Store(u, true)

			if nextErr != nil {
				return
			}
		}()
	}

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	return result, nil
}
