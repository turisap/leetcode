package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=CLSReg  -run=XXX  -cpuprofile=cpu.pprof
      2) cpu profile interactive
   		go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=CLSReg -run=XXX -memprofile=mem.pprof
      4) mem profile interactive
   		go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=CLSL  -run=XXX  -cpuprofile=cpu.pprof
     2) cpu profile interactive
        go tool pprof cpu.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=CLSL -run=XXX -memprofile=mem.pprof
     4) mem profile interactive
        go tool pprof mem.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count
       go test -count=10  -bench=BenchmarkCLSReg -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkCLSL -run=XXX -benchmem > stat2
     7) compare with
         benchstat stat1 stat2
*/

func BenchmarkSubCLSRegWith(b *testing.B) {
	for _, testCase := range testCases {
		testCase := testCase
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CLSReg(testCase.input)
			}
		})
	}
}

func BenchmarkSubCLSDynamic(b *testing.B) {
	for _, testCase := range testCases {
		testCase := testCase
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CLSDynamic(testCase.input)
			}
		})
	}
}

func BenchmarkCLSReg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CLSReg(25)
	}
}

func BenchmarkCLSL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CLSL(25)
	}
}

func BenchmarkCLSDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CLSDynamic(25)
	}
}
