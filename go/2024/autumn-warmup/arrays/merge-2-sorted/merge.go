package main

import "sort"

func MergeNaive(nums1 []int, m int, nums2 []int, n int) {
	l := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[l]
		l++
	}

	sort.Ints(nums1)
}
