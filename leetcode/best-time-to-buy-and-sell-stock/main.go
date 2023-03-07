package main

func maxProfit(prices []int) int {
	var buyPrice = 10001
	var maxProfit int

	for _, price := range prices {
		if price < buyPrice {
			buyPrice = price
		}

		if sellProfit := price - buyPrice; maxProfit < sellProfit {
			maxProfit = sellProfit
		}
	}

	return maxProfit
}
