package main

import "testing"

/*
	1) Test Recursive
	    go test -run TestRLLRecursiveSubTest
	1a)Test Iterative
	    go test -run TestRLLIterativeSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil || l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func TestRLLRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := rllRecursive(testCase.head)
			if !isEqual(result, testCase.result) {
				t.Errorf("RLL %v produces %v, got %v", testCase.head, testCase.result, result)
			}
		})
	}
}

func TestRLLIterativeSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := rllIterative(testCase.head)
			if !isEqual(result, testCase.result) {
				t.Errorf("RLL %v produces %v, got %v", testCase.head, testCase.result, result)
			}
		})
	}
}
