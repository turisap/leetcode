package main

import "testing"

/*
	1) go test -run TestCLSSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestCLSSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := CLSReg(testCase.input)
			if result != testCase.result {
				t.Errorf("CLS input = %d produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}

func TestCLSLSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := CLSL(testCase.input)
			if result != testCase.result {
				t.Errorf("CLS input = %d produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}
func TestCLSDSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := CLSDynamic(testCase.input)
			if result != testCase.result {
				t.Errorf("CLS input = %d produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}
