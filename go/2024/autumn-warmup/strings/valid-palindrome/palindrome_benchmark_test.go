package main

import "testing"

var res bool
var inputShort = "`l;`` 1o1 ??;l`"
var inputLong = "X;;];;&;2;H;;];J;;#;[;;6;Z;<;N;_;>; > _ N < Z 6  [ #  J ]  H 2 &  ]  X"

/*
BenchmarkPalindromeReverseShort-8     	 1037161	      1126 ns/op	    1047 B/op	      17 allocs/op
BenchmarkPalindromeTwoPointsShort-8   	10193415	       117.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkPalindromeReverseLong-8      	  468519	      2556 ns/op	    1305 B/op	      19 allocs/op
BenchmarkPalindromeTwoPointsLong-8    	 1814292	       659.3 ns/op	      10 B/op	      10 allocs/op
*/
func BenchmarkPalindromeReverseShort(b *testing.B) {
	for _ = range b.N {
		res = isPalindromeReverse(inputShort)
	}
}

func BenchmarkPalindromeTwoPointsShort(b *testing.B) {
	for _ = range b.N {
		res = isPalindromeTwoPointer(inputShort)
	}
}
func BenchmarkPalindromeReverseLong(b *testing.B) {
	for _ = range b.N {
		res = isPalindromeReverse(inputLong)
	}
}

func BenchmarkPalindromeTwoPointsLong(b *testing.B) {
	for _ = range b.N {
		res = isPalindromeTwoPointer(inputLong)
	}
}
