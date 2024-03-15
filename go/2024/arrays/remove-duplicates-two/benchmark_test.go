package main

import "testing"

/*
   == 1st implementation ==
    1) go test -v -bench=rdwSwap  -run=XXX  -cpuprofile=cpu1.pprof
    2) cpu profile interactive
     go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
    3) go test -v -bench=rdwSwap -run=XXX -memprofile=mem1.pprof
    4) mem profile interactive
     go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   == 2nd implementation ==
   1) go test -v -bench=RDWIterative  -run=XXX  -cpuprofile=cpu2.pprof
   2) cpu profile interactive
      go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
   3) go test -v -bench=RDWIterative -run=XXX -memprofile=mem2.pprof
   4) mem profile interactive
      go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   == comparison ==
   5) compare results of two implementations in console
      go test -v -bench . -benchmem -run=XXX
   6) create cpu profile  with count
     go test -count=10  -bench=BenchmarkrdwSwap -run=XXX -benchmem  > stat1 &&
     go test -count=10  -bench=BenchmarkRDWIterative -run=XXX -benchmem > stat2
   7) compare with
       benchstat stat1 stat2
*/

//  func BenchmarkSubRDWRegWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				rdwSwap(testCase.s)
//  			}
//  		})
//  	}
//  }

//  func BenchmarkSubRDWIterativeWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				rdwIterative(testCase.s)
//  			}
//  		})
//  	}
//  }

func BenchmarkRDWSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rdwSwap(testCases[3].s)
	}
}

func BenchmarkRDWIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rdwTwo(testCases[3].s)
	}
}
