package main

func canPartition(nums []int) bool {
	n := len(nums)
	sum, max := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if sum % 2 == 1 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}
	// 其中 dp[i][j]表示从数组的[0,i]下标范围内选取若干个正整数（可以是0个）
	// 是否存在一种选取方案使得被选取的正整数的和等于 j
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// initiate first column
	for i := 0; i < n; i++ {
		// if dont pick any num, target is 0 as well, that makes sense
		dp[i][0] = true 
	}
	// initiate first row
	dp[0][nums[0]] = true
	// compute first table
	for i := 1; i < n; i++ {
		num := nums[i]
		for j := 1; j <= target; j++ {
			// j 是当前的target
			if j >= num {
				// 如果不选取 nums[i] 则 dp[i][j] = dp[i−1][j]
				// 如果选取 nums[i] 则 dp[i][j] = dp[i−1][j−nums[i]]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
			} else {
				// 因为j已经大于nums[i] 必然不可能选择nums[i]
				// 所以状态同无法选取nums[i]
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

func main() {
	
}