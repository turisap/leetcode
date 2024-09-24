package is_subsequence

/*
"ahbgdc"
"abc"
*/
func isSubSequence(s, t string) bool {
	if s == "" {
		return true
	}

	j := 0
	for i := 0; i < len(t); i++ {
		if t[i] != s[j] {
			continue
		}

		if j == len(s)-1 {
			return true
		}

		j++
	}

	return false
}
