package main

import "fmt"

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	
	for i := 0; i < l; i++ {
		steps := nums[i]
		for j := 1; j <= steps; j++ {
			if i + j < l && dp[i+j] == 0 {
				fmt.Println(i, j, i+1)
				dp[i+j] = dp[i] + 1
			}
			if i + j == l - 1 {
				return dp[i+j]
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{1,2,1,1,1}))
}