package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	1) Test Naive
	    go test -run TestPTRNaiveSubTest
	1a)Test Iterative
	    go test -run TestPTRImprovedSubTest
	2) check coverage
	    go test -coverprofile=coverage.out
	3) check coverage by function:
	    go tool cover -func=coverage.out
	4) check coverage:
	    go tool cover -html=coverage.out
*/

func sliceToString(slice interface{}) string {
	value := reflect.ValueOf(slice)

	str := "["
	for i := 0; i < value.Len(); i++ {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", value.Index(i))
	}
	str += "]"

	return str
}

func TestPTRNaiveSubTest(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := ptrNaive(testCase.n)
			if sliceToString(result) != sliceToString(testCase.result) {
				t.Errorf("PTR %d produces %s, got %s", testCase.n, sliceToString(testCase.result), sliceToString(result))
			}
		})
	}
}
