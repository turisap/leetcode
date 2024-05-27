package main

const N = 10 * 1024 * 1024 // 10M
var n = N / 2

func f() byte {
	var a [N]byte // allocated on stack
	for i := range a {
		a[i] = byte(i)
	}
	var s = a[:] // a slice with 10M elements

	var p = &a
	return s[n] + p[n]
}
