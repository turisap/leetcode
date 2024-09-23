package lcp

import (
	"slices"
)

func longestCommonPrefixNaive(s []string) string {
	lcp := s[0]

	for i := 1; i < len(s); i++ {
		w := s[i]
		for j := 0; j < len(w) && j < len(lcp); j++ {
			if len(lcp) == 0 {
				return lcp
			}

			if w[j] != lcp[j] {
				lcp = lcp[:j]
				break
			}
		}
		if len(lcp) > len(w) {
			lcp = lcp[:len(w)]
		}
	}

	return lcp
}

func longestCommonPresort(s []string) string {
	slices.Sort(s)

	f := s[0]
	l := s[len(s)-1]

	i, j := 0, 0

	var lcp string
	for i < len(f) && j < len(l) {
		if l[i] != f[j] {
			return lcp
		}

		lcp = f[:i+1]
		i++
		j++
	}

	return lcp
}
