package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=APDRecursive  -run=XXX  -cpuprofile=cpu1.pprof
      2) cpu profile interactive
   		go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=APDRecursive -run=XXX -memprofile=mem1.pprof
      4) mem profile interactive
   		go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=APDIterative  -run=XXX  -cpuprofile=cpu2.pprof
     2) cpu profile interactive
        go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=APDIterative -run=XXX -memprofile=mem2.pprof
     4) mem profile interactive
        go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count and compare
       go test -count=10  -bench=BenchmarkAPDRecursive -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkAPDIterative -run=XXX -benchmem > stat2 &&
       benchstat stat1 stat2
*/

//  func BenchmarkSubAPDRegWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				apdRecursive(testCase.s)
//  			}
//  		})
//  	}
//  }

//  func BenchmarkSubAPDIterativeWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				apdIterative(testCase.s)
//  			}
//  		})
//  	}
//  }

var initial = []int{1, 2, 3, 4, 5, 6, 7, 8}
var appending = []int{1, 2, 3, 4, 6}

var globalC []int

func BenchmarkAPDCustom(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = appendCustom(initial, appending...)
	}
	globalC = s
}

var globalS []int

func BenchmarkAPDStandard(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(initial, appending...)
	}

	globalS = s
}
