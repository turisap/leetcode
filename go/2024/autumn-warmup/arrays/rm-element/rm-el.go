package rm_element

import "slices"

func removeElementNaive(nums []int, val int) int {
	c := 0
	i := 0
	for i < len(nums) {
		curr := nums[i]

		if curr == val {
			nums = slices.Delete(nums, i, i)
			i++
			continue
		}

		i++
		c++
	}

	return c
}
