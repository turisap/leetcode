package main

/*
   == 1st implementation ==
    1) go test -v -bench=OPTRecursive  -run=XXX  -cpuprofile=cpu1.pprof
    2) cpu profile interactive
     go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
    3) go test -v -bench=OPTRecursive -run=XXX -memprofile=mem1.pprof
    4) mem profile interactive
     go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   == 2nd implementation ==
   1) go test -v -bench=OPTIterative  -run=XXX  -cpuprofile=cpu2.pprof
   2) cpu profile interactive
      go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
   3) go test -v -bench=OPTIterative -run=XXX -memprofile=mem2.pprof
   4) mem profile interactive
      go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   == comparison ==
   5) compare results of two implementations in console
      go test -v -bench . -benchmem -run=XXX
   6) create cpu profile  with count
     go test -count=10  -bench=BenchmarkOPTRecursive -run=XXX -benchmem  > stat1 &&
     go test -count=10  -bench=BenchmarkOPTIterative -run=XXX -benchmem > stat2
   7) compare with
       benchstat stat1 stat2
*/

//  func BenchmarkSubOPTRegWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				optRecursive(testCase.s)
//  			}
//  		})
//  	}
//  }

//  func BenchmarkSubOPTIterativeWith(b *testing.B) {
//  	for _, testCase := range testCases {
//  		testCase := testCase
//  		b.Run(testCase.name, func(b *testing.B) {
//  			for i := 0; i < b.N; i++ {
//  				optIterative(testCase.s)
//  			}
//  		})
//  	}
//  }

//func BenchmarkOPTRecursive(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		optRecursive()
//	}
//}
//
//func BenchmarkOPTIterative(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		optIterative()
//	}
//}
