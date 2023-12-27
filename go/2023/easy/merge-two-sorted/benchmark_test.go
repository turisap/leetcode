package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=MTSReg  -run=XXX  -cpuprofile=cpu.pprof
      2) cpu profile interactive
   		go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=MTSReg -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
   		go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   	  == 2nd implementation ==
       1) go test -v -bench=MTSReg  -run=XXX  -cpuprofile=cpu.pprof
       2) cpu profile interactive
      	go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=MTSReg -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
      	go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      5) compare results of two implementations in console
   		go test -v -bench . -benchmem -run=XXX
      6) create cpu profile  with count
        go test -count=10  -bench=BenchmarkMTSReg -run=XXX -benchmem  > stat3 &&
        go test -count=10  -bench=BenchmarkMTSL -run=XXX -benchmem > stat4 &&
		benchstat stat1 stat2

*/

//  func BenchmarkSubMTSRegWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				MTSReg(testCase.s)
//  			}
//  		})
//  	}
//  }

func BenchmarkMTSReg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}

		list2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 10}}}}}

		MTSReg(list1, list2)
	}
}

func BenchmarkMTSL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}

		list2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 10}}}}}

		MTSL(list1, list2)
	}
}
