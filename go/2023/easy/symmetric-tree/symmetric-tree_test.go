package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestSYTRecursiveSubTest
	1a)Test improved
	    go test -run TestSYTIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestSYTRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := sytRecursive(testCase.tree)
			if result != testCase.result {
				t.Errorf("SYT %v produces %t, got %t", *testCase.tree, testCase.result, result)
			}
		})
	}
}

func TestSYTIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := sytIterative(testCase.tree)
			if result != testCase.result {
				t.Errorf("SYT %v produces %t, got %t", *testCase.tree, testCase.result, result)
			}
		})
	}
}
