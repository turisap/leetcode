package main

import "testing"

/*
	1) Test Naive
	    go test -v -run TestBSSNaiveSubTest
	1a)Test Improved
	    go test -v -run TestBSSImprovedSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestBSSNaiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.skip {
				t.Skip()
			}
			result := bssNaive(testCase.input)
			if result != testCase.result {
				t.Errorf("BSS %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}

func TestBSSImprovedSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.skip {
				t.Skip()
			}
			result := bssImproved(testCase.input)
			if result != testCase.result {
				t.Errorf("BSS %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}
