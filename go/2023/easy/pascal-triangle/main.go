package main

func ptrNaive(n int) [][]int {
	// @TODO n length
	res := [][]int{{1}, {1, 1}}

	if n < 3 {
		return res[:n]
	}

	for i := 3; i <= n; i++ {
		res = append(res, make([]int, i))
	}

	for r := 2; r < len(res); r++ {
		row := res[r]
		prevRow := res[r-1]

		for el := 0; el < len(row); el++ {
			if el == 0 || el == (len(row)-1) {
				row[el] = 1
			} else {
				row[el] = prevRow[el-1] + prevRow[el]
			}
		}
	}

	return res
}
