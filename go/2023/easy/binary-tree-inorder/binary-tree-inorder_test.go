package main

import (
	"fmt"
	"testing"
)

/*
	1) Test naive
	    go test -run TestBTIRecursiveSubTest
	1a)Test improved
	    go test -run BTIIterative
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func TestBTIRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := fmt.Sprintf("%v", btiRecursive(testCase.tree))

			fmt.Println(result, testCase.result)
			if result != testCase.result {
				t.Errorf("BTI %v produces %s, got %s", testCase.tree, testCase.result, result)
			}
		})
	}
}

func TestBTIIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := fmt.Sprintf("%v", btiIterative(testCase.tree))

			fmt.Println(result, testCase.result)
			if result != testCase.result {
				t.Errorf("BTI %v produces %s, got %s", testCase.tree, testCase.result, result)
			}
		})
	}
}
