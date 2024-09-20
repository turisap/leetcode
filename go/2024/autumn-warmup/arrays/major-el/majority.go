package major_el

func majorityNaive(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	mEl := nums[0]
	for k, v := range m {
		mCount := m[mEl]
		if v > mCount {
			mEl = k
		}
	}

	return mEl
}

func majorityNaivePreallocate(nums []int) int {
	m := make(map[int]int, len(nums))

	for _, v := range nums {
		m[v]++
	}

	mEl := nums[0]
	for k, v := range m {
		mCount := m[mEl]
		if v > mCount {
			mEl = k
		}
	}

	return mEl
}

func majorityMoor(nums []int) int {
	mEl := nums[0]
	mC := 1

	for _, v := range nums {
		if mEl == v {
			mC++
		} else {
			mC--
		}

		if mC == 0 {
			mEl = v
			mC = 1
		}
	}

	return mEl
}
