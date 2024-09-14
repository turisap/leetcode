package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type DoneCh chan os.Signal

const P = 10
const C = 5

func main() {
	dataChan := make(chan int)
	done := make(DoneCh)
	var wg sync.WaitGroup
	defer func() {
		fmt.Println("exit main")
	}()

	for v := range P {
		wg.Add(1)
		go Producer(dataChan, &wg, done, v)
	}

	for v := range C {
		go Consume(dataChan, v)
	}

	time.Sleep(time.Second * 2)
	close(done)

	wg.Wait()
	close(dataChan)
	time.Sleep(time.Second * 2)
}

func Consume(ch <-chan int, i int) {
	defer func() {
		fmt.Printf("exit %d consumer\n", i)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func Producer(ch chan<- int, wg *sync.WaitGroup, done DoneCh, i int) {
	defer wg.Done()
	defer func() {
		fmt.Println("exit producer", i)
	}()

	ticker := time.NewTicker(time.Millisecond * 100)
	var c int

	for {
		select {
		case <-ticker.C:
			c++
			ch <- c
		case <-done:
			return
		}
	}
}
