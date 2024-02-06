package main

var testCases = []struct {
	name   string
	s      []int
	item   int
	result []int
}{
	{"Add one integer to a slice", []int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
	{"Add one integer to an empty slice", []int{}, 1, []int{1}},
}

var testMulti = []struct {
	name   string
	s      []int
	args   []int
	result []int
}{
	{"Add a set of integers to an array", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
	{"Add a set of integers to an empty array", []int{}, []int{2, 3, 4}, []int{2, 3, 4}},
	{"Add a set of integers to an array when cap is enough", make([]int, 0, 5), []int{2, 3, 4}, []int{2, 3, 4}},
}
