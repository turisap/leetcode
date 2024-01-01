package main

var testCases = []struct {
	name   string
	input  []int
	result int
}{
	{"get 2 for [2,2,1]", []int{2, 2, 1}, 1},
	{"get 1 for [1]", []int{1}, 1},
	{"get 4 for {1,1,2,2,3,4,4}", []int{1, 1, 2, 2, 3, 4, 4}, 3},
}
