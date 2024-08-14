package dynamicProgramming

func rodCutting(prices []int, n int) int {
	memo := make([]int, n+1)

	for i := range memo {
		memo[i] = -1
	}

	return rodCuttingHelper(prices, n, memo)
}

func rodCuttingHelper(prices []int, n int, memo []int) int {

	if n == 0 {
		return 0
	}

	if memo[n] != -1 {
		return memo[n]
	}

	maxProfit := 0

	for i := 1; i < n; i++ {
		profit := prices[i-1] + rodCuttingHelper(prices, n-1, memo)
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	memo[n] = maxProfit
	return maxProfit
}
