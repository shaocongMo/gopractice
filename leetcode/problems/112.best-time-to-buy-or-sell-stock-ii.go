package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	old_prices := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > old_prices {
			profit += prices[i] - old_prices
		}
		old_prices = prices[i]
	}
	return profit
}
