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


// TOn SO1 

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0 
    }

    max := 0 
    minBuy := prices[0]

    for i := 1; i< len(prices); i++  {
        price := prices[i]
        if price < minBuy {
            minBuy = price 
        } else {
            prof := price-minBuy
            if prof > max {
                max = prof
            }
        }
    }
    return max 
}

