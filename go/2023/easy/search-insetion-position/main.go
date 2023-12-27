package main

import (
	"fmt"
)

func SIPReg(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}

	for idx, v := range nums {
		if v == target {
			return idx
		}

		if target < v && target > nums[idx-1] {
			fmt.Println("here", target, v, idx)
			return idx
		}
	}

	return len(nums)
}

func SIPL(nums []int, target int) (p int) {
	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := (left + right) / 2

		if nums[pivot] == target {
			return pivot
		}

		if target < nums[pivot] {
			right = pivot - 1
		}

		if target > nums[pivot] {
			left = pivot + 1
		}
	}

	return left
}
