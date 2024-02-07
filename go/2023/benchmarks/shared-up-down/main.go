package main

import (
	"math/rand"
)

func sharedDown(n int) {
	b := make([]byte, n, n)

	r := internalD(&b)
	_ = r
}

func sharedUp(n int) {
	r := internalU(n)

	_ = r
}

func internalD(b *[]byte) *[]byte {
	for i, _ := range *b {
		(*b)[i] = byte(rand.Intn(255))
	}

	return b
}

func internalU(n int) *[]byte {
	b := make([]byte, n, n)

	internalD(&b)

	for i, _ := range b {
		b[i] = byte(rand.Intn(255))
	}

	return &b
}
