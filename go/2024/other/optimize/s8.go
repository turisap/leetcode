package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
Byte slice map key to my assignment
Stack allocation,
memory ballast
Manual inline
For range to for
Writing global writing local
method on an interface and on a concrete struct
Map  pointers instead of values
struct sorting fields
preallocate map (or GC ballast)
@TODO use ints instead of floats
@TODO Boundary checks
@TODO Allocate the backing array of a slice on stack even if its size is larger than or equal to 64K (but not larger than 10M)
@TODO Avoid unnecessary pointer dereferences in a loop
@TODO Avoid accessing fields of a struct in a loop though pointers to the struct
@TODO Use index tables instead of maps which key types have only a small set of possible
*/

/*
outside of this
Stack allocation for recursive functions
defer (separately)
map prealloc
slice prealloc
*/

const LENGTH = 50000

func s8(f *os.File) Result {
	scanner := bufio.NewScanner(f)
	fMap := map[[32]byte]*FahrenheitReorder{}
	cMap := map[[32]byte]*CelsiusReorder{}
	//targetSources := [][32]byte{}
	targetSources := make([][32]byte, 0, LENGTH)
	c := 0

	const separatorIndex = 32
	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		var randomStringBytes [separatorIndex]byte
		copy(randomStringBytes[:], lineBytes[:separatorIndex])
		randomFloatPart := lineBytes[separatorIndex+2:]

		temperature, err := strconv.ParseFloat(string(randomFloatPart), 32)

		if err != nil {
			continue
		}

		if c%2 == 0 {
			targetSources = append(targetSources, randomStringBytes)
		}

		fMap[randomStringBytes] = &FahrenheitReorder{
			temp:   temperature,
			time:   time.Now(),
			active: false,
			source: string(randomStringBytes[:]),
		}
		c++
	}

	// map to CelsiusReorder
	for k, v := range fMap {
		e := v.temp * 2.0
		l := v.temp / 3.87
		a := 7*e + 3*l

		cMap[k] = &CelsiusReorder{
			temp:      (v.temp - 32) * 0.55,
			source:    v.source,
			footprint: a,
		}
	}

	var (
		minTempLocal = 0.0
		maxTempLocal = 0.0
		fKoeffLocal  = 0.0
		sumLocal     = 0.0
	)

	for i := 0; i < len(targetSources); i++ {
		s := targetSources[i]
		m := cMap[s]

		if m == nil {
			break
		}
		curr := m.temp

		if curr < minTempLocal {
			minTempLocal = curr
		}

		if curr > maxTempLocal {
			maxTempLocal = curr
		}

		if i%TARGET_KOEF_IDX == 0 {
			a := curr + 3.4
			b := curr - 20.32
			k := curr * 3.1 / 4.1
			fKoeffLocal = 3.0*a + 5.2*b - k*2.2 // complex float manipulations
		}

		sumLocal += curr
	}

	avgStr := fmt.Sprintf("%.2f", sumLocal/float64((len(targetSources))))
	minTStr := fmt.Sprintf("%.2f", minTempLocal)
	maxTStr := fmt.Sprintf("%.2f", maxTempLocal)
	fKoefStr := fmt.Sprintf("%.2f", fKoeffLocal)

	return Result{
		avg:     avgStr,
		minTemp: minTStr,
		maxTemp: maxTStr,
		fK:      fKoefStr,
	}
}
