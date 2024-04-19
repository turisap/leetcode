package main

import (
	"testing"
)

var testCases = []struct {
	name   string
	a      []int
	result []int
}{
	{"Simple array", []int{1, 2, 3}, []int{3, 2, 1}},
	{"Mirror array", []int{1, 2, 2, 1}, []int{1, 2, 2, 1}},
}

func TestAITRecursiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.a)

			for i := range result {
				if result[i] != tt.result[i] {
					t.Errorf("%s returns %v, expected %v\n", tt.name, result, tt.result)
				}
			}
		})
	}
}
