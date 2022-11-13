package main

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[length-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	
}