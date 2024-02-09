package main

func reverse[T any](s []T) []T {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}

	return s
}
