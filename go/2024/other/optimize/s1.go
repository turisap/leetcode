package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
@TODO list
Byte slice map key to my assignment
Stack allocation, memory ballast
Manual inline
For range to for
Boundary checks
Writing global writing local
No defer
use ints instead of floats
Byte slice map key to my assignment
Map  pointers instead of values
Custom hash map
struct sorting fields
Allocate the backing array of a slice on stack even if its size is larger than
or equal to 64K (but not larger than 10M)
preallocate map (or GC ballast)
Avoid unnecessary pointer dereferences in a loop
Avoid accessing fields of a struct in a loop though pointers to the struct
Use index tables instead of maps which key types have only a small set of possible
values
method on an interface and on a concrete struct
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

type Measure interface {
	getTemperature() float64
	getSource() string
}

type Celsius struct {
	temp      float64
	time      time.Time
	active    bool
	source    string
	footprint float64
}

func (c Celsius) getTemperature() float64 {
	return c.temp
}

func (c Celsius) getSource() string {
	return c.source
}

type Fahrenheit struct {
	temp      float64
	time      time.Time
	active    bool
	source    string
	footprint float64
}

func (f Fahrenheit) getTemperature() float64 {
	return f.temp
}

func (f Fahrenheit) getSource() string {
	return f.source
}

var sum float64 = 0.0
var maxTemp float64 = 0.0
var minTemp float64 = 0.0
var fKoeff = 0.0

type Result struct {
	avg     float64
	minTemp float64
	maxTemp float64
	fK      float64
}

func s1(f *os.File) Result {
	scanner := bufio.NewScanner(f)
	fMap := map[string]Fahrenheit{}
	cMap := map[string]Celsius{}
	// @TODO pre-alloc
	targetSources := []string{}

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

		if isTargetSource(source) {
			targetSources = append(targetSources, source)
		}

		fMap[source] = Fahrenheit{
			temp:   temperature,
			time:   time.Now(),
			active: false,
			source: source,
		}
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

	// @TODO for range keys access temperatures
	// calculate avg, min, max
	//

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

		if i%17 == 0 {
			fKoeff = calcFloatSth(curr) // complex float manipulations
		}

		sum += curr
	}

	return Result{
		avg:     sum / float64((len(targetSources))),
		minTemp: minTemp,
		maxTemp: maxTemp,
		fK:      fKoeff,
	}
}

// @TODO manual inline
func calcFloatSth(f float64) float64 {
	a := f + 3.4
	b := f - 20.32
	k := f * 3.1 / 4.1

	return 3.0*a + 5.2*b - k*2.2
}

// @TODO call a method on concrete type, not on an interface
func fToCelsius(m Measure) Measure {
	// @TODO assign repetitive cals to getTemperature to a var
	return Celsius{
		temp:      degreesFtoCelsius(m.getTemperature()),
		source:    m.getSource(),
		footprint: calcFootprint(m.getTemperature()),
	}
}

// @TODO to inline
// @TODO float to int s
func degreesFtoCelsius(f float64) float64 {
	return (f - 32) * 0.55
}

func calcFootprint(f float64) float64 {
	e := f * 2.0
	k := f / 3.87

	a := 7*e + 3*k

	return a / 5
}

// @TODO inline
func isTargetSource(s string) bool {
	return rand.Intn(2) == 1
}
