package main

import "testing"

/*
      == 1st implementation ==
      1) go test -v -bench=STCStruct  -run=XXX  -cpuprofile=cpu1.pprof
      2) cpu profile interactive
   		go tool pprof cpu1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
      3) go test -v -bench=STCStruct -run=XXX -memprofile=mem1.pprof
      4) mem profile interactive
   		go tool pprof mem1.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == 2nd implementation ==
     1) go test -v -bench=STCSlice  -run=XXX  -cpuprofile=cpu2.pprof
     2) cpu profile interactive
        go tool pprof cpu2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"
     3) go test -v -bench=STCSlice -run=XXX -memprofile=mem2.pprof
     4) mem profile interactive
        go tool pprof mem2.pprof | Interactive command "top", "web", "list <<FUNC NAME>>"

     == comparison ==
     5) compare results of two implementations in console
        go test -v -bench . -benchmem -run=XXX
     6) create cpu profile  with count and compare
       go test -count=10  -bench=BenchmarkSTCStruct -run=XXX -benchmem  > stat1 &&
       go test -count=10  -bench=BenchmarkSTCSlice -run=XXX -benchmem > stat2 &&
       benchstat stat1 stat2
*/

const n = 1000000

func BenchmarkSTCStruct(b *testing.B) {
	input := make([]Foo, n)

	for i := 0; i < n; i++ {
		c := int64(i)
		input[i] = Foo{a: c, b: c}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumFoo(input)
	}
}

func BenchmarkSTCSlice(b *testing.B) {
	input := Bar{
		a: make([]int64, n),
		b: make([]int64, n),
	}

	for i := 0; i < n; i++ {
		c := int64(i)
		input.a[i] = c
		input.b[i] = c
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumBar(input)
	}
}
