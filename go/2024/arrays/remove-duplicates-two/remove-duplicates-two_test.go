package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestRDWSwapSubTest
	1a)Test Iterative
	    go test -run TestRDWIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestRDWSwapSubTest(t *testing.T) {
	for _, testCase := range testCases {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			result := rdwSwap(tt.s)
			if result != tt.result {
				t.Errorf("RDW %v produces %d, got %d", tt.s, tt.result, result)
			}
		})
	}
}

//func TestRDWIterativeSubTest(t *testing.T) {
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			result := rdwIterative()
//			if result != testCase.result {
//				t.Errorf("RDW %s produces %d, got %d", testCase.s, testCase.result, result)
//			}
//		})
//	}
//}
