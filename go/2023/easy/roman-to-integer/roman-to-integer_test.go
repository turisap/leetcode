package main

import "testing"

/*
	1) go test -run TestRTISubTest
	2) check coverage "go test -coverprofile=coverage.out"
	3) check coverage by function: "go tool cover -func=coverage.out"
	4) check coverage: "go tool cover -html=coverage.out"
*/

func TestRTISubTest(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		result int
	}{
		{"II", "II", 2},
		{"IV", "IV", 4},
		{"IX", "IX", 9},
		{"L", "L", 50},
		{"C", "C", 100},
		{"CM", "CM", 900},
		{"CD", "CD", 400},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := RTIReg(testCase.s)
			if result != testCase.result {
				t.Errorf("Roman %s produces %d, got %d", testCase.s, testCase.result, result)
			}
		})
	}
}

func TestRTILSubTest(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		result int
	}{
		{"II", "II", 2},
		{"IV", "IV", 4},
		{"IX", "IX", 9},
		{"L", "L", 50},
		{"C", "C", 100},
		{"CM", "CM", 900},
		{"CD", "CD", 400},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := RTIL(testCase.s)
			if result != testCase.result {
				t.Errorf("Roman %s produces %d, got %d", testCase.s, testCase.result, result)
			}
		})
	}
}
