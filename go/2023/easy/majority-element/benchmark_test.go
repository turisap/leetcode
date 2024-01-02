package main

import "testing"

/*
	   == 1st implementation ==
	   1) go test -v -bench=MJENaive  -run=XXX  -cpuprofile=cpu1.pprof
	   2) cpu profile interactive
			go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
	   3) go test -v -bench=MJENaive -run=XXX -memprofile=mem1.pprof
	   4) mem profile interactive
			go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

	  == 2nd implementation ==
	  1) go test -v -bench=MJEImproved  -run=XXX  -cpuprofile=cpu2.pprof
	  2) cpu profile interactive
	     go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
	  3) go test -v -bench=MJEImproved -run=XXX -memprofile=mem2.pprof
	  4) mem profile interactive
	     go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

	  == comparison ==
	  5) compare results of two implementations in console
	     go test -v -bench . -benchmem -run=XXX
	  6) create cpu profile  with count and compare
	    go test -count=10  -bench=BenchmarkMJENaive -run=XXX -benchmem  > stat1 &&
	    go test -count=10  -bench=BenchmarkMJEImproved -run=XXX -benchmem > stat2 &&
	    benchstat stat1 stat2
*/
func BenchmarkMJENaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mjeNaive(testCases[0].input)
	}
}

func BenchmarkMJEImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mjeImproved(testCases[0].input)
	}
}

func BenchmarkMJEVoting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mjeVoting(testCases[0].input)
	}
}

func BenchmarkMJEEditorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mjeEDITORIAL(testCases[0].input)
	}
}
