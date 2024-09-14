package main

import "fmt"

func getSrcChan() <-chan int {
	out := make(chan int)
	go func() {
		for v := range 10 {
			out <- v
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func dv(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v / 2
		}
		close(out)
	}()

	return out
}

func main() {

	res := dv(sq(getSrcChan()))

	for m := range res {
		fmt.Println(m)
	}
}
