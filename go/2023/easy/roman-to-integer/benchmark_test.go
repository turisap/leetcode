package main

import "testing"

/*
	1) run benchmarks without unit tests "go test -v -bench=RTI -run=XXX -benchmem"
	2) run benchMarks for specific implementation "go test -bench=BenchmarkRTIReg -run=XXX -count=10 -benchmem > rtireg.txt"
	3) run benchMarks for another implementation "go test -bench=BenchmarkRTIIt -run=XXX -count=10 -benchmem > rtirit.txt"
	3) compare results "benchstat rtireg.txt rtiit.txt"
*/

func BenchmarkSubRTIRegWith(b *testing.B) {
	testCases := []struct {
		name   string
		s      string
		result int
	}{
		{"II", "II", 2},
		{"CM", "CM", 900},
		{"MMMCMXCIX", "MMMCMXCIX", 3999},
	}

	for _, testCase := range testCases {
		testCase := testCase
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RTIReg(testCase.s)
			}
		})
	}
}

func BenchmarkRTIReg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RTIReg("MMMCMXCIX")
	}
}

func BenchmarkRTIL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RTIL("MMMCMXCIX")
	}
}
