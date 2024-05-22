package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
@TODO Byte slice map key to my assignment
@TODO Stack allocation, memory ballast
@TODO Manual inline
@TODO For range to for
@TODO Boundary checks
Writing global writing local
@TODO No defer
@TODO use ints instead of floats
@TODO Byte slice map key to map assignment
@TODO Map  pointers instead of values
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

func s2(f *os.File) Result {
	scanner := bufio.NewScanner(f)
	fMap := map[string]Fahrenheit{}
	cMap := map[string]Celsius{}
	// @TODO pre-alloc
	targetSources := []string{}

	c := 0
	for scanner.Scan() {
		ln := scanner.Text()
		parts := strings.Split(ln, ": ")

		if len(parts) != 2 {
			continue
		}

		source := parts[0]
		temperature, err := strconv.ParseFloat(parts[1], 32)

		if err != nil {
			continue
		}

		if isTargetSource(c) {
			targetSources = append(targetSources, source)
		}

		fMap[source] = Fahrenheit{
			temp:   temperature,
			time:   time.Now(),
			active: false,
			source: source,
		}
		c++
	}

	// map to celsius @TODO (pre-alloc map)
	for _, v := range fMap {
		m := fToCelsius(v)
		c, ok := m.(Celsius)

		if !ok {
			continue
		}

		cMap[c.source] = c

	}

	var (
		minTempLocal = 0.0
		maxTempLocal = 0.0
		fKoeffLocal  = 0.0
		sumLocal     = 0.0
	)

	// @TODO replace with for i
	for i, v := range targetSources {
		curr := cMap[v].temp

		if curr < minTempLocal {
			// @TODO global to local
			minTempLocal = curr
		}

		if curr > maxTempLocal {
			// @TODO global to local
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
