package main

import "testing"

/*
	1) Test Naive
	    go test -run TestMJENaiveSubTest
	1a)Test Improved
	    go test -run TestMJEImprovedSubTest
	1b)Test Voting
	    go test -run TestMJEVotingSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestMJENaiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := mjeNaive(testCase.input)
			if result != testCase.result {
				t.Errorf("MJE %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}

func TestMJEImprovedSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := mjeImproved(testCase.input)
			if result != testCase.result {
				t.Errorf("MJE %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}

func TestMJEVotingSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := mjeVoting(testCase.input)
			if result != testCase.result {
				t.Errorf("MJE %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}
