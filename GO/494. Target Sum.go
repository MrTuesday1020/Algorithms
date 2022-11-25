package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum-target < 0 || (sum-target) % 2 == 1 {
		return 0
	}
	neg := (sum-target) / 2
	n := len(nums)
	// init dp table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	// compute dp table
	//  dp[i][j] 表示在数组 nums 的前 i 个数中选取元素，使得这些元素之和等于 j 的方案数
	// 如果 j < num 则不能选 num 此时有 dp[i][j] = dp[i−1][j]
	// 如果 j ≥ num 分两种case讨论:
	// 		如果不选 num 方案数是 dp[i−1][j]
	//		如果选 num 方案数是 dp[i−1][j−num]
	// 		总结两种case 即 dp[i][j] = dp[i−1][j] + dp[i−1][j−num]
	dp[0][0] = 1
	for i := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	
	return dp[n][neg]
}

func main() {
	
}