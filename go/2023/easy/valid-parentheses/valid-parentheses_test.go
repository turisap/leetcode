package main

import "testing"

/*
	1) go test -run TestVPSSubTest
	2) check coverage
		go test -coverprofile=coverage.out
	3) check coverage by function:
		go tool cover -func=coverage.out
	4) check coverage:
		go tool cover -html=coverage.out
*/

func TestVPSSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := VPSReg(testCase.name)
			if result != testCase.result {
				t.Errorf("VPS %s produces %t, got %t", testCase.name, testCase.result, result)
			}
		})
	}
}

func TestVPSLSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := VPSL(testCase.name)
			if result != testCase.result {
				t.Errorf("VPS %s produces %t, got %t", testCase.name, testCase.result, result)
			}
		})
	}
}
