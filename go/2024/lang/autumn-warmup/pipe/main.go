package main

import (
	"fmt"
	"time"
)

type Done chan struct{}

func getSrcChan(done Done) <-chan int {
	out := make(chan int)

	go func() {
		defer func() {
			fmt.Println("exit  out")
			close(out)
		}()

		for v := range 100 {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
	}()

	return out
}

func sq(in <-chan int, done Done) <-chan int {
	out := make(chan int)

	go func() {
		defer func() {
			fmt.Println("exit  sq")
			close(out)
		}()
		for v := range in {
			select {
			case out <- v * v:
			case <-done:
				return
			}
		}
		close(out)
	}()

	return out
}

func dv(in <-chan int, done Done) <-chan int {
	out := make(chan int)

	go func() {
		defer func() {
			fmt.Println("exit  dv")
			close(out)
		}()
		for v := range in {
			select {
			case out <- v * 2:
			case <-done:
				return
			}
		}
		close(out)
	}()

	return out
}

func main() {
	done := make(Done)

	res := dv(sq(getSrcChan(done), done), done)

	go func() {
		time.Sleep(time.Second)
		close(done)
	}()

	for m := range res {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(m)
	}
}
