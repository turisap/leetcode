package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=ISPRecursive  -run=XXX  -cpuprofile=cpu1.pprof
      2) cpu profile interactive
   		go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=ISPRecursive -run=XXX -memprofile=mem1.pprof
      4) mem profile interactive
   		go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=ISPIterative  -run=XXX  -cpuprofile=cpu2.pprof
     2) cpu profile interactive
        go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=ISPIterative -run=XXX -memprofile=mem2.pprof
     4) mem profile interactive
        go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count and compare
       go test -count=10  -bench=BenchmarkISPRecursive -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkISPIterative -run=XXX -benchmem > stat2 &&
       benchstat stat1 stat2
*/

func BenchmarkISPRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ispRecursive("akaabakolokabaaka")
	}
}

func BenchmarkISPIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ispIterative("akaabakolokabaaka")
	}
}
