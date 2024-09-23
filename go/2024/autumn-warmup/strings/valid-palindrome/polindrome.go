package main

import (
	"regexp"
	"slices"
	"strings"
	"unicode"
)

//func main() {
//	res := ""
//	for _ = range 100 {
//		l := string(rand.Intn(100))
//		res = fmt.Sprintf("%s;%s %s", l, res, l)
//	}
//
//	fmt.Println(res)
//}

func isPalindromeReverse(s string) bool {
	tr := strings.Trim(s, "")
	l := strings.ToLower(tr)
	regEx := regexp.MustCompile("[^a-zA-Z0-9]+")
	prepared := regEx.ReplaceAllString(l, "")

	split := strings.Split(prepared, "")
	slices.Reverse(split)
	r := strings.Join(split, "")

	return r == prepared
}

func isPalindromeTwoPointer(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j {
		l := string(s[i])
		r := string(s[j])

		if !isAlph(l) {
			i++
			continue
		}

		if !isAlph(r) {
			j--
			continue
		}

		if strings.ToLower(l) != strings.ToLower(r) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlph(s string) bool {
	r := rune(s[0])
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
