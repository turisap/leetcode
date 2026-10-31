package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type MyErr struct{}

func (*MyErr) Error() string { return "err" }

func do() error {
	var e *MyErr
	return e
}

// func main() { fmt.Println(do() == nil) }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := a[1:3]
// 	b = append(b, 100)
//
// 	fmt.Println(a, b, len(b), cap(b))
// }

func f(s []int) {
	s = append(s, 1)
	s[0] = 9
}

func main1() {
	s := make([]int, 1, 2) // then: make([]int, 1, 1)
	// s := make([]int, 1, 1) // then: make([]int, 1, 1)
	f(s)
	fmt.Println(s)
	fmt.Println(s[0:2])
}

func f2(s *[]int) {
	*s = append(*s, 1)
	(*s)[0] = 9
}

func main2() {
	// s := make([]int, 1, 2) // then: make([]int, 1, 1)
	s := make([]int, 1, 1) // then: make([]int, 1, 1)
	f2(&s)
	fmt.Println(s)
}

func f1() (r int) {
	defer func() { r++ }()
	return 5
}

func g() int {
	x := 1
	defer func() { x++ }()
	return x
}

func h() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

func main3() {
	h()
}

func main5() {
	arr := [3]int{1, 2, 3}
	// slice := []int{1, 2, 3}
	for i, v := range &arr {
		// for i, v := range arr {
		// for i, v := range slice {
		arr[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
	// What changes if arr is a slice? What if you range over &arr?
}

type Base struct{}

func (Base) Name() string    { return "base" }
func (b Base) Hello() string { return "hello " + b.Name() }

type Derived struct{ Base }

func (Derived) Name() string { return "derived" }
func main6() {
	fmt.Println(Derived{}.Hello())
}

func main7() {
	s := "héllo"
	fmt.Println(len(s), len([]rune(s)))
	for i, r := range s {
		fmt.Print(i, ":", string(r), " ")
	}
}

func main8() {
	const (
		a = iota
		b
		_
		c = "x"
		d
		m = "x"
		e = iota
	)
	fmt.Println(a, b, c, d, e)
}

func main9() {
	var a, b interface{} = []int{1}, []int{1}
	fmt.Println(a == b)
}

func main10() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	panic("boom")
}

func main11() {
	e := work()
	fmt.Println(e)
}

func work() (err error) {
	if true {
		_, err := strconv.Atoi("x")
		if err != nil {
			return err // does not compile
			// return err - fix
		}
	}
	return
}

func main12() {
	ch := make(chan int)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()
	fmt.Println(<-ch)
	wg.Wait()
}

func main13() {
	ssec()
}

func ssec() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()
	for i := 0; i < 3; i++ {
		// ii := i
		// go func() { ch <- ii }()
		ch <- i
	}
	close(ch)

	wg.Wait()
}

func first() int {
	ch := make(chan int) // unbuffered
	for i := 0; i < 3; i++ {
		go func() { ch <- i }() // 3 goroutines, each tries to send
	}
	return <-ch // receives exactly ONE value
}

func main() {
	fmt.Println(runtime.NumGoroutine()) // 1
	first()
	time.Sleep(100 * time.Millisecond)  // let goroutines settle
	fmt.Println(runtime.NumGoroutine()) // 3  <- 2 leaked
}
