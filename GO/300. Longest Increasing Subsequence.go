package main

func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = 1
	ans := 1
	for i := 1; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	
}