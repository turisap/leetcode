package rotate_arr

import "slices"

func rotateNaive(nums []int, k int) []int {
	targetEls := slices.Clone(nums[len(nums)-k:])

	return append(targetEls, nums[:len(nums)-k]...)
}
