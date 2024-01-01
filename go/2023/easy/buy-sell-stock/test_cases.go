package main

var testCases = []struct {
	name   string
	skip   bool
	input  []int
	result int
}{
	{"get 5", false, []int{7, 1, 5, 3, 6, 4}, 5},
	{"get 0", false, []int{7, 6, 4, 3, 1}, 0},
}
