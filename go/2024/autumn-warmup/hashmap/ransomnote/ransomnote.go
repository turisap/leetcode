package ransomnote

import (
	"strings"
)

func canConstructNaive(ransomNote string, magazine string) bool {
	m := map[int32]int{}

	for _, c := range ransomNote {
		m[c]++
	}

	for _, c := range magazine {
		if _, ok := m[c]; ok {
			m[c]--
		}
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}

func canConstructNaiveImproved(ransomNote string, magazine string) bool {
	m := make(map[int32]int, len(ransomNote))

	for _, c := range magazine {
		m[c]++
	}

	for _, c := range ransomNote {
		if _, ok := m[c]; ok {
			m[c]--
			if m[c] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func canConstructOneCount(ransomNote string, magazine string) bool {
	for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
		}
	}

	return true
}
