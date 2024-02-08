package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestRDCRecursiveSubTest

	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestRDCRecursiveSubTestInt(t *testing.T) {
	for _, testCase := range intCases {
		tt := testCase
		t.Run(testCase.name, func(t *testing.T) {
			result := reduceCustom[int, []int](tt.input, func(acc, val int) int {
				return acc + val
			})
			if result != tt.result {
				t.Errorf("RDC %v produces %d, got %d", tt.input, tt.result, result)
			}
		})
	}
}

func TestRDCRecursiveSubTestString(t *testing.T) {
	for _, testCase := range stringCases {
		tt := testCase
		t.Run(testCase.name, func(t *testing.T) {
			result := reduceCustom[string, []string](tt.input, func(acc, val string) string {
				return acc + val
			})
			if result != tt.result {
				t.Errorf("RDC %v produces %s, got %s", tt.input, tt.result, result)
			}
		})
	}
}
