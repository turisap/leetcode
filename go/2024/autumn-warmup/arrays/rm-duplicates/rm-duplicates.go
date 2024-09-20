package rm_duplicates

import (
	"slices"
)

func removeDuplicatesNaive(nums []int) int {
	uniqueCount := 1
	currUnique := nums[0]
	i := 1
	c := 1
	maxIt := len(nums)

	for {
		if c == maxIt {
			break
		}

		c++
		curr := nums[i]
		if curr == currUnique {
			slices.Delete(nums, i, i+1)
			continue
		}

		currUnique = curr
		uniqueCount++
		i++
	}

	return uniqueCount
}

func removeDuplicatesTwoPointers(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return j + 1
}

/*
	0, 0, 1, 1, 1, 2, 2, 3, 3, 4
	0, 0, 1, 1, 1, 2, 2, 3, 3, 4
*/
