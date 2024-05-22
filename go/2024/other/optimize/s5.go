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
@TODO Stack allocation,
@TODO memory ballast
@TODO Manual inline
For range to for
@TODO Boundary checks
Writing global writing local
@TODO No defer
@TODO use ints instead of floats
@TODO Byte slice map key to map assignment
Map  pointers instead of values
@TODO Custom hash map
@TODO struct sorting fields
@TODO Allocate the backing array of a slice on stack even if its size is larger than
@TODO or equal to 64K (but not larger than 10M)
@TODO preallocate map (or GC ballast)
@TODO Avoid unnecessary pointer dereferences in a loop
@TODO Avoid accessing fields of a struct in a loop though pointers to the struct
@TODO Use index tables instead of maps which key types have only a small set of possible
@TODO values
@TODO method on an interface and on a concrete struct
@TODO defer (separately)
*/

// @TODO all optimizations for utils.go file

func s5(f *os.File) Result {
	scanner := bufio.NewScanner(f)
	fMap := map[[32]byte]*Fahrenheit{}
	cMap := map[[32]byte]*Celsius{}
	// @TODO pre-alloc
	targetSources := [][32]byte{}
	c := 0

	const separatorIndex = 32
	for scanner.Scan() {
		lineBytes := scanner.Bytes()

		// @TODO make const 32
		var randomStringBytes [separatorIndex]byte
		copy(randomStringBytes[:], lineBytes[:32])
		randomFloatPart := lineBytes[separatorIndex+2:]

		temperature, err := strconv.ParseFloat(string(randomFloatPart), 32)

		if err != nil {
			continue
		}

		if isTargetSource(c) {
			targetSources = append(targetSources, randomStringBytes)
		}

		fMap[randomStringBytes] = &Fahrenheit{
			temp:   temperature,
			time:   time.Now(),
			active: false,
			source: string(randomStringBytes[:]),
		}
		c++
	}

	// map to celsius @TODO (pre-alloc map)
	for k, v := range fMap {
		m := fToCelsius(v)
		c, ok := m.(Celsius)

		if !ok {
			continue
		}

		cMap[k] = &c

	}

	var (
		minTempLocal = 0.0
		maxTempLocal = 0.0
		fKoeffLocal  = 0.0
		sumLocal     = 0.0
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

		if i%TARGET_KOEF_IDX == 0 {
			fKoeffLocal = calcFloatSth(curr) // complex float manipulations
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
