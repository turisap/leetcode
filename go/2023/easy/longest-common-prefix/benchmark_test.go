package main

import "testing"

/*
	1) go test -v -bench=  -run=XXX  -cpuprofile=cpu.pprof
	2) cpu profile interactive "go tool pprof cpu.pprof". Interactive command "top", "web", "list <<FUNC NAME>>"
	3) go test -v -bench= -run=XXX -memprofile=mem.pprof
	4) mem profile interactive "go tool pprof mem.pprof". Interactive command "top", "web", "list <<FUNC NAME>>"
	5) create cpu profile  with count
		go test -count=10 -v -bench=BenchmarkLCPReg -run=XXX -benchmem  >> cpu1 &&
		go test -count=10 -v -bench=BenchmarkLCPL -run=XXX -benchmem >> cpu2
	6) compare with
		benchstat cpu1 cpu2
*/

func BenchmarkLCPReg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCPReg(testCases[0].strs)
	}
}

func BenchmarkLCPL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCPL(testCases[0].strs)
	}
}

func BenchmarkLCPSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCPSort(testCases[0].strs)
	}
}
