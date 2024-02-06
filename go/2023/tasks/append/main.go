package main

func appendCustom[T any](slice []T, args ...T) []T {
	// if the length of the resulting slice does not exceed cap(slice)
	curL := len(slice)
	curC := cap(slice)
	availableSlots := curC - curL
	neededSlots := len(args)

	if availableSlots >= neededSlots {
		c := len(slice)
		k := 0
		for k < len(args) {
			slice[c] = args[k]
			c++
			k++
		}

		return slice
	}

	l := neededSlots + curL
	s := make([]T, l)

	for i := 0; i < l; i++ {
		if i < curL {
			s[i] = slice[i]
			continue
		}

		s[i] = args[i-curL]
	}

	return s
}
