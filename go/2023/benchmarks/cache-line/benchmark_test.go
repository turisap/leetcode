package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=CHL2  -run=XXX  -cpuprofile=cpu1.pprof
      2) cpu profile interactive
   		go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=CHL2 -run=XXX -memprofile=mem1.pprof
      4) mem profile interactive
   		go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=CHL8  -run=XXX  -cpuprofile=cpu2.pprof
     2) cpu profile interactive
        go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=CHL8 -run=XXX -memprofile=mem2.pprof
     4) mem profile interactive
        go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count and compare
       go test -count=10  -bench=BenchmarkCHL2 -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkCHL8 -run=XXX -benchmem > stat2 &&
       benchstat stat1 stat2
*/

func BenchmarkCHL2(b *testing.B) {
	l := 10000
	s := make([]int64, l)

	for i := 0; i < l; i++ {
		s[i] = int64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum2(s)
	}
}

func BenchmarkCHL8(b *testing.B) {
	l := 10000
	s := make([]int64, l)

	for i := 0; i < l; i++ {
		s[i] = int64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum8(s)
	}
}
