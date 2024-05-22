package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type TestCase struct {
	name   string
	input  *os.File
	result Result
}

/*
	1) Test Recursive
	    go test -run TestOPTRecursiveSubTest
	1a)Test Iterative
	    go test -run TestOPTIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestOPTRecursiveSubTest(t *testing.T) {
	testCase := TestCase{
		name:   "basic test",
		input:  nil,
		result: Result{},
	}
	t.Run(testCase.name, func(t *testing.T) {
		result := s1(testCase.input)
		assert.Equal(t, result, testCase.result)
	})
}

//func TestOPTIterativeSubTest(t *testing.T) {
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			result := optIterative()
//			if result != testCase.result {
//				t.Errorf("OPT %s produces %d, got %d", testCase.s, testCase.result, result)
//			}
//		})
//	}
//}
