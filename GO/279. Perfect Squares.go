package main

// dp[i] 表示数字i最少可以由几个完全平方数相加构成
// 位置i只依赖 i-j*j 的位置，如 i-1、i-4、i-9 等等位置，才能满足完全平方分割的条件。
// 因此dp[i]可以取的最小值即为 1 + min(dp[i-1],dp[i-4],dp[i-9]...)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n ;i++ {
		dp[i] = i   // worst case 所有都由1组成
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	
}