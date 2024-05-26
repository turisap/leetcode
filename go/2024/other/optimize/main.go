package main

import "testing"

func main() {
	//f, err := getMeasurementsFile()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer f.Close()
	//
	//_ = s7(f)

	stat := func(f func() byte) int {
		allocs := testing.AllocsPerRun(1, func() {
			f()
		})
		return int(allocs)
	}
	println(stat(f)) // 0
}
