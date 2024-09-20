package rm_duplicates

import (
	"fmt"
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
		fmt.Println(i, curr, currUnique, nums)
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

/*
0, 0, 1, 1, 1, 2, 2, 3, 3, 4
0, 1, 1, 1, 2, 2, 3, 3, 4, 0

*/
