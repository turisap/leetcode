package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=btiRecursive  -run=XXX  -cpuprofile=cpu.pprof
      2) cpu profile interactive
   		go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=btiRecursive -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
   		go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=BTIIterative  -run=XXX  -cpuprofile=cpu.pprof
     2) cpu profile interactive
        go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=BTIIterative -run=XXX -memprofile=mem.pprof
     4) mem profile interactive
        go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count
       go test -count=10  -bench=BenchmarkBTIRecursive -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkBTIIterative -run=XXX -benchmem > stat2
     7) compare with
         benchstat stat1 stat2
*/

func BenchmarkBTIRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btiRecursive(testCases[0].tree)
	}
}

func BenchmarkBTIIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btiIterative(testCases[0].tree)
	}
}
