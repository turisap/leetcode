package main

import "testing"

/*
	1) go test -run TestSIPSubTest || go test -run TestSIPLSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestSIPSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := SIPReg(testCase.nums, testCase.target)
			if result != testCase.result {
				t.Errorf("SIP %v and %d produces %d, got %d", testCase.nums, testCase.target, testCase.result, result)
			}
		})
	}
}

func TestSIPLSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := SIPL(testCase.nums, testCase.target)
			if result != testCase.result {
				t.Errorf("SIP %v and %d produces %d, got %d", testCase.nums, testCase.target, testCase.result, result)
			}
		})
	}
}
