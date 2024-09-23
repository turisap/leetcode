package l_word

import "strings"

func lastWordNaive(str string) int {
	trimmed := strings.TrimSpace(str)
	words := strings.Split(trimmed, " ")

	return len(words[len(words)-1])
}

func lastWordIterative(s string) int {
	c := 0
	started := false
	for i := len(s) - 1; i >= 0; i-- {
		curr := string(s[i])
		if curr != " " && !started {
			c++
			started = true
			continue
		}

		if curr != " " && started {
			c++
			continue
		}

		if curr == " " && started {
			return c
		}
	}

	return c
}
