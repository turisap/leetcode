package main

func CLSReg(n int) int {

	var closure func(a, n int) int

	closure = func(a, n int) int {
		if a > n {
			return 0
		}

		if a == n {
			return 1
		}

		return closure(a+1, n) + closure(a+2, n)

	}

	return closure(0, n)
}

func CLSL(n int) int {

	var closure func(a, n int) int
	memo := make([]int, n)

	closure = func(a, n int) int {
		if a > n {
			return 0
		}

		if a == n {
			return 1
		}

		if memo[a] == 0 {
			memo[a] = closure(a+1, n) + closure(a+2, n)
		}

		return memo[a]
	}

	return closure(0, n)
}

func CLSDynamic(n int) int {
	if n <= 2 {
		return n
	}

	steps := [2]int{1, 2}

	for i := 2; i < n; i++ {
		next := steps[0] + steps[1]
		temp := steps[1]

		steps[0] = temp
		steps[1] = next
	}

	return steps[1]
}
