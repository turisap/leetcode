package main

import "fmt"

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

func main() {
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
