package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestIBTRecursiveSubTest
	1a)Test Iterative
	    go test -run TestIBTIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestIBTRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := ibtRecursive(testCase.result)
			if result != testCase.result {
				t.Errorf("IBT %v produces %v, got %v", testCase.head, testCase.result, result)
			}
		})
	}
}

func TestIBTIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := ibtIterative(testCase.head)
			if result != testCase.result {
				t.Errorf("IBT %v produces %v, got %v", testCase.head, testCase.result, result)
			}
		})
	}
}
