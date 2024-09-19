package main

import "testing"

var res []int

var (
	a   = []int{1, 2, 3, 0, 0, 0}
	x   = []int{2, 5, 6}
	m   = 3
	n   = 3
	out = []int{1, 2, 2, 3, 5, 6}
)

func BenchmarkNaive(b *testing.B) {
	for _ = range b.N {
		// go:noinline
		MergeNaive(a, m, x, n)
	}
}
