package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	1) Test Recursive
	    go test -run TestAPDRecursiveSubTest
	1a)Test Iterative
	    go test -run TestAPDIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestAPDItem(t *testing.T) {
	for _, testCase := range testCases {
		tt := testCase

		t.Run(testCase.name, func(t *testing.T) {
			result := appendCustom(tt.s, tt.item)
			assert.ElementsMatch(t, result, tt.result)
		})
	}
}

func TestAPDArgs(t *testing.T) {
	for _, testCase := range testMulti {
		tt := testCase

		t.Run(testCase.name, func(t *testing.T) {
			result := appendCustom(tt.s, tt.args...)
			assert.ElementsMatch(t, result, tt.result)
		})
	}
}
