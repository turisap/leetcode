package main

import (
	"strings"
	"unicode/utf8"
)

func ispRecursive(input string) bool {
	if len(input) == 0 {
		return false
	}

	inner := func(s string) bool {

		count := len(strings.Split(s, ""))

		if count < 2 {
			return true
		}

		slice := strings.Split(s, "")
		if count == 2 {
			return slice[0] == slice[1]
		}

		return ispRecursive(s[1:len(s)-1]) && (slice[0] == slice[len(slice)-1])
	}

	return inner(input)
}

func ispRecursiveUTF(input string) bool {
	if len(input) == 0 {
		return false
	}

	inner := func(s string) bool {

		b := []byte(s)
		count := utf8.RuneCount(b)

		if count < 2 {
			return true
		}

		slice := strings.Split(s, "")
		if count == 2 {
			return slice[0] == slice[1]
		}

		return ispRecursive(s[1:len(s)-1]) && (slice[0] == slice[len(slice)-1])
	}

	return inner(input)
}

func ispIterative(input string) bool {
	s := strings.Split(input, "")

	if len(s) == 0 {
		return false
	}

	l := 0
	r := len(s) - 1

	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
