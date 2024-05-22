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
@TODO Writing global writing local
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
*/

/* What is needed?
DONE a map
DONE  a value struct
DONE a slice
DONE some functions to inline
@TODO defer (separately)
DONE a global (some kind of average)
DONE a for range to replace with for i
DONE another map cycle for access
DONE another map for keys (maybe a part of them)
DONE some function which accepts interfaces
*/

var sum float64 = 0.0
var maxTemp float64 = 0.0
var minTemp float64 = 0.0
var fKoeff = 0.0

func s1(f *os.File) Result {
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

	// @TODO replace with for i
	for i, v := range targetSources {
		curr := cMap[v].temp

		if curr < minTemp {
			// @TODO global to local
			minTemp = curr
		}

		if curr > maxTemp {
			// @TODO global to local
			maxTemp = curr
		}

		if i%TARGET_KOEF_IDX == 0 {
			fKoeff = calcFloatSth(curr) // complex float manipulations
		}

		sum += curr
	}

	avgStr := fmt.Sprintf("%.2f", sum/float64((len(targetSources))))
	minTStr := fmt.Sprintf("%.2f", minTemp)
	maxTStr := fmt.Sprintf("%.2f", maxTemp)
	fKoefStr := fmt.Sprintf("%.2f", fKoeff)

	return Result{
		avg:     avgStr,
		minTemp: minTStr,
		maxTemp: maxTStr,
		fK:      fKoefStr,
	}
}
