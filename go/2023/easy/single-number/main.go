package main

func snrXOR(nums []int) (res int) {
	res = nums[0]

	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	return
}

func snrHashTable(nums []int) (res int) {
	ht := map[int]bool{}

	for _, number := range nums {
		if v, ok := ht[number]; !ok {
			ht[number] = true
		} else {
			ht[number] = !v
		}
	}

	for k, v := range ht {
		if v == true {
			return k
		}
	}

	return
}
