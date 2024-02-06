package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestRDCRecursiveSubTest
	1a)Test Iterative
	    go test -run TestRDCIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestRDCRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := reduceCustom()
			if result != testCase.result {
				t.Errorf("RDC %s produces %d, got %d", testCase.s, testCase.result, result)
			}
		})
	}
}
