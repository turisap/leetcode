package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nProducers := 20
	nWorkers := 5
	jobs := make(chan string, nProducers)

	for i := 0; i < nProducers; i++ {
		t := time.NewTicker(time.Millisecond * time.Duration(rand.Intn(1000)))
		defer t.Stop()

		go func(pN int) {
			for _ = range t.C {
				jobs <- fmt.Sprintf("P: %d, msg: %d", pN, rand.Intn(10))
			}
		}(i)
	}

	for i := 0; i < nWorkers; i++ {

		go func(wN int) {
			for msg := range jobs {
				fmt.Printf("WN %d, recieve: %s\n", wN, msg)
			}
		}(i)
	}

	select {}
}
