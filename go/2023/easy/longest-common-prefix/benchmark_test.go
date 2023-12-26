package main

import "testing"

/*
   	1) run benchmarks without unit tests "go test -v -bench=LCP -run=XXX -benchmem"
   	2) run benchMarks for specific implementation "go test -bench=BenchmarkLCPReg -run=XXX -count=10 -benchmem > lcpreg.txt"
   	3) run benchMarks for another implementation "go test -bench=BenchmarkLCPit -run=XXX -count=10 -benchmem > lcprit.txt"
   	4) compare results "benchstat lcpreg.txt lcpit.txt" // CAN BE DONE WITH CPU AS WELL
   	5) cpu profile "go test -cpuprofile cpu.prof -bench=LCP -run=XXX"
    6) cpu profile interactive "go tool pprof cpu.prof". Interactive command "top", "web", "list <<FUNC NAME>>"
    7) memory profile "go test -memprofile mem.prof -bench=LCP -run=XXX".
    8) memory profile "go tool pprof mem.prof". Interactive command "top", "web", "list <<FUNC NAME>>"

	=========================
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

//func BenchmarkSubLCPRegWith(b *testing.B) {
//
//	for _, testCase := range testCases {
//		testCase := testCase
//		b.Run(testCase.name, func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				LCPReg(testCase.strs)
//			}
//		})
//	}
//}

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
