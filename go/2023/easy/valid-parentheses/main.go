package main

var openMap = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func VPSL(s string) (valid bool) {
	//stack := []rune{}
	stack := make([]rune, len(s))

	for _, ch := range s {
		if _, ok := openMap[ch]; ok {
			stack = append(stack, ch)
			continue
		} else if len(stack) == 0 {
			return false
		} else if openMap[stack[len(stack)-1]] != ch {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func VPSReg(s string) bool {
	stack := []rune{}

	if len(s)%2 != 0 {
		return false
	}

	for _, r := range s {

		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}

			el := stack[len(stack)-1]

			stack = append([]rune{}, stack[:len(stack)-1]...)

			opposite := openMap[el]

			if opposite != r {
				return false
			}
		}

	}

	if len(stack) > 0 {
		return false
	}

	return true
}
