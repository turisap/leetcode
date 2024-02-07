package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	1) Test All
	    go test -run TestAPD
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

		t.Run(tt.name, func(t *testing.T) {
			result := appendCustom(tt.s, tt.item)
			assert.ElementsMatch(t, result, tt.result)
		})
	}
}

func TestAPDArgs(t *testing.T) {
	for _, testCase := range testMulti {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			result := appendCustom(tt.s, tt.args...)
			assert.ElementsMatch(t, result, tt.result)
		})
	}
}

func TestAPDEnoughCap(t *testing.T) {
	s := make([]int, 2, 5)

	t.Run("Add a set of integers to an array when cap is enough", func(t *testing.T) {
		result := appendCustom(s, []int{2, 3, 4}...)
		fmt.Println(result)
		assert.ElementsMatch(t, result, []int{0, 0, 2, 3, 4})
	})
}
