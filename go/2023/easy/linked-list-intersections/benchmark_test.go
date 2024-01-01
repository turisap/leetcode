package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=LLIHashTable  -run=XXX  -cpuprofile=cpu.pprof
      2) cpu profile interactive
   		go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=LLIHashTable -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
   		go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=LLITwoPointersImproved  -run=XXX  -cpuprofile=cpu.pprof
     2) cpu profile interactive
        go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=LLITwoPointersImproved -run=XXX -memprofile=mem.pprof
     4) mem profile interactive
        go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count
       go test -count=10  -bench=BenchmarkLLIHashTable -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkLLITwoPointersImproved -run=XXX -benchmem > stat2
     7) compare with
         benchstat stat1 stat2
*/

func BenchmarkLLIHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lliHashTable(testCases[0].head1, testCases[0].head2)
	}
}

func BenchmarkLLITwoPointersPlain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lliTwoPointersPlain(testCases[0].head1, testCases[0].head2)
	}
}

func BenchmarkLLITwoPointersImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lliTwoPointersImproved(testCases[0].head1, testCases[0].head2)
	}
}
