package main

import "testing"

/*
	1) Test TwoPointers
	    go test -run TestLHSTwoPointersSubTest
	1a)Test Iterative
	    go test -run TestLHSIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestLHSTwoPointersSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lhsTwoPointers(testCase.head)
			if result != testCase.result {
				t.Errorf("LHS %v produces %t, got %t", testCase.head, testCase.result, result)
			}
		})
	}
}

func TestLHSIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lhsIterative(testCase.head)
			if result != testCase.result {
				t.Errorf("LHS %v produces %t, got %t", testCase.head, testCase.result, result)
			}
		})
	}
}
