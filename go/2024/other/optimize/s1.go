package main

import (
	"bufio"
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
a map
a value struct
a slice
some functions to inline
defers
a global (some kind of average)
a for range
another map cycle for access
another map for keys (maybe a part of them)
some function which accepts interfaces
*/

type Measure struct {
	temp   float64
	time   time.Time
	active bool
	source string
}

var avg float64 = 0.0

func s1(f *os.File) float64 {
	scanner := bufio.NewScanner(f)
	measuresMap := map[string]Measure{}
	keysMap := map[string]bool{}

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

		keysMap[source] = true
		measuresMap[source] = Measure{
			temp:   temperature,
			time:   time.Now(),
			active: false,
			source: source,
		}
	}

	return 0.0
}
