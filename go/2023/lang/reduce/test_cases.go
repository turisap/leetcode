package main

var intCases = []struct {
	name   string
	input  []int
	result any
}{
	{"Array of numbers", []int{1, 2, 3}, 6},
	{"Array of one element (zero)", []int{0}, 0},
	{"Array of one element (non-zero)", []int{5}, 5},
	{"Array of negative numbers", []int{1, -2, -3}, -4},
}

var stringCases = []struct {
	name   string
	input  []string
	result any
}{
	{"Empty Array", []string{}, ""},
	{"Array of strings", []string{"hello", "world", "!"}, "helloworld!"},
}
