package main

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bssNaive(d []int) int {
	res := 0

	for i := 0; i < len(d); i++ {

		for j := i + 1; j < len(d); j++ {
			res = getMax(res, d[j]-d[i])
		}
	}

	return res
}

func bssImproved(d []int) int {
	res := 0
	minPrice := d[0]

	for _, v := range d {
		if v < minPrice {
			minPrice = v
			continue
		}

		res = getMax(res, v-minPrice)
	}

	return res
}
