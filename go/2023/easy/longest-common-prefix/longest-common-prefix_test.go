package main

import "testing"

/*
	1) go test -run TestLCPSubTest
	2) check coverage "go test -coverprofile=coverage.out"
	3) check coverage by function: "go tool cover -func=coverage.out"
	4) check coverage: "go tool cover -html=coverage.out"
*/

func TestLCPSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := LCPReg(testCase.strs)
			if result != testCase.result {
				t.Errorf("LCP %v produces %s, got %s", testCase.strs, testCase.result, result)
			}
		})
	}
}

func TestLCPLSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := LCPL(testCase.strs)
			if result != testCase.result {
				t.Errorf("LCP %v produces %s, got %s", testCase.strs, testCase.result, result)
			}
		})
	}
}
