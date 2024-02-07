package main

var testCases = []struct {
	name   string
	input  string
	result bool
}{
	{"Even palindrome", "anna", true},
	{"Uneven palindrome", "argra", true},
	{"Various cases false palindrome", "aRgra", false},
	{"Not palindrome", "test", false},
	{"Empty string", "", false},
	{"Long string", "akaabakolokabaaka", true},
}
