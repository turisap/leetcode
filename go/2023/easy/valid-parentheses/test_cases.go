package main

var testCases = []struct {
	name   string
	result bool
}{
	{"()", true},
	{"{}", true},
	{"[]", true},
	{"[]()", true},
	{"[({})]", true},
	{"(}", false},
	{"}", false},
	{"{})", false},
	{"((", false},
	{"))", false},
}
