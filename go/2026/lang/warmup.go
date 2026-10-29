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

func main() {
	// s := make([]int, 1, 2) // then: make([]int, 1, 1)
	s := make([]int, 1, 1) // then: make([]int, 1, 1)
	f2(&s)
	fmt.Println(s)
}
