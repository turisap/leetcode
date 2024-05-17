package main

// @TODO named import benchmark
// @TODO map pointers
// @TODO map to custom hash table
// @TODO make it no range
func darSimple(nums []int, k int) bool {
	m := map[int]int{}

	for curridx, v := range nums {
		previdx, ok := m[v]

		if ok {
			if curridx-previdx <= k {
				return true
			}

		}
		m[v] = curridx
	}

	return false
}

func darMapPointers(nums []int, k int) bool {
	m := map[int]*int{}

	for curridx, v := range nums {
		previdx, ok := m[v]

		if ok {
			if curridx-*previdx <= k {
				return true
			}

		}
		m[v] = &curridx
	}

	return false
}
