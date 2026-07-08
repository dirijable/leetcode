package easy


func maxProfit(prices []int) int {
	resultProfit := 0
	for start, end := 0, 1; end < len(prices); end++ {
		if prices[start] < prices[end] {
			resultProfit = max(resultProfit, prices[end]-prices[start])
			continue
		}
		start = end

	}

	return resultProfit
}
