package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DoneCh chan os.Signal

const P = 10
const C = 5

func main() {
	dataChan := make(chan int)
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM)
	defer func() {
		fmt.Println("exit main")
	}()

	for v := range P {
		go Producer(dataChan, done, v)
	}

	for v := range C {
		go Consume(dataChan, v)
	}

	signal.Notify(done)
	select {
	case <-done:
		time.Sleep(time.Second)
		return
	}
}

func Consume(ch <-chan int, i int) {
	defer func() {
		fmt.Printf("exit %d consumer\n", i)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func Producer(ch chan<- int, done DoneCh, i int) {
	ticker := time.NewTicker(time.Millisecond * 100)
	var c int
	defer func() {
		fmt.Println("exit producer", i)
	}()

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
