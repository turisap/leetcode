package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=SIPReg  -run=XXX  -cpuprofile=cpu.pprof
      2) cpu profile interactive
   		go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=SIPReg -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
   		go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   	  == 2nd implementation ==
       1) go test -v -bench=SIPL  -run=XXX  -cpuprofile=cpu.pprof
       2) cpu profile interactive
      	go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=SIPL -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
      	go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      5) compare results of two implementations in console
   		go test -v -bench . -benchmem -run=XXX
      6) create cpu profile  with count
        go test -count=10  -bench=BenchmarkSIPReg -run=XXX -benchmem  > stat1 &&
        go test -count=10  -bench=BenchmarkSIPL -run=XXX -benchmem > stat2
      7) compare with
          benchstat stat1 stat2
*/

//  func BenchmarkSubSIPRegWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				SIPReg(testCase.s)
//  			}
//  		})
//  	}
//  }

func BenchmarkSIPReg(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		SIPReg(input, 18)
	}
}

func BenchmarkSIPL(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		SIPL(input, 18)
	}
}
