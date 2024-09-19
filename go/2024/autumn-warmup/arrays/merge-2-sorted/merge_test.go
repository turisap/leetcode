package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	a, b []int
	m, n int
	out  []int
}

var tcc = []TestCase{
	{
		a:   []int{1, 2, 3, 0, 0, 0},
		b:   []int{2, 5, 6},
		m:   3,
		n:   3,
		out: []int{1, 2, 2, 3, 5, 6},
	},
	{
		a:   []int{10, 11, 15, 0, 0},
		b:   []int{1, 2},
		m:   3,
		n:   2,
		out: []int{1, 2, 10, 11, 15},
	},
}

func TestSortNaive(t *testing.T) {
	for _, tc := range tcc {
		MergeNaive(tc.a, tc.m, tc.b, tc.n)
		assert.Equal(t, tc.out, tc.a)
	}
}
