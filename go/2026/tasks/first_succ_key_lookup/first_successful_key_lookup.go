package first_succ_key_lookup

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// struct which rules the process (with methods)
// fun-out to N services - config ?? restriction to max addresses at once
// cancel all if get success
// timeouts - config

type Config struct {
	addresses []string
	timeout   time.Duration
}

type Getter struct {
	cfg *Config
	cl  *http.Client
}

var (
	ErrNoSuccess = fmt.Errorf("no success")
)

func NewGetter(cfg *Config) *Getter {
	transport := &http.Transport{
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     0,
		IdleConnTimeout:     time.Second * 3,
	}
	return &Getter{
		cfg: cfg,
		cl: &http.Client{
			Transport: transport,
			Timeout:   cfg.timeout,
		},
	}
}

func (g *Getter) Get() (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res := make(chan string, len(g.cfg.addresses))

	wg := &sync.WaitGroup{}

	for _, add := range g.cfg.addresses {
		wg.Add(1)

		go func(endpoint string) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)

			if err != nil {
				fmt.Println(err)
				return
			}

			resp, err := g.cl.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("non-200 response:", resp.StatusCode)
				return
			}

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			res <- string(b)
			cancel()
		}(add)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	k, ok := <-res
	if !ok {
		return "", ErrNoSuccess
	}

	return k, nil
}
