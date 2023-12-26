package main

import (
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func LCPSort(strs []string) (lcp string) {
	sort.Sort(ByLength(strs))

	lcp = strs[0]

	for i := 1; i < len(strs); i++ {
		curr := strs[i]
		l := findMinLen(len(curr), len(lcp))

		lcp = lcp[:l]

		for j := 0; j < l; j++ {
			if curr[j] != lcp[j] {
				lcp = curr[:j]
				break
			}
		}
	}

	return
}

func LCPReg(strs []string) (lcp string) {
	lcp = strs[0]

	for i := 1; i < len(strs); i++ {
		curr := strs[i]
		l := findMinLen(len(curr), len(lcp))

		lcp = lcp[:l]

		for j := 0; j < l; j++ {
			if curr[j] != lcp[j] {
				lcp = curr[:j]
				break
			}
		}
	}

	return
}

func LCPL(strs []string) (lcp string) {
	if len(strs) == 1 {
		return strs[0]
	}

	lcp = strs[0]

	for i := 1; i < len(strs); i++ {
		curr := strs[i]
		j := 0

		for j < len(curr) && j < len(lcp) && curr[j] == lcp[j] {
			j++
		}

		lcp = lcp[:j]

		if len(lcp) == 0 {
			return ""
		}
	}

	return
}

func findMinLen(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func main() {
//	LCPSort([]string{"abs", "abdfs", "a", "aaa"})
//}
