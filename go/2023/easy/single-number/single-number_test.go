package main

import "testing"

/*
	1) Test XOR
	    go test -run TestSNRXORSubTest
	1a)Test HashTablejj
	    go test -run TestSNRHashTableSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestSNRXORSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := snrXOR(testCase.input)
			if result != testCase.result {
				t.Errorf("SNR %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}

func TestSNRHashTableSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := snrHashTable(testCase.input)
			if result != testCase.result {
				t.Errorf("SNR %v produces %d, got %d", testCase.input, testCase.result, result)
			}
		})
	}
}
