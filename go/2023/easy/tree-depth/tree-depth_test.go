package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestTRDRecursiveSubTest
	1a)Test Iterative
	    go test -run TestTRDIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestTRDRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := trdRecursive(testCase.input)
			if result != testCase.result {
				t.Errorf("TRD %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}

func TestTRDIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := trdIterative(testCase.input)
			if result != testCase.result {
				t.Errorf("TRD %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}
