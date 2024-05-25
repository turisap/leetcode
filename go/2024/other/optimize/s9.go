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
method on an interface and on a concrete struct
Stack allocation for recursive functions
byte arrays map keys
map pointer to values
map iteration deref
defer (separately)
map prealloc
slice prealloc
*/

type CelsiusInt struct {
	time      time.Time
	source    string
	temp      int
	footprint int
	active    bool
}

type FahrenheitInt struct {
	time      time.Time
	source    string
	temp      int
	footprint int
	active    bool
}

const MULT_FACTOR = 1000

func s9(f *os.File) Result {
	scanner := bufio.NewScanner(f)
	fMap := map[[32]byte]*FahrenheitInt{}
	cMap := map[[32]byte]*CelsiusInt{}
	targetSources := [][32]byte{}
	c := 0

	const separatorIndex = 32
	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		var randomStringBytes [separatorIndex]byte
		copy(randomStringBytes[:], lineBytes[:32])
		randomFloatPart := lineBytes[separatorIndex+2:]

		temperature, err := strconv.ParseFloat(string(randomFloatPart), 32)
		tempInt := int(temperature * MULT_FACTOR)

		if err != nil {
			continue
		}

		if c%2 == 0 {
			targetSources = append(targetSources, randomStringBytes)
		}

		fMap[randomStringBytes] = &FahrenheitInt{
			temp:   tempInt,
			time:   time.Now(),
			active: false,
			source: string(randomStringBytes[:]),
		}
		c++
	}

	k1 := int(2.0 * MULT_FACTOR)
	k2 := int(3.87 * MULT_FACTOR)
	k3 := int(0.55 * MULT_FACTOR)
	for k, v := range fMap {
		e := v.temp * k1
		l := v.temp / k2
		a := 7*e + 3*l

		cMap[k] = &CelsiusInt{
			temp:      (v.temp - 32) * k3,
			source:    v.source,
			footprint: a,
		}
	}

	var (
		minTempLocal = 0
		maxTempLocal = 0
		fKoeffLocal  = 0
		sumLocal     = 0
	)

	for i := 0; i < len(targetSources); i++ {
		s := targetSources[i]
		curr := cMap[s].temp

		if curr < minTempLocal {
			minTempLocal = curr
		}

		if curr > maxTempLocal {
			maxTempLocal = curr
		}

		k1 := int(3.4 * MULT_FACTOR)
		k2 := int(20.32 * MULT_FACTOR)
		k3 := int(3.1 * MULT_FACTOR)
		k4 := int(4.1 * MULT_FACTOR)
		k5 := int(5.2 * MULT_FACTOR)
		k6 := int(2.2 * MULT_FACTOR)
		if i%TARGET_KOEF_IDX == 0 {
			a := curr + k1
			b := curr - k2
			k := curr * k3 / k4
			fKoeffLocal = 3*a + k5*b - k*k6 // complex float manipulations
		}

		sumLocal += curr
	}

	avgStr := fmt.Sprintf("%.2f", (float64(sumLocal)/float64(len(targetSources)))/MULT_FACTOR)
	minTStr := fmt.Sprintf("%.2f", float64(minTempLocal/MULT_FACTOR))
	maxTStr := fmt.Sprintf("%.2f", float64(maxTempLocal/MULT_FACTOR))
	fKoefStr := fmt.Sprintf("%.2f", float64(fKoeffLocal/MULT_FACTOR))

	return Result{
		avg:     avgStr,
		minTemp: minTStr,
		maxTemp: maxTStr,
		fK:      fKoefStr,
	}
}
