package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestDARRecursiveSubTest
	1a)Test Iterative
	    go test -run TestDARIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestDARRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := darSimple(testCase.slice, testCase.k)
			if result != testCase.result {
				t.Errorf("DAR %v produces %t, got %t", testCase.slice, testCase.result, result)
			}
		})
	}
}

func TestDARIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := darMapPointers(testCase.slice, testCase.k)
			if result != testCase.result {
				t.Errorf("DAR %v produces %t, got %t", testCase.slice, testCase.result, result)
			}
		})
	}
}
