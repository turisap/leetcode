package rm_element

import (
	"slices"
)

func removeElementNaive(nums []int, val int) int {
	c := 0
	i := 0
	for i < len(nums) {
		curr := nums[i]

		if curr == val {
			nums = slices.Delete(nums, i, i+1)
			continue
		}

		i++
		c++
	}

	return c
}

func removeElementsPointers(nums []int, val int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return j
}
