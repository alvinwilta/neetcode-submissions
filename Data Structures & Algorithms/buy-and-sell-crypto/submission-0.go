func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]
	for _, price := range prices {
		if price < min {
			min = price
		}

		if price-min > profit {
			profit = price-min
		}
	}

	return profit
}
