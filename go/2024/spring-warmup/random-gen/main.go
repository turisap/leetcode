package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			out <- r.Int()
		}
		close(out)
	}()

	return out
}

func main() {
	g := generator(5)

	for v := range g {
		fmt.Println(v)
	}
}
