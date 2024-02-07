package main

func aitIterative(a, b []int) []int {
	set := map[int]bool{}
	added := map[int]bool{}
	// @TODO benchmark with len
	res := []int{}

	for _, v := range a {
		set[v] = true
	}

	for _, v := range b {
		if set[v] && !added[v] {
			// @TODO index access
			res = append(res, v)
			added[v] = true
		}
	}

	return res
}

// @COOL good improvement
func aitImproved(a, b []int) []int {
	rL := len(b)

	if len(a) < len(b) {
		rL = len(a)
	}

	set := make(map[int]bool, rL)
	added := make(map[int]bool, rL)
	res := make([]int, rL, rL)

	for _, v := range a {
		set[v] = true
	}

	c := 0
	for _, v := range b {
		if set[v] && !added[v] {
			res[c] = v
			added[v] = true
			c++
		}
	}

	return res[:c]
}
