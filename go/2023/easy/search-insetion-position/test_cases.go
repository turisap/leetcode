package main

var testCases = []struct {
	name   string
	nums   []int
	target int
	result int
}{
	{"nums = [1,3,5,6], target = 5", []int{1, 3, 5, 6}, 5, 2},
	{"nums = [1,3,5,6], target = 2", []int{1, 3, 5, 6}, 2, 1},
	{"nums = [1,3,5,6], target = 7", []int{1, 3, 5, 6}, 7, 4},
	{"nums = [2,3,5,6], target = 1", []int{1, 3, 5, 6}, 1, 0},
	{"nums = [1,3,5,6], target = 0", []int{1, 3, 5, 6}, 0, 0},
}
