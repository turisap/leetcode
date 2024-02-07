package main

var testCases = []struct {
	name   string
	a      []int
	b      []int
	result []int
}{
	{"Simple intersection #2", []int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
	{"Simple intersection #1", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
}
