func maxProfit(prices []int) int {
	max := 0
	// days := [2]int
	minBuy := prices[0]

	for _, val := range prices {
		if val < minBuy {
			minBuy = val
			continue
		}

		if val-minBuy > max {
			max = val - minBuy
		}
	}
	return max
}