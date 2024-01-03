package main

import "testing"

/*

	1a)Test Iterative
	    go test -run TestPLLIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestPLLIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := pllIterative(testCase.head)
			if result != testCase.result {
				t.Errorf("PLL %v produces %t, got %t", testCase.head, testCase.result, result)
			}
		})
	}
}
