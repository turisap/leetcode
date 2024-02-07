package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	1) Test Recursive
	    go test -run TestAITRecursiveSubTest
	1a)Test Iterative
	    go test -run TestAITIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestAITRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			result := aitImproved(tt.a, tt.b)
			if !assert.ElementsMatch(t, result, tt.result) {
				t.Errorf("AIT %v and %v should produce %v, got %v", tt.a, tt.b, tt.result, result)
			}
		})
	}
}

func TestAITIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			result := aitIterative(tt.a, tt.b)
			if !assert.ElementsMatch(t, result, tt.result) {
				t.Errorf("AIT %v and %v should produce %v, got %v", tt.a, tt.b, tt.result, result)
			}
		})
	}
}
