package main

var testCases = []struct {
	name   string
	s      []int
	result int
}{
	{"returns 3 for [1,1,1,2]", []int{1, 1, 1, 2}, 3},
	{"returns 5 for [1,1,1,2,2,3]", []int{1, 1, 1, 2, 2, 3}, 5},
	{"returns 7 for [0,0,1,1,1,1,2,3,3]", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	{"returns 6 for [1,1,1,2,2,2,3,3]", []int{1, 1, 1, 2, 2, 2, 3, 3}, 6},
}
