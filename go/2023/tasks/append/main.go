package main

func appendCustom[T any](slice []T, args ...T) []T {
	curL := len(slice)
	curC := cap(slice)
	availableSlots := curC - curL
	neededSlots := len(args)

	if availableSlots >= neededSlots {
		c := len(slice)
		k := 0
		w := slice[:curC]
		for k < len(args) {
			w[c] = args[k]
			c++
			k++
		}

		return w
	}

	l := neededSlots + curL
	s := make([]T, l, l)

	for i := 0; i < l; i++ {
		if i < curL {
			s[i] = slice[i]
			continue
		}

		s[i] = args[i-curL]
	}

	return s
}
