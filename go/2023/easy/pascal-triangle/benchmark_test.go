package main

import "testing"

/*
	   == 1st implementation ==
	   1) go test -v -bench=PTRNaive  -run=XXX  -cpuprofile=cpu.pprof
	   2) cpu profile interactive
			go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
	   3) go test -v -bench=PTRNaive -run=XXX -memprofile=mem.pprof
	   4) mem profile interactive
			go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"


	  == comparison ==
	  5) compare results of two implementations in console
	     go test -v -bench . -benchmem -run=XXX
	  6) create cpu profile  with count
	    go test -count=10  -bench=BenchmarkPTRNaive -run=XXX -benchmem  > stat1 &&
*/

func BenchmarkPTRNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptrNaive(50)
	}
}
