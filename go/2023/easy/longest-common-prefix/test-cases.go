package main

var testCases = []struct {
	name   string
	strs   []string
	result string
}{
	{"fl", []string{"flower", "flow", "flight"}, "fl"},
	{"none", []string{"dog", "racecar", "car"}, ""},
	{"a", []string{"ab", "a"}, "a"},
}
