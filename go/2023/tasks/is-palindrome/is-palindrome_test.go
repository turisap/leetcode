package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestISPRecursiveSubTest
	1a)Test Iterative
	    go test -run TestISPIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestISPRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tt := testCase
			result := ispRecursive(tt.input)

			if result != tt.result {
				t.Errorf("ISP %v produces %t, got %t", tt.input, tt.result, result)
			}
		})
	}
}

func TestISPIterative(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tt := testCase
			result := ispIterative(tt.input)

			if result != tt.result {
				t.Errorf("ISP %v produces %t, got %t", tt.input, tt.result, result)
			}
		})
	}
}
