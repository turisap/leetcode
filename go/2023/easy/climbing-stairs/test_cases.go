package main

var testCases = []struct {
	name   string
	input  int
	result int
}{
	{"1 stair", 1, 1},
	{"2 stairs", 2, 2},
	{"3 stairs", 3, 3},
	{"44 stairs", 44, 1134903170},
}
