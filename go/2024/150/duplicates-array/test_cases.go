package main

var testCases = []struct {
	name   string
	slice  []int
	k      int
	result bool
}{
	{"simple valid", []int{1, 2, 3, 1}, 3, true},
	{"simple valid2", []int{1, 0, 1, 1}, 1, true},
	{"simple invalid", []int{1, 2, 3, 1, 2, 3}, 2, false},
}
