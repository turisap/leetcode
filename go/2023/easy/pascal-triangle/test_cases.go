package main

var testCases = []struct {
	name   string
	n      int
	result [][]int
}{
	{"[1]", 1, [][]int{{1}}},
	{"[3]", 3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
	{"[5]", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
}
