package main

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][3]int, length)
	dp[0][0] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0]+prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[length-1][1], dp[length-1][2])
}

func max(i ,j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	
}