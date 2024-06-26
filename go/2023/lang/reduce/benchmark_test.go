package main

import (
	"github.com/thoas/go-funk"
	"testing"
)

/*
      == 1st implementation ==
      1) go test -v -bench=RDCRecursive  -run=XXX  -cpuprofile=cpu1.pprof
      2) cpu profile interactive
   		go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=RDCRecursive -run=XXX -memprofile=mem1.pprof
      4) mem profile interactive
   		go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=RDCIterative  -run=XXX  -cpuprofile=cpu2.pprof
     2) cpu profile interactive
        go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=RDCIterative -run=XXX -memprofile=mem2.pprof
     4) mem profile interactive
        go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count and compare
       go test -count=10  -bench=BenchmarkRDCRecursive -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkRDCIterative -run=XXX -benchmem > stat2 &&
       benchstat stat1 stat2
*/

//  func BenchmarkSubRDCRegWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				rdcRecursive(testCase.s)
//  			}
//  		})
//  	}
//  }

var testCase = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func reducer(acc, val int) int {
	return acc + val
}

func BenchmarkRDCLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funk.Reduce(testCase, reducer, 0)
	}
}

func BenchmarkRDC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduceCustom[int, []int](testCase, reducer)
	}
}
