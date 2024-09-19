package main

import "fmt"

func pipe(in chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for msg := range in {
			out <- msg * msg
		}
	}()

	return out
}

func main() {
	in := make(chan int)

	go func() {
		defer close(in)
		for _, v := range []int{1, 4, 6} {
			in <- v
		}
	}()

	out := pipe(in)

	for msg := range out {
		fmt.Println(msg)
	}
}
