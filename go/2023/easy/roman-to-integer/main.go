package main

var strMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var doubleMap = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func RTIReg(s string) (r int) {
	for i := 0; i < len(s); {
		single := string(s[i])
		sEnd := i + 2

		if sEnd > len(s) {
			sEnd = len(s)
		}

		double := string(s[i:sEnd])

		v, ok := doubleMap[double]

		if ok {
			r += v
			i += 2

			continue
		}

		r += strMap[single]
		i++
	}

	return
}

func RTIL(s string) (r int) {

	for i, v := range s {
		r += strMap[string(v)]
		if i != 0 {
			if strMap[string(s[i-1])] < strMap[string(v)] {
				r -= 2 * strMap[string(s[i-1])]
			}
		}
	}

	return r
}
