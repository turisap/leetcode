package main

import "testing"

/*
   == 1st implementation ==
    1) go test -v -bench=DARRecursive  -run=XXX  -cpuprofile=cpu1.pprof
    2) cpu profile interactive
     go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
    3) go test -v -bench=DARRecursive -run=XXX -memprofile=mem1.pprof
    4) mem profile interactive
     go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   == 2nd implementation ==
   1) go test -v -bench=DARIterative  -run=XXX  -cpuprofile=cpu2.pprof
   2) cpu profile interactive
      go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
   3) go test -v -bench=DARIterative -run=XXX -memprofile=mem2.pprof
   4) mem profile interactive
      go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

   == comparison ==
   5) compare results of two implementations in console
      go test -v -bench . -benchmem -run=XXX
   6) create cpu profile  with count
     go test -count=10  -bench=BenchmarkDARRecursive -run=XXX -benchmem  > stat1 &&
     go test -count=10  -bench=BenchmarkDARIterative -run=XXX -benchmem > stat2
   7) compare with
       benchstat stat1 stat2
*/

var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1718, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var testK = 2

func BenchmarkDARSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = darSimple(testSlice, testK)
	}
}

func BenchmarkDARMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = darMapPointers(testSlice, testK)
	}
}
