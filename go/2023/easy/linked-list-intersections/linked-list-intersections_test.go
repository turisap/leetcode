package main

import "testing"

/*
	1) Test HashTable
	    go test -run TestLLIHashTableSubTest
	1a)Test TwoPointers
	    go test -run TestLLITwoPointersSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestLLIHashTableSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lliHashTable(testCase.head1, testCase.head2)
			if result != testCase.result {
				t.Errorf("LLI %s produces %v, got %v", testCase.name, testCase.result, result)
			}
		})
	}
}

func TestLLITwoPointersPlainSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lliTwoPointersPlain(testCase.head1, testCase.head2)
			if result != testCase.result {
				t.Errorf("LLI %s produces %v, got %v", testCase.name, testCase.result, result)
			}
		})
	}
}

func TestLLITwoPointersImprovedSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := lliTwoPointersImproved(testCase.head1, testCase.head2)
			if result != testCase.result {
				t.Errorf("LLI %s produces %v, got %v", testCase.name, testCase.result, result)
			}
		})
	}
}
