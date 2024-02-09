package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(in chan int) {
		defer close(in)

		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}(ch1)

	go func(in, out chan int) {
		for i := 0; i < 10; i++ {
			msg := <-ch1
			ch2 <- msg
		}

		close(out)
	}(ch1, ch2)

	for v := range ch2 {
		fmt.Println("msg", v)
	}
}
