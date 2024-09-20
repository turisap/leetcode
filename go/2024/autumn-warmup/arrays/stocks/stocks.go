package stocks

func stocksProfit(prices []int) int {
	profit := 0
	valley := prices[0]

	for _, v := range prices {
		if v < valley {
			valley = v
			continue
		}

		curDiff := v - valley

		if curDiff > profit {
			profit = curDiff
		}
	}

	return profit
}
