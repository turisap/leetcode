package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type ReturnType interface {
	constraints.Ordered | constraints.Unsigned | constraints.Signed
}

type Reducer[T ReturnType, V any] func(acc T, value V) T

func reduceCustom[T ReturnType, V []T](arr V, reducer Reducer[T, T]) (r T) {
	var initialAcc T

	for _, v := range arr {
		initialAcc = reducer(initialAcc, v)
	}

	return initialAcc
}

func main() {
	s := []string{"k", "skjd"}

	res := reduceCustom[string, []string](s, func(acc, val string) string {
		return acc + val
	})

	fmt.Println(res)
}
