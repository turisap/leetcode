package main

var testCases = []struct {
	name   string
	strs   []string
	result string
}{
	{"fl", []string{"flower", "flow", "flight"}, "fl"},
	{"none", []string{"dog", "racecar", "car"}, ""},
	{"a", []string{"ab", "a"}, "a"},
	{"first for len 1", []string{"first"}, "first"},
	{"compare for sorted", []string{
		"first", "firsty", "firstmanonthemooon", "fist", "fillip", "finica",
		"fififi", "ffffff", "forage", "fillipmoris", "foster", "f"}, "f"},
}
