package main

import "sort"

func mjeNaive(nums []int) int {
	ht := map[int]int{}
	m := 0
	maxEl := 0

	for _, v := range nums {
		if count, ok := ht[v]; !ok {
			ht[v] = 1

			if m < 1 {
				m = 1
				maxEl = v
			}
		} else {
			ht[v] = count + 1
			if (count + 1) > m {
				m = count + 1
				maxEl = v
			}

		}
	}

	return maxEl
}

func mjeImproved(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

func mjeVoting(nums []int) int {
	maxEl, c := nums[0], 1

	for i := 1; i < len(nums); i++ {

		if nums[i] == maxEl {
			c++
			continue
		}

		if c == 0 {
			maxEl = nums[i]
			c = 1
			continue
		}

		c--
	}

	return maxEl
}

func mjeEDITORIAL(nums []int) int {
	maxEl, c := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if c == 0 {
			maxEl, c = nums[i], 1
		} else {
			if nums[i] == maxEl {
				c += 1
			} else {
				c -= 1
			}
		}
	}
	return maxEl
}
