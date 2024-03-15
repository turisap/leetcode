package main

// 0, 1, 1, 1, 2,2
func rdwSwap(s []int) int {
	count := 1
	i := 1
	j := 1

	for i < len(s) {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count > 2 {
			i++
		} else {
			s[j] = s[i]
			i++
			j++
		}
	}

	return j

}

func rdwTwo(s []int) int {
	j := 2

	for i := 2; i < len(s); i++ {
		if s[i] != s[j-2] {
			s[j] = s[i]
			j++
		}
	}

	return j
}
