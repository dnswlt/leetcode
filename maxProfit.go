package leetcode

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	sell := make([]int, l)
	buy := make([]int, l)
	sell[0] = 0
	buy[0] = -prices[0]
	if prices[1] > prices[0] {
		sell[1] = prices[1] - prices[0]
	}
	buy[1] = max(-prices[0], -prices[1])
	for i := 2; i < l; i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
	}
	return sell[l-1]
}
