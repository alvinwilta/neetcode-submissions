func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	min := prices[0]

	for i:=1;i<len(prices);i++ {
		profit := prices[i]-min
		if profit>maxProfit {
			maxProfit = profit
		}
		if prices[i]<min {
			min = prices[i]
		}
	}

	return maxProfit
}
